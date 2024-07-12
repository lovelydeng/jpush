package push

import (
	"context"
	"fmt"
	"github.com/lovelydeng/jpush/internal/core"
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

func (c *Client) SimplePush(ctx context.Context, registrationIds []string, message *Message) error {
	entity := PushEntity{
		Platform: PlatformAll,
		Audience: Audience{
			RegistrationId: registrationIds,
		},
		Message: message,
	}

	resp, err := c.Core.DoPostRequest(ctx, NormalPushAction, entity)
	if err != nil {
		return err
	}
	_ = resp
	fmt.Printf("httpStatus: %d body: %+v\n", resp.Status, string(resp.Body))
	return nil
}
