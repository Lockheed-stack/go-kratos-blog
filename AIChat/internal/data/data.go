package data

import (
	"AIChat/internal/conf"
	"encoding/json"
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
var ProviderSet = wire.NewSet(NewData, NewAIChatRepo)

type cloudflareAPISetting struct {
	apiBaseUrl string
	token      string
}
type CloudflareResp struct {
	Success bool `json:"success"`
}

// Data .
type Data struct {
	// TODO wrapped database client
	cfAPI            *cloudflareAPISetting
	ETCD_reg         *etcd.Registry
	log              *log.Helper
	http_client      *http.Client
	aiModelTextOnly  map[string]string
	aiModelTextToImg map[string]string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	data := &Data{
		log: log.NewHelper(logger),
		http_client: &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost: 5,
			},
		},
		aiModelTextOnly:  make(map[string]string),
		aiModelTextToImg: make(map[string]string),
	}

	// cloudflare api
	cfAPI := &cloudflareAPISetting{
		apiBaseUrl: c.Cloudflare.ApiBaseUrl,
		token:      c.Cloudflare.Token,
	}
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

	data.cfAPI = cfAPI
	err := cloudflareApiVerify(data) // cloudflare api testing
	if err != nil {
		return data, nil, err
	}

	// etcd setting
	etcd_client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return data, nil, err
	} else {
		data.log.Info("connecting to etcd successfully.")
	}
	data.ETCD_reg = etcd.New(etcd_client)

	cleanup := func() {
		etcd_client.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}

	return data, cleanup, nil
}

func cloudflareApiVerify(data *Data) error {

	verifyURL := "https://api.cloudflare.com/client/v4/user/tokens/verify"
	req, err := http.NewRequest("GET", verifyURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+data.cfAPI.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	resp_json := &CloudflareResp{}
	err = dec.Decode(resp_json)
	if err != nil {
		return err
	}
	if !resp_json.Success {
		return errors.New(500, "", "cloudflare api verification failed.")
	} else {
		data.log.Info("Cloudflare api verification success.")
	}

	return nil
}
