# jpush-sdk-go


#### go get github.com/lovelydeng/jpush

```
   package main
   import("github.com/lovelydeng/jpush/push")
   cli = push.NewWithDefaultCore(os.Getenv("APP_KEY"), os.Getenv("APP_MASTER_SECRET"))
   err := cli.SimplePush(context.Background(), []string{"435464565463"}, &Message{
		MsgContent: "哈哈哈哈",
		Title:      "测试",
		Extras:     nil,
	})
	if err != nil {
		panic(err)
	}

```
