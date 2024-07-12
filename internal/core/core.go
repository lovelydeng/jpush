package core

import (
	"go.uber.org/zap"
	"net/http"
)

// Core 封装了与 JPush 服务端通信的功能
type Core struct {
	appKey          string
	appMasterSecret string
	baseURL         string
	httpClient      *http.Client
	Logger          *zap.Logger
}

const defaultBaseUrl = "https://api.jpush.cn"

// NewClient 创建一个新的 Core 实例
func New(appKey, appMasterSecret string, options ...Option) *Core {
	JpushCore := &Core{appKey: appKey, appMasterSecret: appMasterSecret}
	for _, opt := range options {
		opt.Apply(JpushCore)
	}

	if JpushCore.baseURL == "" {
		JpushCore.baseURL = defaultBaseUrl
	}

	if JpushCore.httpClient == nil {
		JpushCore.httpClient = http.DefaultClient
	}

	if JpushCore.Logger == nil {
		logger, _ := zap.NewProduction()
		JpushCore.Logger = logger
	}
	return JpushCore
}

type Option interface {
	Apply(c *Core)
}

func WithHttClient(client *http.Client) Option {
	return withHttpClient{client: client}
}

type withHttpClient struct {
	client *http.Client
}

func (w withHttpClient) Apply(c *Core) {
	c.httpClient = w.client
}

func WithBaseUrl(url string) Option {
	return withBaseUrl{baseUrl: url}
}

type withBaseUrl struct {
	baseUrl string
}

func (w withBaseUrl) Apply(c *Core) {
	c.baseURL = w.baseUrl
}

func WithZapLogger(logger *zap.Logger) Option {
	return withLogger{logger: logger}
}

type withLogger struct {
	logger *zap.Logger
}

func (w withLogger) Apply(c *Core) {
	c.Logger = w.logger
}
