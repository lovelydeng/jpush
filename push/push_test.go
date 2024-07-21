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
	err := cli.SimplePush(context.Background(), PlatformAndroid, &Audience{
		RegistrationId: []string{"18071adc0206ccfe21d"},
	}, &Notification{
		Alert: "报警通知111",
		Android: Android{
			Alert:       "动检报警3333",
			Title:       "动检报警",
			Intent:      Intent{Url: "xvr://vistaflow.live/systemMessages"},
			BadgeAddNum: 1,
			BadgeClass:  "com.gscam.xvr.ui.main.MainActivity",
		},
	})
	if err != nil {
		panic(err)
	}
}
