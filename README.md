# jpush-sdk-go

## 目录结构

```plaintext
jpush-sdk-go/
├── internal/
│   ├── client/
│   │   ├── client.go           # 客户端 HTTP 请求处理
│   │   └── client_test.go      # 客户端测试代码
│   ├── push/
│   │   ├── push.go             # 推送功能实现
│   │   └── push_test.go        # 推送功能测试代码
│   ├── report/
│   │   ├── report.go           # 统计报告功能实现
│   │   └── report_test.go      # 统计报告功能测试代码
│   └── util/
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

