package middlewares

import (
	"gateway/internal/conf"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMids)

var Kratos_log *log.Helper

type Mids struct {
	JWTHandler            func() gin.HandlerFunc
	CorsHandler           func() gin.HandlerFunc
	NormalUserAuthHandler func() gin.HandlerFunc
}

func NewMids(c *conf.Middlewares, logger log.Logger) *Mids {
	JwtKey = []byte(c.Jwt.JwtKey)
	Kratos_log = log.NewHelper(logger)

	if len(c.Jwt.AesKey)%16 != 0 {
		panic("the length of AES key should be 16, please check the config.yaml")
	}
	AESKey = []byte(c.Jwt.AesKey)
	if len(c.Jwt.Nonce)%12 != 0 {
		panic("the length of Nonce should be 12, please check the config.yaml")
	}
	Nonce = []byte(c.Jwt.Nonce)

	return &Mids{
		JWTHandler:            JwtMids,
		CorsHandler:           CorsMid,
		NormalUserAuthHandler: NormalUserAuth,
	}
}
