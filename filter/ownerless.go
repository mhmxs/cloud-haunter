package operation

import (
	ctx "github.com/hortonworks/cloud-haunter/context"
	"github.com/hortonworks/cloud-haunter/types"
	"github.com/hortonworks/cloud-haunter/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	ctx.Filters[types.OwnerlessFilter] = ownerless{}
}

type ownerless struct {
}

func (o ownerless) Execute(items []types.CloudItem) []types.CloudItem {
	log.Debugf("[OWNERLESS] Filtering instances (%d): [%s]", len(items), items)
	return filter("OWNERLESS", items, types.ExclusiveFilter, func(item types.CloudItem) bool {
		if !isInstance(item) {
			log.Fatalf("[OWNERLESS] Filter does not apply for cloud item: %s", item.GetName())
			return true
		}
		inst := item.(*types.Instance)
		match := !utils.IsAnyMatch(inst.Tags, ctx.OwnerLabel)
		log.Debugf("[OWNERLESS] Instances: %s match: %v", inst.Name, match)
		return match
	})
}
