package push

import (
	"context"
	"os"
	"testing"
)

var cli *Client

func TestMain(m *testing.M) {
	cli = NewWithDefaultCore(os.Getenv("APP_KEY"), os.Getenv("APP_MASTER_SECRET"))
	m.Run()
}

func TestClient_SimplePush(t *testing.T) {
	err := cli.SimplePush(context.Background(), &Audience{
		Tag:            nil,
		TagAnd:         nil,
		TagNot:         nil,
		Alias:          nil,
		RegistrationId: []string{"18071adc0206ccfe21d"},
		Segment:        nil,
		Abtest:         nil,
		LiveActivityID: "",
	}, &Notification{
		Alert: "报警通知111",
		Android: Android{
			Title:  "动检报警",
			Intent: Intent{Url: "xvr://vistaflow.live/systemMessages"},
		},
	})
	if err != nil {
		panic(err)
	}
}
