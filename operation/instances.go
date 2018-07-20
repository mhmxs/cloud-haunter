package operation

import (
	log "github.com/Sirupsen/logrus"
	ctx "github.com/hortonworks/cloud-haunter/context"
	"github.com/hortonworks/cloud-haunter/types"
)

func init() {
	ctx.Operations[types.Instances] = instances{}
}

type instances struct {
}

func (o instances) Execute(clouds []types.CloudType) []types.CloudItem {
	log.Debugf("[GETINSTANCES] Collecting long instances instances on: [%s]", clouds)
	itemsChan, errChan := collectInstances(clouds)
	return wait(itemsChan, errChan, "[GETINSTANCES] Failed to collect instances instances")
}