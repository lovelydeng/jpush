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

## 目录结构
```plaintext
jpush-sdk-go/
├── internal/
│   ├── core/  核心代码
│   │ 
│   ├── push/  推送逻辑代码
│   │   ├── push.go             # 推送功能实现
│   │   └── push_test.go        # 推送功能测试代码
│   ├── report/  报告核心代码
│   │   ├── report.go           # 统计报告功能实现
│   │   └── report_test.go      # 统计报告功能测试代码
│   └── util/  脚手架
│       └── utils.go            # 辅助工具函数
├── examples/
│   └── main.go                 # 示例代码入口
├── tests/
│   ├── integration/
│   │   └── integration_test.go # 集成测试代码
│   ├── unit/
│   │   └── unit_test.go        # 单元测试代码
├── .gitignore                   # Git 忽略文件列表
├── LICENSE                      # 许可证文件
├── README.md                    # 项目说明文件
├── go.mod                       # Go 模块文件
└── go.sum                       # Go 依赖文件

