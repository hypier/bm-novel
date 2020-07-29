package redis

//
//import (
//	"bm-novel/internal/domain/user"
//	"fmt"
//	"testing"
//	"time"
//
//	uuid "github.com/satori/go.uuid"
//)
//
//func TestCacher_Put(t *testing.T) {
//
//	for i := 0; i < 10; i++ {
//		u := user.User{
//			UserID:   uuid.NewV4(),
//			UserName: "fun",
//		}
//		//usr, _ := json.Marshal(u)
//		err := GetChcher().Put("login:"+u.UserName, []byte(u.UserID.String()), time.Second*20)
//
//		get, err := GetChcher().Get("login:" + u.UserName)
//		fmt.Println(string(get))
//
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}
//
//func TestCacher_HPut(t *testing.T) {
//
//	for i := 0; i < 10; i++ {
//		u := user.User{
//			UserID:   uuid.NewV4(),
//			UserName: "fun",
//		}
//		key := "login:" + u.UserName
//		err := GetChcher().HPut(key, u.UserID.String(), []byte(u.UserID.String()), time.Hour)
//
//		get, err := GetChcher().HGet(key, u.UserID.String())
//		fmt.Println(string(get))
//
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}
//
//func TestCacher_HMSet(t *testing.T) {
//
//	key := "login:fun"
//	err := rdb.HMSet(ctx, key, uuid.NewV4().String(), uuid.NewV4().String(), uuid.NewV4().String(), uuid.NewV4().String())
//
//	fmt.Println(err)
//}
//
//func TestCacher_Exists(t *testing.T) {
//	key := "login:fun"
//	//field := "e212fa4f-de00-4a5b-b0d3-650ee5ebe79b"
//	//err := GetChcher().HExists(key, field)
//	result, err := rdb.Exists(ctx, key).Result()
//	fmt.Println(result, err)
//}
