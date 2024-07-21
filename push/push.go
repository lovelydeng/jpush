package push

import (
	"context"
	"fmt"
	"github.com/lovelydeng/jpush/core"
)

type Client struct {
	Core *core.Core
}

var pushClient *Client

func New(co *core.Core) *Client {
	pushClient = &Client{Core: co}
	return pushClient
}

func NewWithDefaultCore(appKey, appMasterSecret string) *Client {
	co := core.New(appKey, appMasterSecret)
	return &Client{Core: co}
}

func (c *Client) SimplePush(ctx context.Context, platform Platform, audience any, notification *Notification) error {
	entity := PushEntity{
		Platform:     string(platform),
		Audience:     audience,
		Notification: notification,
	}

	resp, err := c.Core.DoPostRequest(ctx, NormalPushAction, entity)
	if err != nil {
		return err
	}
	_ = resp

	if resp.Status != 200 {
		return fmt.Errorf("httpStatus: %d body: %+v\n", resp.Status, string(resp.Body))
	}
	return nil
}
