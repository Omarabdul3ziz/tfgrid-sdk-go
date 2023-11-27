package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/threefoldtech/tfgrid-sdk-go/farmerbot/constants"
	"github.com/threefoldtech/tfgrid-sdk-go/farmerbot/slice"
)

// Config is the configuration for farmerbot
type Config struct {
	Farm  Farm   `json:"farm" yaml:"farm" toml:"farm"`
	Nodes []Node `json:"nodes" yaml:"nodes" toml:"nodes"`
	Power Power  `json:"power" yaml:"power" toml:"power"`
	*sync.Mutex
}

// GetNodeByNodeID gets a node by id
func (c *Config) GetNodeByNodeID(nodeID uint32) (Node, error) {
	for _, n := range c.Nodes {
		if n.ID == nodeID {
			return n, nil
		}
	}

	return Node{}, fmt.Errorf("node %d not found", nodeID)
}

// UpdateNode updates a node in the config
func (c *Config) UpdateNode(node Node) error {
	for i, n := range c.Nodes {
		if n.ID == node.ID {
			c.Lock()
			c.Nodes[i] = node
			c.Unlock()
			return nil
		}
	}

	return fmt.Errorf("node %d not found", node.ID)
}

// FilterNodesPower filters ON or OFF nodes
func (c *Config) FilterNodesPower(states []PowerState) []Node {
	filtered := make([]Node, 0)
	for _, node := range c.Nodes {
		if slice.Contains(states, node.PowerState) {
			filtered = append(filtered, node)
		}
	}
	return filtered
}

// FilterAllowedNodesToShutDown filters nodes that are allowed to shut down
//
// nodes with public config can't be shutdown
// Do not shutdown a node that just came up (give it some time)
func (c *Config) FilterAllowedNodesToShutDown() []Node {
	filtered := make([]Node, 0)
	for _, node := range c.Nodes {
		if node.IsUnused() && !node.PublicConfig && !node.NeverShutDown &&
			time.Since(node.LastTimePowerStateChanged) >= constants.PeriodicWakeUpDuration {
			filtered = append(filtered, node)
		}
	}
	return filtered
}
