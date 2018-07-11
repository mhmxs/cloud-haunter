package operation

import (
	"testing"
	"time"

	ctx "github.com/hortonworks/cloud-cost-reducer/context"
	"github.com/hortonworks/cloud-cost-reducer/types"
	"github.com/stretchr/testify/assert"
)

func TestLongRunningInit(t *testing.T) {
	assert.NotNil(t, ctx.Operations[types.LongRunning])
}

func TestLongRunningFilter(t *testing.T) {
	now := time.Now()
	items := []types.CloudItem{
		&types.Instance{
			CloudType: types.AWS,
			Name:      "short running",
			Created:   now.Add(-defaultRunningPeriod).Add(1 * time.Second),
		},
		&types.Instance{
			CloudType: types.AWS,
			Name:      "long running",
			Created:   now.Add(-defaultRunningPeriod).Add(-1 * time.Second),
		},
		&types.Instance{
			CloudType: types.AWS,
			Name:      "ignored",
			Tags:      types.Tags{ctx.AwsIgnoreLabel: "true"},
		},
	}

	filteredItems := longRunning{defaultRunningPeriod}.filter(items)

	assert.Equal(t, 1, len(filteredItems))
}
