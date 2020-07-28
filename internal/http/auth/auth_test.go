package auth

import (
	"bm-novel/internal/domain/user"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestSetAuth(t *testing.T) {

	u := user.User{
		UserID:   uuid.NewV4(),
		UserName: "fun",
	}

	for i := 0; i < 10; i++ {
		_, _ = setClientToken(&u, uuid.NewV4().String())
	}

}
