package config

import (
	"fmt"
	"testing"
)

func TestLoadDatabase(t *testing.T) {

	LoadConfig()

	fmt.Println(Config)
}
