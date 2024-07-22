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
		Alert: "human detect alarm",
		Android: Android{
			Alert:  "human detect alarm",
			Title:  "human detect alarm",
			Intent: Intent{Url: "xvr://vistaflow.live/systemMessages"},
		},
	})
	if err != nil {
		panic(err)
	}
}
