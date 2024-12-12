package data

import (
	"AI_Service/internal/conf"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
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

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		etcd_client.Close()
	}

	return data, cleanup, nil
}
