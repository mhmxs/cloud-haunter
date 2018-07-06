package operation

import (
	"time"

	"os"

	log "github.com/Sirupsen/logrus"
	ctx "github.com/hortonworks/cloud-cost-reducer/context"
	"github.com/hortonworks/cloud-cost-reducer/types"
	"github.com/hortonworks/cloud-cost-reducer/utils"
)

var runningPeriod = 24 * time.Hour

func init() {
	ctx.Operations[types.LONGRUNNING] = LongRunning{}
	running := os.Getenv("RUNNING_PERIOD")
	if len(running) > 0 {
		if duration, err := time.ParseDuration(running); err != nil {
			log.Warnf("[LONGRUNNING] err: %s", err)
		} else {
			runningPeriod = duration
		}
	}
	log.Infof("[LONGRUNNING] running period set to: %s", runningPeriod)
}

type LongRunning struct {
}

func (o LongRunning) Execute(clouds []types.CloudType) []types.CloudItem {
	if ctx.DryRun {
		log.Debugf("Collecting long running instances on: [%s]", clouds)
	}
	itemsChan, errChan := collectRunningInstances(clouds)
	items := wait(itemsChan, errChan, "[LONGRUNNING] Failed to collect long running instances")
	return o.filter(items)
}

func (o LongRunning) filter(items []types.CloudItem) []types.CloudItem {
	if ctx.DryRun {
		log.Debugf("Filtering instances (%d): [%s]", len(items), items)
	}
	return filter(items, func(item types.CloudItem) bool {
		ignoreLabel, ok := ctx.IgnoreLabels[item.GetCloudType()]
		inst := item.(*types.Instance)
		match := (!ok || !utils.IsAnyMatch(inst.Tags, ignoreLabel)) && item.GetCreated().Add(runningPeriod).Before(time.Now())
		if ctx.DryRun {
			log.Debugf("Instances: %s match: %b", inst.Name, match)
		}
		return match
	})
}
