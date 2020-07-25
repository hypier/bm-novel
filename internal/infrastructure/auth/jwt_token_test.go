package auth

import (
	"testing"
)

func TestSetAuth(t *testing.T) {

	auth := UserAuth{
		UserID:   "c67ee30c-9c76-4a2b-8d6c-063f114c616c",
		UserName: "fun",
		RoleCode: "admin",
	}

	for i := 0; i < 10; i++ {
		SetToken(auth)
	}

}
