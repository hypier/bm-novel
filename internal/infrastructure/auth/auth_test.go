package auth

import (
	"bm-novel/internal/domain/user"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestSetAuth(t *testing.T) {

	u := user.User{
		UserID:   uuid.NewV4(),
		UserName: "fun",
	}

	for i := 0; i < 10; i++ {
		_, _ = SetJWT(&u)
	}

}
