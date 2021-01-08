package model

import (
	"fmt"
	"testing"
)

func TestWXModel_GetAccessToken(t *testing.T) {
	token, _ := NewWXModel().GetAccessToken()
	fmt.Println(token)
}
