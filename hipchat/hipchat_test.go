package hipchat

import (
	"net/http"
	"testing"
	"time"

	"github.com/hortonworks/cloud-cost-reducer/types"
	"github.com/stretchr/testify/assert"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

func TestSend(t *testing.T) {
	mockClient := mockNotificationClient{make(chan *hipchat.NotificationRequest, 0)}

	go send("room", "message", mockClient)

	notifReq := <-mockClient.notifReqChan

	assert.Equal(t, "message", notifReq.Message)
	assert.Equal(t, hipchat.Color("green"), notifReq.Color)
	assert.Equal(t, "text", notifReq.MessageFormat)
}

func TestGenerateMessage(t *testing.T) {
	items := []types.CloudItem{
		types.Instance{
			CloudType: types.AWS,
			Name:      "instance",
			Created:   time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Owner:     "owner",
			Region:    "region",
		},
		types.Access{
			CloudType: types.AWS,
			Name:      "access",
			Created:   time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Owner:     "owner",
		},
	}
	op := types.LongRunning
	message := dispatcher.generateMessage(&op, items)

	assert.Equal(t, "/code\n[AWS] instance name: instance created: 1970-01-01 00:00:00 +0000 UTC owner: owner region: region\n[AWS] access: access created: 1970-01-01 00:00:00 +0000 UTC owner: owner\n", message)
}

type mockNotificationClient struct {
	notifReqChan chan *hipchat.NotificationRequest
}

func (m mockNotificationClient) Notification(id string, notifReq *hipchat.NotificationRequest) (r *http.Response, e error) {
	m.notifReqChan <- notifReq
	return
}
