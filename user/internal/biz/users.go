package biz

import (
	"crypto/sha256"
	"encoding/hex"
	pb "user/api/users"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);uniqueIndex;not null"`
	Password string `gorm:"type:char(64);not null"`
	Role     uint8  `gorm:"type:tinyint;UNSIGNED;DEFAULT:2"`
}

type UserRepo interface {
	CheckDuplicateUsername(name string) bool

	CreateUser(name string, psw string) error
	RemoveUser(id uint32) error
	AuthLogin(name string, pwd string) (uint32, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) CreateOneUser(username string, passwd string) error {

	if uc.repo.CheckDuplicateUsername(username) {
		return pb.ErrorErrUserPreExisting("")
	}

	hash := sha256.New()
	hash.Write([]byte(passwd))
	hex_pwd := hex.EncodeToString(hash.Sum(nil))

	err := uc.repo.CreateUser(username, hex_pwd)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_USER_PRE_EXISTING":
			return pb.ErrorErrUserPreExisting("")
		default:
			return pb.ErrorErrUserInvalidRequest("")
		}
	}
	return nil
}

func (uc *UserUsecase) UserLoginAuth(username string, passwd string) (uint32, error) {
	hash := sha256.New()
	hash.Write([]byte(passwd))
	hex_pwd := hex.EncodeToString(hash.Sum(nil))
	UID, err := uc.repo.AuthLogin(username, hex_pwd)

	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return 0, pb.ErrorErrUserUsernamePasswordWrong("")
		}
	}

	return UID, nil
}

func (uc *UserUsecase) RemoveOneUser(id uint32) error {
	err := uc.repo.RemoveUser(id)

	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_USER_NOT_FOUND":
			return pb.ErrorErrUserNotFound("")
		default:
			return pb.ErrorErrUserInvalidRequest("")
		}
	}
	return nil
}
