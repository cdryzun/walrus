package server

import (
	"context"
	"database/sql"

	"k8s.io/client-go/kubernetes"

	"github.com/seal-io/walrus/pkg/cache"
	"github.com/seal-io/walrus/pkg/casdoor"
	"github.com/seal-io/walrus/pkg/database"
	"github.com/seal-io/walrus/pkg/health"
	"github.com/seal-io/walrus/pkg/k8s"
	"github.com/seal-io/walrus/utils/gopool"
)

// registerHealthCheckers registers the health checkers into the global health registry.
func (r *Server) registerHealthCheckers(ctx context.Context, opts initOptions) error {
	k8sClientSet, err := kubernetes.NewForConfig(opts.K8sConfig)
	if err != nil {
		return err
	}

	cs := health.Checkers{
		health.CheckerFunc("k8s", getKubernetesHealthChecker(k8sClientSet)),
		health.CheckerFunc("k8sctrl", getKubernetesControllerHealthChecker(opts.K8sCacheReady)),
		health.CheckerFunc("database", getDatabaseHealthChecker(opts.DatabaseDriver)),
		health.CheckerFunc("gopool", getGoPoolHealthChecker()),
	}

	if opts.CacheDriver != nil {
		cs = append(cs, health.CheckerFunc("cache", getCacheHealthChecker(opts.CacheDriver)))
	}

	if r.EnableAuthn {
		cs = append(cs, health.CheckerFunc("casdoor", casdoor.IsConnected))
	}

	return health.Register(ctx, cs)
}

func getKubernetesHealthChecker(cs *kubernetes.Clientset) health.Check {
	return func(ctx context.Context) error {
		return k8s.IsConnected(ctx, cs.RESTClient())
	}
}

func getKubernetesControllerHealthChecker(done <-chan struct{}) health.Check {
	return func(ctx context.Context) error {
		select {
		case <-done:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func getDatabaseHealthChecker(databaseDriver *sql.DB) health.Check {
	return func(ctx context.Context) error {
		return database.IsConnected(ctx, databaseDriver)
	}
}

func getCacheHealthChecker(cacheDriver cache.Driver) health.Check {
	return func(ctx context.Context) error {
		return cache.IsConnected(ctx, cacheDriver)
	}
}

func getGoPoolHealthChecker() health.Check {
	return func(_ context.Context) error {
		return gopool.IsHealthy()
	}
}
