package middlewares

import (
	"context"
	"gateway/internal/conf"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewMids)

var (
	kratos_log *log.Helper
	redis_cli  *redis.Client
)

type Mids struct {
	JWTHandler            func() gin.HandlerFunc
	CorsHandler           func() gin.HandlerFunc
	NormalUserAuthHandler func() gin.HandlerFunc
}

func NewMids(c_mids *conf.Middlewares, c_data *conf.Data, logger log.Logger) *Mids {
	JwtKey = []byte(c_mids.Jwt.JwtKey)
	kratos_log = log.NewHelper(logger)

	// redis config
	rdb := redis.NewClient(&redis.Options{
		Addr: c_data.Redis.Addr,
		DB:   0,
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	redis_cli = rdb

	// scheduled purge redis AI visitors key.
	now := time.Now() // the time of service start
	nowStamp := now.Unix()
	tomorrowStamp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1).Unix()
	timer := time.NewTimer(time.Second * time.Duration(tomorrowStamp-nowStamp))
	go func() {
		for range timer.C {
			timer.Reset(time.Hour * 24)
			redis_Remove_Visitors_AI_key()
		}
	}()

	// check jwt setting
	if len(c_mids.Jwt.AesKey)%16 != 0 {
		panic("the length of AES key should be 16, please check the config.yaml")
	}
	AESKey = []byte(c_mids.Jwt.AesKey)
	if len(c_mids.Jwt.Nonce)%12 != 0 {
		panic("the length of Nonce should be 12, please check the config.yaml")
	}
	Nonce = []byte(c_mids.Jwt.Nonce)

	return &Mids{
		JWTHandler:            JwtMids,
		CorsHandler:           CorsMid,
		NormalUserAuthHandler: NormalUserAuth,
	}
}
