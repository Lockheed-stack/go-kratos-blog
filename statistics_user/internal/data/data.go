package data

import (
	"statistics_user/internal/biz"
	"statistics_user/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewStatUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
	ETCD_reg *etcd.Registry
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	data := &Data{}

	// gorm
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}
	data.db = db
	// auto migrate struct to mysql
	err = db.AutoMigrate(&biz.StatUser{})
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	// init etcd etcd_client
	etcd_client, err := clientv3.New(clientv3.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		return data, nil, err
	}
	data.ETCD_reg = etcd.New(etcd_client)

	cleanup := func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Error(err)
		} else {
			sqlDB.Close()
		}
		etcd_client.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
