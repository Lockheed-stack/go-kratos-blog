package data

import (
	"blog/internal/biz"
	"blog/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewArticleRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
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
	err = db.AutoMigrate(&biz.Article{})
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	//closing tasks
	cleanup := func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.NewHelper(logger).Error(err)
		} else {
			sqlDB.Close()
		}

		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
