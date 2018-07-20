package operation

import (
	log "github.com/Sirupsen/logrus"
	ctx "github.com/hortonworks/cloud-haunter/context"
	"github.com/hortonworks/cloud-haunter/types"
	"github.com/hortonworks/cloud-haunter/utils"
)

func init() {
	ctx.Filters[types.Ownerless] = ownerless{}
}

type ownerless struct {
}

func (o ownerless) Execute(items []types.CloudItem) []types.CloudItem {
	log.Debugf("[OWNERLESS] Filtering instances (%d): [%s]", len(items), items)
	return filter(items, func(item types.CloudItem) bool {
		if !isInstance(item) {
			log.Debugf("[OWNERLESS] Filter does not apply for cloud item: %s", item.GetName())
			return true
		}
		inst := item.(*types.Instance)
		match := !utils.IsAnyMatch(inst.Tags, ctx.OwnerLabels[item.GetCloudType()])
		log.Debugf("[OWNERLESS] Instances: %s match: %b", inst.Name, match)
		return match
	})
}
