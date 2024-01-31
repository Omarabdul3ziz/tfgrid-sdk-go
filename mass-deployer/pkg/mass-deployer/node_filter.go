package deployer

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

func filterNodes(ctx context.Context, tfPluginClient deployer.TFPluginClient, group NodesGroup) ([]int, error) {
	log.Debug().Str("Node group", group.Name).Msg("Filter nodes")

	filter := types.NodeFilter{}

	statusUp := "up"
	freeMRU := convertMBToBytes(uint64(group.FreeMRU * 1024))

	filter.Status = &statusUp
	filter.TotalCRU = &group.FreeCPU
	filter.FreeMRU = &freeMRU

	if group.FreeSRU > 0 {
		freeSRU := convertGBToBytes(group.FreeSRU)
		filter.FreeSRU = &freeSRU
	}
	if group.FreeHRU > 0 {
		freeHRU := convertGBToBytes(group.FreeHRU)
		filter.FreeHRU = &freeHRU
	}
	if group.Region != "" {
		filter.Region = &group.Region
	}
	if group.Certified {
		certified := "Certified"
		filter.CertificationType = &certified
	}
	if group.PublicIP4 {
		filter.IPv4 = &group.PublicIP4
	}
	if group.PublicIP6 {
		filter.IPv6 = &group.PublicIP6
	}
	if group.Dedicated {
		filter.Dedicated = &group.Dedicated
	}
	freeSSD := []uint64{group.FreeSRU}
	if group.FreeSRU == 0 {
		freeSSD = nil
	}
	freeHDD := []uint64{group.FreeHRU}
	if group.FreeHRU == 0 {
		freeHDD = nil
	}

	nodes, err := deployer.FilterNodes(ctx, tfPluginClient, filter, freeSSD, freeHDD, nil, int(group.NodesCount))
	if err != nil {
		return []int{}, err
	}

	nodesIDs := []int{}
	for _, node := range nodes {
		nodesIDs = append(nodesIDs, node.NodeID)
	}

	return nodesIDs, nil
}