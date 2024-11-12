package data

import (
	"context"
	"gateway/api/users"
	"gateway/internal/biz"
	"strconv"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type gatewayUserRepo struct {
	data         *Data
	log          *log.Helper
	timer        *time.Timer
	lock         sync.Mutex
	active_users map[uint64]struct{}
}

func NewGatewayUserRepo(data *Data, logger log.Logger) biz.GatewayUserRepo {

	now := time.Now() // the time of service start
	nowStamp := now.Unix()
	tomorrowStamp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1).Unix()
	timer := time.NewTimer(time.Second * time.Duration(tomorrowStamp-nowStamp))
	repo := &gatewayUserRepo{
		data:         data,
		log:          log.NewHelper(logger),
		timer:        timer,
		lock:         sync.Mutex{},
		active_users: make(map[uint64]struct{}),
	}

	// start a scheduled task
	go func() {
		for {
			select {
			case <-repo.timer.C:
				{
					// reset timer
					repo.timer.Reset(time.Hour * 24)
					// save user statistics info to DB
					repo.saveStatisticsInfoToDB()
				}
			case <-repo.data.Cancel_CTX.Done():
				{
					repo.timer.Stop()
					repo.log.Info("user timer stopped!")
					return
				}
			}
		}
	}()
	return repo
}

func (r *gatewayUserRepo) GRPC_CreateUser(req *users.CreateUsersRequest) (*users.CreateUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.CreateUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_DeleteUser(req *users.DeleteUsersRequest) (*users.DeleteUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.DeleteUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_AuthUser(req *users.AuthUsersRequest) (*users.AuthUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.AuthUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_GetSelectedUsers(req *users.GetSelectedUsersRequest) (*users.GetSelectedUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.GetSelectedUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_GetUserStatisticsInfo(req *users.GetStatisticsRequest) (*users.GetStatisticsReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.GetUserStatisticsInfo(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_UpdateUserPublicInfo(req *users.UpdateUserPublicInfoRequest) (*users.UpdateUserPublicInfoReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.UpdateUserPublicInfo(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}

// non-gRPC functions
func (r *gatewayUserRepo) MaintainUserStatisticsInfo(uid uint64, ip string, pv uint32, is_newBlog bool) error {

	// if add a new blog
	uid_str := strconv.Itoa(int(uid))
	if is_newBlog {
		err := SetTodayUserNewBlogNumRedis(r.data.Redis_cli, uid_str)
		if err != nil {
			return err
		}
		return nil
	}

	// set unique view
	err := SetUserUniqueviewAndPageviewRedis(r.data.Redis_cli, uid_str, ip)
	if err != nil {
		r.log.Error(err)
		return err
	}

	// update active users
	r.lock.Lock()
	r.active_users[uid] = struct{}{}
	r.lock.Unlock()

	return nil
}
func (r *gatewayUserRepo) TodayUserStatisticsInfo(uid uint64) (*users.StatisticsInfo, error) {

	resp := &users.StatisticsInfo{}

	key := strconv.Itoa(int(uid))
	pv, uv, err := GetUserUniqueviewAndPageviewRedis(r.data.Redis_cli, key)
	if err != nil && err != redis.Nil {
		return resp, err
	}
	blogs_num, err := GetTodayUserNewBlogNumRedis(r.data.Redis_cli, key)
	if err != nil && err != redis.Nil {
		return resp, err
	}

	resp.TotalPageviews = uint64(pv)
	resp.TotalUniqueviews = uint64(uv)
	resp.TotalBlogs = uint64(blogs_num)

	return resp, nil
}

// internal functions
func (r *gatewayUserRepo) saveStatisticsInfoToDB() error {

	req := &users.UpdateUserStatisticsInfoRequest{}

	// prepare data for update
	r.lock.Lock()
	activeUserNum := len(r.active_users)
	infos := make([]*users.StatisticsInfo, activeUserNum)
	keys := make([]string, activeUserNum)
	idx := 0
	for k := range r.active_users {
		tmp := new(users.StatisticsInfo)
		tmp.ID = k
		infos[idx] = tmp
		keys[idx] = strconv.Itoa(int(k))
		idx += 1
		delete(r.active_users, k)
	}
	r.lock.Unlock()

	// get statistics info from redis
	pv, uv, new_blogs_num, err := GetAllUsersStatisticsInfo(r.data.Redis_cli, keys)
	if err != nil {
		r.log.Error(err)
		return err
	}
	for i := range infos {
		id_str := strconv.Itoa(int(infos[i].ID))
		pv_num, _ := strconv.Atoi(pv[id_str])
		blogs_num, _ := strconv.Atoi(new_blogs_num[id_str])
		infos[i].TotalPageviews = uint64(pv_num)
		infos[i].TotalUniqueviews = uint64(uv[id_str])
		infos[i].TotalBlogs = uint64(blogs_num)
	}

	// save to DB
	req.Infos = infos
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	resp, err := client.UpdateUserStatisticsInfo(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return err
	} else if resp.Code != 200 {
		r.log.Error(resp.Msg)
	} else {
		r.log.Info("User statistics info has been saved to DB")
	}

	// clean redis key
	err = DelUserUniqueviewAndPageviewRedis(r.data.Redis_cli, keys)
	if err != nil {
		r.log.Errorf("An error occurd when clean User redis key:%v\n", err)
	}
	return nil
}
