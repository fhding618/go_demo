package model

import (
	"fmt"
	"testing"
)

func TestWXModel_GetAccessToken(t *testing.T) {
	token := NewAccessTokenModel()
	access_token := token.GetAccessToken()
	fmt.Println(access_token)
}
