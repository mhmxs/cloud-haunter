package action

import (
	"encoding/json"
	"fmt"

	ctx "github.com/hortonworks/cloud-haunter/context"
	"github.com/hortonworks/cloud-haunter/types"
	"github.com/hortonworks/cloud-haunter/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	ctx.Actions[types.Json] = new(jsonAction)
}

type jsonAction struct {
}

func (a jsonAction) Execute(op types.OpType, filter []types.FilterType, items []types.CloudItem) {
	log.Infof("[JSON] Number of items generated by operation %s and filters %s on accounts %s: %d", op.String(), filter, utils.GetCloudAccountNames(), len(items))
	out, _ := json.MarshalIndent(items, "", "  ")
	fmt.Println(string(out))
}
