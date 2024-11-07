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
	Username         string `gorm:"type:varchar(30);uniqueIndex;not null"`
	Password         string `gorm:"type:char(64);not null"`
	Role             uint8  `gorm:"type:tinyint;UNSIGNED;DEFAULT:2"`
	Avatar           string `gorm:"type:longtext"`
	SelfDesc         string `gorm:"type:varchar(150);DEFAULT:'nothing'"`
	TotalLoginDays   uint64 `gorm:"DEFAULT:1"`
	TotalBlogs       uint64 `gorm:"DEFAULT:1"`
	TotalPageviews   uint64 `gorm:"DEFAULT:1"`
	TotalUniqueviews uint64 `gorm:"DEFAULT:1"`
}

type UserRepo interface {
	CheckDuplicateUsername(name string) bool
	CreateUser(name string, psw string) error
	RemoveUser(id uint64) error
	AuthLogin(name string, pwd string) (*pb.UserPublicInfo, error)
	GetSelectedUsers(selectedFields []string, IDs []uint64) ([]*pb.UserPublicInfo, error)
	GetUserStatitics(id uint64) (*pb.StatisticsInfo, error)
	UpdateUserStatisticsInfo(infos []*pb.StatisticsInfo) (int64, error)
	UpdateUserPublicInfo(info *pb.UserPublicInfo) error
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

func (uc *UserUsecase) UserLoginAuth(username string, passwd string) (*pb.UserPublicInfo, error) {
	hash := sha256.New()
	hash.Write([]byte(passwd))
	hex_pwd := hex.EncodeToString(hash.Sum(nil))
	result, err := uc.repo.AuthLogin(username, hex_pwd)

	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return nil, pb.ErrorErrUserUsernamePasswordWrong("")
		}
	}

	return result, nil
}

func (uc *UserUsecase) RemoveOneUser(id uint64) error {
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

func (uc *UserUsecase) GetUsersByIDs(IDs []uint64) ([]*pb.UserPublicInfo, error) {
	if len(IDs) > 5 {
		return nil, pb.ErrorErrUserInvalidRequest("REQUEST_TOO_LONG")
	}

	selectedFidlds := []string{"id", "username", "avatar", "self_desc"}
	result, err := uc.repo.GetSelectedUsers(selectedFidlds, IDs)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return nil, pb.ErrorErrUserInvalidRequest("")
		}
	}

	return result, nil
}

func (uc *UserUsecase) GetStatisticsInfoByID(id uint64) (*pb.StatisticsInfo, error) {
	result, err := uc.repo.GetUserStatitics(id)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return nil, pb.ErrorErrUserInvalidRequest("")
		}
	}

	return result, nil
}

func (uc *UserUsecase) BatchUpdateUserStatisticsInfo(infos []*pb.StatisticsInfo) error {

	result, err := uc.repo.UpdateUserStatisticsInfo(infos)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return pb.ErrorErrUserInvalidRequest("")
		}
	} else if result < int64(len(infos)) {
		return pb.ErrorErrUserInvalidRequest("expect %v records updated, but only %v records updated", len(infos), result)
	}
	return nil
}

func (uc *UserUsecase) UpdateOneUserPublicInfo(info *pb.UserPublicInfo) error {

	err := uc.repo.UpdateUserPublicInfo(info)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_USER_PRE_EXISTING":
			return pb.ErrorErrUserPreExisting("username exist")
		default:
			return pb.ErrorErrUserInvalidRequest("")
		}
	}
	return nil
}
