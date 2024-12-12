package data

import (
	"AI_Service/internal/conf"
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAICloudflareRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	log *log.Helper
	// cloudflare setting
	cfAPIBaseUrl  string
	cfToken       string
	cf_httpClient *http.Client
	// etcd
	ETCD_Reg *etcd.Registry
	// AI models
	aiModelTextOnly  map[string]string
	aiModelTextToImg map[string]string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	data := &Data{
		log:          log.NewHelper(logger),
		cfAPIBaseUrl: c.Cloudflare.ApiBaseUrl,
		cfToken:      c.Cloudflare.Token,
		cf_httpClient: &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost: 5,
			},
		},
		aiModelTextOnly:  make(map[string]string),
		aiModelTextToImg: make(map[string]string),
	}

	// etcd setting
	etcd_client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return data, nil, err
	}
	data.ETCD_Reg = etcd.New(etcd_client)

	// read AI models name
	if len(c.Cloudflare.ModelTextOnly) == 0 || len(c.Cloudflare.ModelTextToImg) == 0 {
		err := errors.New(500, "", "Reading ModelTextOnly or ModelTextToImg Failed")
		return data, nil, err
	}
	for _, v := range c.Cloudflare.ModelTextOnly {
		tmp := strings.Split(v, "/")[2]
		data.aiModelTextOnly[tmp] = v
	}
	for _, v := range c.Cloudflare.ModelTextToImg {
		tmp := strings.Split(v, "/")[2]
		data.aiModelTextToImg[tmp] = v
	}

	// clean resource
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		etcd_client.Close()
	}

	return data, cleanup, nil
}
