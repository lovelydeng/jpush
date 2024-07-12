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
	err := cli.SimplePush(context.Background(), []string{"435464565463"}, &Message{
		MsgContent: "哈哈哈哈",
		Title:      "测试",
		Extras:     nil,
	})
	if err != nil {
		panic(err)
	}
}
