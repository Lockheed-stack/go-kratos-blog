package data

import (
	"context"
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratos_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGatewayUploadRepo,
	NewGatewayBlogRepo,
	NewGatewayCategoryRepo,
	NewGatewayUserRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	ETCD_reg          *etcd.Registry
	ConnGRPC_blog     *grpc.ClientConn
	ConnGRPC_category *grpc.ClientConn
	ConnGRPC_user     *grpc.ClientConn
	Redis_cli         *redis.Client
	// qiniuyun
	Qiniu_AccessKey      string
	Qiniu_SecretKey      string
	Qiniu_Bucket_Img     string
	Qiniu_Bucket_Article string
	cdnImg               string
	cdnArticle           string
	// context: for cancel goroutines and cleanup resources
	Cancel_CTX context.Context
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	data := &Data{}

	// etcd config
	etcd_client, err := clientv3.New(clientv3.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		return data, nil, err
	}
	data.ETCD_reg = etcd.New(etcd_client)

	// redis config
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   0,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		return data, nil, err
	}
	data.Redis_cli = rdb

	// connection of gprc
	conn1, err := kratos_grpc.DialInsecure(
		context.Background(),
		kratos_grpc.WithEndpoint("discovery:///blog"),
		kratos_grpc.WithDiscovery(data.ETCD_reg),
	)
	if err != nil {
		return data, nil, err
	}
	data.ConnGRPC_blog = conn1

	conn2, err := kratos_grpc.DialInsecure(
		context.Background(),
		kratos_grpc.WithEndpoint("discovery:///category"),
		kratos_grpc.WithDiscovery(data.ETCD_reg),
	)
	if err != nil {
		return data, nil, err
	}
	data.ConnGRPC_category = conn2

	conn3, err := kratos_grpc.DialInsecure(
		context.Background(),
		kratos_grpc.WithEndpoint("discovery:///user"),
		kratos_grpc.WithDiscovery(data.ETCD_reg),
	)
	if err != nil {
		return data, nil, err
	}
	data.ConnGRPC_user = conn3

	// qiniuyun config
	data.Qiniu_AccessKey = c.Qiniuyun.AccessKey
	data.Qiniu_SecretKey = c.Qiniuyun.SecretKey
	data.Qiniu_Bucket_Img = c.Qiniuyun.BucketImg
	data.Qiniu_Bucket_Article = c.Qiniuyun.BucketArticle
	data.cdnImg = c.Qiniuyun.CdnImg
	data.cdnArticle = c.Qiniuyun.CdnArticle

	// context
	ctx, cancel := context.WithCancel(context.Background())
	data.Cancel_CTX = ctx

	cleanup := func() {
		cancel()
		conn1.Close()
		conn2.Close()
		conn3.Close()
		etcd_client.Close()
		rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
