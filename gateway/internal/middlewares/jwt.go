package middlewares

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"crypto/aes"
	"crypto/cipher"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpiredDuration  = 1 * time.Hour
	RefreshTokenExpiredDuration = 30 * 24 * time.Hour
	TokenIssuer                 = ""
)

var (
	JwtKey []byte
	AESKey []byte
	Nonce  []byte
)

type MyClaims struct {
	Username string `json:"username"`
	UserID   string `json:"userid"`
	jwt.RegisteredClaims
}

func GetJWTtime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

// generate token
func GenerateToken(username string, userID uint64) (string, error) {

	RegisteredClaims := jwt.RegisteredClaims{
		ExpiresAt: GetJWTtime(AccessTokenExpiredDuration),
		Issuer:    "kratosBlog",
	}

	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, userID)

	encrypt_name, _ := encryptAES_GCM(block, []byte(username))
	encrypt_id, _ := encryptAES_GCM(block, buf.Bytes())
	SetClaims := MyClaims{
		encrypt_name,
		encrypt_id,
		RegisteredClaims,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	signedString, err := accessToken.SignedString(JwtKey)
	if err != nil {
		Kratos_log.Error(err.Error())
		return "", err
	}
	return signedString, nil
}

func encryptAES_GCM(block cipher.Block, src []byte) (string, error) {

	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aes_gcm.Seal(nil, Nonce, src, nil)

	return hex.EncodeToString(cipherText), nil
}
func decryptAES_GCM(block cipher.Block, ciphertext []byte) error {
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	_, err = aes_gcm.Open(ciphertext[:0], Nonce, ciphertext, nil)
	if err != nil {
		return err
	}
	return nil
}

// authenticate token
func AuthToken(token string) (uint64, error) {
	parseToken, err := jwt.ParseWithClaims(
		token,
		&MyClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)

	if err != nil {
		return 0, err
	}

	if claim, ok := parseToken.Claims.(*MyClaims); ok && parseToken.Valid {

		block, err := aes.NewCipher(AESKey)
		if err != nil {
			return 0, err
		}
		cipher_id, _ := hex.DecodeString(claim.UserID)
		err = decryptAES_GCM(block, cipher_id)
		if err != nil {
			return 0, err
		}

		plain_id := binary.BigEndian.Uint64(cipher_id)
		Kratos_log.Infof("auth user id:%v, Token is valid\n", plain_id)
		return plain_id, nil
	}
	return 0, nil
}

func JwtMids() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			ctx.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Unkown token type",
			})
			ctx.Abort()
			return
		}

		userid, err := AuthToken(checkToken[1])
		if err != nil || userid == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Authentication Failed",
			})
			ctx.Abort()
			return
		}
		ctx.Set("request_userid", int(userid))
	}
}
