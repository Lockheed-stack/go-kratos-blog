package data

import (
	"context"
	"gateway/api/stat_user"
	"gateway/api/users"
	"gateway/internal/biz"
	"strconv"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type gatewayUserRepo struct {
	data                *Data
	log                 *log.Helper
	timer               *time.Timer
	lock                sync.Mutex
	users_being_vivited map[uint64]struct{}
	active_users        map[uint64]struct{}
	stat_user_repo      biz.GatewayStatUserRepo
}

func NewGatewayUserRepo(data *Data, logger log.Logger, stat_user_repo biz.GatewayStatUserRepo) biz.GatewayUserRepo {

	now := time.Now() // the time of service start
	nowStamp := now.Unix()
	tomorrowStamp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1).Unix()
	timer := time.NewTimer(time.Second * time.Duration(tomorrowStamp-nowStamp))
	// timer := time.NewTimer(time.Second * 40)

	repo := &gatewayUserRepo{
		data:                data,
		log:                 log.NewHelper(logger),
		timer:               timer,
		lock:                sync.Mutex{},
		users_being_vivited: make(map[uint64]struct{}),
		active_users:        make(map[uint64]struct{}),
		stat_user_repo:      stat_user_repo,
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

	// if User Authentication Success
	if result.Code == 200 {
		r.lock.Lock()
		r.active_users[result.SelectedUser.ID] = struct{}{}
		r.lock.Unlock()
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
	// Only the user has been logged can call this function successfully, therefore we add 1 to TotalLoginDays.
	result.Info.TotalLoginDays += 1
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

	// update users_being_vivited
	r.lock.Lock()
	r.users_being_vivited[uid] = struct{}{}
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

	req_user := &users.UpdateUserStatisticsInfoRequest{}
	req_stat_user := &stat_user.SetUserStatInfoRequest{}

	// prepare data for update
	r.lock.Lock()
	visitedUserNum := len(r.users_being_vivited)
	activeUserNum := len(r.active_users)
	visited_and_active_user_map := make(map[uint64]*users.StatisticsInfo)
	today_user_stat_map := make(map[uint64]*stat_user.DayStatistics)

	req_user_infos := make([]*users.StatisticsInfo, 0, visitedUserNum+activeUserNum)
	req_stat_user_data := make([]*stat_user.DayStatistics, 0, visitedUserNum)
	keys := make([]string, visitedUserNum)

	// iterate over users_being_vivited
	idx := 0
	for k := range r.users_being_vivited {
		// UpdateUserStatisticsInfoRequest
		tmp_info := new(users.StatisticsInfo)
		tmp_info.ID = k
		visited_and_active_user_map[k] = tmp_info

		// SetUserStatInfoRequest
		tmp_data := new(stat_user.DayStatistics)
		tmp_data.Uid = k
		today_user_stat_map[k] = tmp_data

		// redis key
		keys[idx] = strconv.Itoa(int(k))
		idx += 1
	}

	// iterate over active_users
	for k := range r.active_users {
		if _, ok := visited_and_active_user_map[k]; !ok {
			tmp_info := new(users.StatisticsInfo)
			tmp_info.ID = k
			tmp_info.TotalLoginDays = 1
			visited_and_active_user_map[k] = tmp_info
		} else {
			visited_and_active_user_map[k].TotalLoginDays = 1
		}
	}

	// clear map
	r.active_users = make(map[uint64]struct{})
	r.users_being_vivited = make(map[uint64]struct{})
	r.lock.Unlock()

	// get statistics info from redis
	pv, uv, new_blogs_num, err := GetAllUsersStatisticsInfo(r.data.Redis_cli, keys)
	if err != nil {
		r.log.Error(err)
		return err
	}

	for k, v := range visited_and_active_user_map {
		id_str := strconv.Itoa(int(k))
		pv_num, _ := strconv.Atoi(pv[id_str])
		blogs_num, _ := strconv.Atoi(new_blogs_num[id_str])
		uv_num := uv[id_str]

		v.TotalPageviews = uint64(pv_num)
		v.TotalBlogs = uint64(blogs_num)
		v.TotalUniqueviews = uint64(uv_num)
		req_user_infos = append(req_user_infos, v)

		if val, ok := today_user_stat_map[k]; ok {
			val.Pv = uint64(pv_num)
			val.Uv = uint64(uv_num)
			req_stat_user_data = append(req_stat_user_data, val)
		}
	}

	// Save to the corresponding table in the database

	if len(req_user_infos) > 0 {
		// save to the user table in DB
		req_user.Infos = req_user_infos
		client := users.NewUsersClient(r.data.ConnGRPC_user)
		resp, err := client.UpdateUserStatisticsInfo(context.Background(), req_user)
		if err != nil {
			r.log.Error(err)
			return err
		} else if resp.Code != 200 {
			r.log.Error(resp.Msg)
		} else {
			r.log.Info("User statistics info has been saved to the user table in DB")
		}
	} else {
		r.log.Info("There are none User statistics info should be saved to the user table in DB")
	}

	if len(req_stat_user_data) > 0 {
		// save to the stat_user table in DB
		req_stat_user.Data = req_stat_user_data
		resp2, err := r.stat_user_repo.GRPC_SetUserTodayStatData(req_stat_user)
		if err != nil {
			r.log.Error(err)
			return err
		} else if resp2.Code != 200 {
			r.log.Error(resp2.Msg)
		} else {
			r.log.Info("User statistics info has been saved to the stat_user table in DB")
		}
	} else {
		r.log.Info("There are none User statistics info should be saved to the stat_user table in DB")
	}

	// clean redis key
	err = DelUserUniqueviewAndPageviewRedis(r.data.Redis_cli, keys)
	if err != nil {
		r.log.Errorf("An error occurd when clean User redis key:%v\n", err)
	}
	return nil
}
