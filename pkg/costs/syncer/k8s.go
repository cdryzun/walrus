package syncer

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/seal-io/walrus/pkg/costs/collector"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/costreport"
	opk8s "github.com/seal-io/walrus/pkg/operator/k8s"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/timex"
)

const (
	maxCollectTimeRange = 24 * time.Hour
	defaultStep         = 1 * time.Hour
)

type K8sCostSyncer struct {
	client model.ClientSet
	logger log.Logger
}

func NewK8sCostSyncer(client model.ClientSet, logger log.Logger) *K8sCostSyncer {
	if logger == nil {
		logger = log.WithName("cost")
	}

	return &K8sCostSyncer{
		client: client,
		logger: logger,
	}
}

func (in *K8sCostSyncer) SetLogger(logger log.Logger) {
	in.logger = logger
}

func (in *K8sCostSyncer) Sync(ctx context.Context, conn *model.Connector, startTime, endTime *time.Time) error {
	err := in.syncCost(ctx, conn, startTime, endTime)
	return err
}

func (in *K8sCostSyncer) syncCost(ctx context.Context, conn *model.Connector, startTime, endTime *time.Time) error {
	in.logger.Debugf("collect cost for connector: %s", conn.Name)

	apiConfig, _, err := opk8s.LoadApiConfig(*conn)
	if err != nil {
		return err
	}

	// NB(thxCode): disable timeout as we don't know the maximum time-cost of once full-series costs synchronization,
	// and rely on the session context timeout control,
	// which means we don't close the underlay kubernetes client operation until the `ctx` is cancel or timeout.
	restCfg, err := opk8s.GetConfig(*conn, opk8s.WithoutTimeout())
	if err != nil {
		return err
	}

	clusterName := apiConfig.CurrentContext

	collect, err := collector.NewCollector(restCfg, conn, clusterName)
	if err != nil {
		return err
	}

	startTime, endTime, err = in.timeRange(ctx, restCfg, conn, startTime, endTime)
	if err != nil {
		return err
	}

	in.logger.Debugf("connector: %s, current sync costs within %s, %s", conn.Name, startTime, endTime)

	curTimeRange := endTime.Sub(*startTime)
	maxTimeRange := maxCollectTimeRange

	if curTimeRange < maxTimeRange {
		maxTimeRange = curTimeRange
	}

	stepStart := *startTime
	for endTime.After(stepStart) {
		stepEnd := stepStart.Add(maxTimeRange)
		in.logger.Debugf("connector: %s, step sync within %s, %s", conn.Name, stepStart.String(), stepEnd.String())

		ac, err := collect.K8sCosts(&stepStart, &stepEnd, defaultStep)
		if err != nil {
			return err
		}

		if len(ac) == 0 {
			in.logger.Debugf("connector: %s, finished step sync within %s, %s, no record found",
				conn.Name, stepStart.String(), stepEnd.String())

			stepStart = stepEnd

			continue
		}

		if err = in.batchCreateCostReports(ctx, ac); err != nil {
			return fmt.Errorf("error creating item costs: %w", err)
		}

		in.logger.Debugf("connector: %s, finished step sync within %s, %s, created %d record",
			conn.Name, stepStart.String(), stepEnd.String(), len(ac))

		stepStart = stepEnd
	}

	return nil
}

func (in *K8sCostSyncer) batchCreateCostReports(ctx context.Context, costs []*model.CostReport) error {
	batchCreate := func(ctx context.Context, cs []*model.CostReport) error {
		return in.client.CostReports().CreateBulk().
			Set(cs...).
			OnConflictColumns(
				costreport.FieldStartTime,
				costreport.FieldEndTime,
				costreport.FieldConnectorID,
				costreport.FieldFingerprint,
			).
			DoNothing().
			Exec(ctx)
	}

	batchSize := 1000
	totalCosts := len(costs)

	for i := 0; i < totalCosts; i += batchSize {
		end := i + batchSize
		if end > totalCosts {
			end = totalCosts
		}

		batch := costs[i:end]

		err := batchCreate(ctx, batch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (in *K8sCostSyncer) timeRange(
	ctx context.Context,
	restCfg *rest.Config,
	conn *model.Connector,
	startTime,
	endTime *time.Time,
) (*time.Time, *time.Time, error) {
	// Time range existed.
	if startTime != nil && endTime != nil {
		return startTime, endTime, nil
	}

	// Time range from cluster.
	clientSet, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return nil, nil, err
	}

	nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	clusterEarliestTime := time.Now()
	for _, v := range nodes.Items {
		if v.CreationTimestamp.Time.Before(clusterEarliestTime) {
			clusterEarliestTime = v.CreationTimestamp.Time
		}
	}

	s := timex.StartTimeOfHour(clusterEarliestTime, time.UTC)
	e := timex.StartTimeOfHour(time.Now(), time.UTC)
	startTime = &s
	endTime = &e

	existed, err := in.client.CostReports().Query().
		Where(costreport.ConnectorID(conn.ID)).
		Order(model.Desc(costreport.FieldEndTime)).
		First(ctx)
	if err != nil {
		if model.IsNotFound(err) {
			return startTime, endTime, nil
		}

		return nil, nil, err
	}

	return &existed.EndTime, endTime, nil
}
