package context

import "github.com/hortonworks/cloud-cost-reducer/types"

var DRY_RUN bool = false

var Operations = make(map[types.OpType]types.Operation)

var CloudProviders = make(map[types.CloudType]types.CloudProvider)
