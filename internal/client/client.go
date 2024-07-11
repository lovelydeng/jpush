package client

import (
	"net/http"
	"sync"
)

// Jpush 封装了与 JPush 服务端通信的功能
type Jpush struct {
	appKey          string
	appMasterSecret string
	baseURL         string
	httpClient      *http.Client
}

const defaultBaseUrl = ""

var once sync.Once
var jpush *Jpush

// NewClient 创建一个新的 Jpush 实例
func NewClient(appKey, appMasterSecret string, options ...Option) *Jpush {
	once.Do(func() {
		jpush = &Jpush{}
		for _, opt := range options {
			opt.Apply(jpush)
		}

		if jpush.baseURL == "" {
			jpush.baseURL = defaultBaseUrl
		}

		if jpush.httpClient == nil {
			jpush.httpClient = http.DefaultClient
		}
	})
	return jpush
}

type Option interface {
	Apply(c *Jpush)
}

func WithHttClient(client *http.Client) Option {
	return withHttpClient{client: client}
}

type withHttpClient struct {
	client *http.Client
}

func (w withHttpClient) Apply(c *Jpush) {
	c.httpClient = w.client
}

func WithBaseUrl(url string) Option {
	return withBaseUrl{baseUrl: url}
}

type withBaseUrl struct {
	baseUrl string
}

func (w withBaseUrl) Apply(c *Jpush) {
	c.baseURL = w.baseUrl
}
