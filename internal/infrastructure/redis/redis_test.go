package redis

import (
	"bm-novel/internal/domain/user"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestCacher_Put(t *testing.T) {

	for i := 0; i < 10; i++ {
		u := user.User{
			UserID:   uuid.NewV4(),
			UserName: "fun",
		}
		//usr, _ := json.Marshal(u)
		err := GetChcher().Put("login:"+u.UserName, []byte(u.UserID.String()), time.Second*60)

		get, err := GetChcher().Get("login:" + u.UserName)
		fmt.Println(string(get))

		if err != nil {
			fmt.Println(err)
		}
	}
}
