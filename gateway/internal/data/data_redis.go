package data

import (
	"bytes"
	"context"
	"encoding/gob"
	"gateway/api/articles"

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

	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal
	result := &articles.DetailArticleInfo{}
	dec := gob.NewDecoder(bytes.NewBufferString(val))
	err = dec.Decode(result)
	if err != nil {
		return nil, err
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

	_, err = rdb.Set(context.Background(), key, buf.Bytes(), 0).Result()

	if err != nil {
		return err
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

func SetPageView(rdb *redis.Client, key string) error {

	return nil
}
