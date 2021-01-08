package wx

import (
	"fmt"
	"testing"
)

func TestMyWX_GetAccessToken(t *testing.T) {
	wx := NewWXClient()
	token, err := wx.GetAccessToken()
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Printf("token: %v", token)
}
