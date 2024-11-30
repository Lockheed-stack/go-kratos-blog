package data

import (
	"bytes"
	"context"
	"encoding/gob"
	"gateway/api/articles"
	"gateway/api/category"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// type BlogInfo struct {
// 	// Detail articles.GetSingleArticleReply_RespondMsg
// 	CreatedAt string `protobuf:"bytes,1,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
// 	UpdatedAt string `protobuf:"bytes,2,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
// 	Title     string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
// 	Desc      string `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
// 	Content   string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
// 	PageView  uint32 `protobuf:"varint,6,opt,name=PageView,proto3" json:"PageView,omitempty"`
// }
// type ByteSlice struct {
// 	Array unsafe.Pointer
// 	Len   int
// 	Cap   int
// }

// func (a *BlogInfo) MarshalBinary() ([]byte, error) {

// 	// is_panic := false

// 	// defer func() {
// 	// 	if e := recover(); e != nil {
// 	// 		is_panic = true
// 	// 	}
// 	// }()

// 	// data := (unsafe.Pointer(a))
// 	// byteSlice := ByteSlice{
// 	// 	Array: data,
// 	// 	Len:   int(unsafe.Sizeof(*a)),
// 	// 	Cap:   int(unsafe.Sizeof(*a)),
// 	// }

// 	// bin_data := *(*[]byte)(unsafe.Pointer(&byteSlice))
// 	// fmt.Printf("%x\n", bin_data)

// 	// if is_panic {
// 	// 	return nil, errors.InternalServer("failed to marshal", "")
// 	// }
// 	// return bin_data, nil
// 	var buf bytes.Buffer
// 	enc := gob.NewEncoder(&buf)
// 	err := enc.Encode(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }
// func (a *BlogInfo) UnmarshalBinary(data []byte) error {
// 	// is_panic := false

// 	// defer func() {
// 	// 	if e := recover(); e != nil {
// 	// 		is_panic = true
// 	// 	}
// 	// }()
// 	// byteSlice := *(*ByteSlice)(unsafe.Pointer(&data))
// 	// tmp := (*BlogInfo)(byteSlice.Array)

// 	// // a.detail = tmp.detail
// 	// a.Desc = tmp.Desc
// 	// a.Title = tmp.Title
// 	// a.Content = tmp.Content
// 	// a.PageView = tmp.PageView
// 	// a.CreatedAt = tmp.CreatedAt
// 	// a.UpdatedAt = tmp.UpdatedAt

// 	// if is_panic {
// 	// 	return errors.InternalServer("failed to Unmarshal", "")
// 	// }
// 	var buf bytes.Buffer
// 	_, err := buf.Write(data)
// 	if err != nil {
// 		return err
// 	}
// 	dec := gob.NewDecoder(&buf)
// 	err = dec.Decode(a)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// blogs relate
func GetOneBlogRedis(rdb *redis.Client, key string) (*articles.DetailArticleInfo, error) {

	pipe := rdb.Pipeline()
	// get blog info
	blog_info := pipe.Get(context.Background(), key)
	// get blog pv
	blog_pv := pipe.Incr(context.Background(), "blog_pv:"+key)
	// val, err := rdb.Get(context.Background(), key).Result()
	_, err := pipe.Exec(context.Background())

	if err != nil {
		return nil, err
	}

	// Unmarshal
	result := &articles.DetailArticleInfo{}
	dec := gob.NewDecoder(bytes.NewBufferString(blog_info.Val()))
	err = dec.Decode(result)
	if err != nil {
		return nil, err
	}
	if blog_pv.Val() > int64(result.PageView) {
		result.PageView = uint32(blog_pv.Val())
	}
	return result, nil
}
func SetOneBlogRedis(rdb *redis.Client, key string, data *articles.DetailArticleInfo) error {

	// marshal
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return err
	}

	// set blog info
	_, err = rdb.SetNX(context.Background(), key, buf.Bytes(), 0).Result()
	if err != nil {
		return err
	}
	// set blog pv
	pv := rdb.Incr(context.Background(), "blog_pv:"+key)
	if pv.Val() < int64(data.PageView) {
		_, err = rdb.Set(context.Background(), "blog_pv:"+key, data.PageView, 0).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
func DelOneBlogKeyRedis(rdb *redis.Client, key string) error {

	_, err := rdb.Del(context.Background(), key).Result()

	if err != nil {
		return err
	}
	return nil
}
func SetBlogsListRedis(rdb *redis.Client, key string, data []*articles.DetailArticleInfo) error {
	// marshal
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return err
	}

	_, err = rdb.Set(context.Background(), key, buf.Bytes(), time.Hour*8).Result()
	if err != nil {
		return err
	}

	return nil
}
func GetBlogsListRedis(rdb *redis.Client, key string) ([]*articles.DetailArticleInfo, error) {

	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	// unmarshal
	var result []*articles.DetailArticleInfo
	dec := gob.NewDecoder(bytes.NewBufferString(val))
	err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func DelBatchBlogKeyRedis(rdb *redis.Client, keys []string) error {
	_, err := rdb.Del(context.Background(), keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

// category relate
func SetCategoryRedis(rdb *redis.Client, key string, data []*category.CategoryInfo) error {
	// marshal
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return err
	}
	_, err = rdb.Set(context.Background(), key, buf.Bytes(), time.Hour*8).Result()

	if err != nil {
		return err
	}
	return nil
}
func GetCategoryRedis(rdb *redis.Client, key string) ([]*category.CategoryInfo, error) {
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var result []*category.CategoryInfo
	dec := gob.NewDecoder(bytes.NewBufferString(val))
	err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// user relate
func SetUserUniqueviewAndPageviewRedis(rdb *redis.Client, key string, uv_value string) error {

	pipe := rdb.Pipeline()
	pipe.SAdd(context.Background(), "user_uv:"+key, uv_value)
	pipe.Expire(context.Background(), "user_uv:"+key, time.Hour*24)

	pipe.HIncrBy(context.Background(), "user_pv", key, 1)
	pipe.Expire(context.Background(), "user_pv", time.Hour*24)
	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func SetTodayUserNewBlogNumRedis(rdb *redis.Client, key string) error {

	pipe := rdb.Pipeline()

	pipe.HIncrBy(context.Background(), "user_blog_num", key, 1)
	pipe.Expire(context.Background(), "user_blog_num", time.Hour*24)

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
func GetUserUniqueviewAndPageviewRedis(rdb *redis.Client, key string) (pv int64, uv int64, err error) {
	pipe := rdb.Pipeline()

	uv_reslut := pipe.SCard(context.Background(), "user_uv:"+key)
	pv_result := pipe.HGet(context.Background(), "user_pv", key)

	_, err = pipe.Exec(context.Background())
	if err != nil {
		return 0, 0, err
	}
	uv = uv_reslut.Val()
	pv_num, _ := strconv.Atoi(pv_result.Val())
	pv = int64(pv_num)

	return pv, uv, nil
}
func GetTodayUserNewBlogNumRedis(rdb *redis.Client, key string) (int, error) {
	result, err := rdb.HGet(context.Background(), "user_blog_num", key).Result()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(result)
	return num, err

}
func GetAllUsersStatisticsInfo(rdb *redis.Client, keys []string) (pv map[string]string, uv map[string]int64, new_blogs_num map[string]string, err error) {
	pipe := rdb.Pipeline()
	pv_result := pipe.HGetAll(context.Background(), "user_pv")                  // get user_pv
	uv_redis_result := make(map[string]*redis.IntCmd, len(keys))                // get user_uv
	new_blogs_num_result := pipe.HGetAll(context.Background(), "user_blog_num") // get user_new_blogs_num

	for _, v := range keys {
		uv_redis_result[v] = pipe.SCard(context.Background(), "user_uv:"+v)
	}
	_, err = pipe.Exec(context.Background())
	if err != nil && err != redis.Nil {
		return nil, nil, nil, err
	}

	uv_result := make(map[string]int64, len(keys))
	for k, v := range uv_redis_result {
		uv_result[k] = v.Val()
	}
	return pv_result.Val(), uv_result, new_blogs_num_result.Val(), nil
}
func DelUserUniqueviewAndPageviewRedis(rdb *redis.Client, keys []string) error {

	// delete the keys: 'user_uv:uid','user_pv','user_blog_num'
	length := len(keys) + 2
	del_keys := make([]string, length)
	for i, v := range keys {
		del_keys[i] = "user_uv:" + v
	}
	del_keys[len(keys)] = "user_pv"
	del_keys[len(keys)+1] = "user_blog_num"

	_, err := rdb.Del(context.Background(), del_keys...).Result()

	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}
