package config

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
)

// CustomConfig is the config of a custom connector.
// It is used to generate the custom connector config.
// E.g. A custom helm connector
//
//	configData := CustomConfig{
//		Attributes: map[string]interface{}{
//			"access_url": "http://localhost:8080",
//		},
//		Dependencies: []Dependency{
//			{
//				Type: "kubernetes",
//				Label: []string{},
//				Attributes: map[string]interface{}{
//					"config_path": "/home/user/.kube/config",
//				},
//			},
//		},
//	}
//
// This will generate the following terraform provider.
//
//	provider "helm" {
//		access_url = "http://localhost:8080"
//		kubernetes {
//			config_path = "/home/user/.kube/config"
//		}
//	}
type CustomConfig struct {
	// Attributes is the custom connector attribute
	// e.g. access_key, secret_key, etc.
	Attributes map[string]any `json:"attributes"`

	// TODO add block support, some custom connector may need Dependencies(blocks)
	// Dependencies is the dependencies of the custom connector.
	Dependencies []Dependency `json:"dependencies"`
}

// Dependency is the dependency of a custom connector.
type Dependency struct {
	Type       string         `json:"type"`
	Label      []string       `json:"label"`
	Attributes map[string]any `json:"attributes"`

	Children []Dependency `json:"children"`
}

// LoadCustomConfig loads the custom connector config from the connector.
func LoadCustomConfig(c *model.Connector) (*CustomConfig, error) {
	if c.Category != types.ConnectorCategoryCustom {
		return nil, fmt.Errorf("connector type is not custom connector: %s", c.ID)
	}

	cc := &CustomConfig{
		Attributes: make(map[string]any),
	}
	for k, d := range c.ConfigData {
		cc.Attributes[k] = d.Value
	}

	return cc, nil
}
