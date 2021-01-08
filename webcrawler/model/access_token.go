package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
	"webcrawler/sdk"
)

type AccessTokenModel struct {
	gorm.Model
	AppId       string `gorm:"column:app_id"`
	appSecret   string
	AccessToken string `gorm:"column:access_token"`
	ExpiresIn   uint32
	ExpiresAt   time.Time
}

func NewAccessTokenModel() *AccessTokenModel {
	appId := sdk.NewWXClient().GetAppId()
	token := new(AccessTokenModel)
	err := db.First(token, "app_id = ?", appId).Error
	if err != nil {
		log.Errorf("DB_NewAccessTokenModel: %v", err)
		return nil
	}
	return token
}

func (tokene *AccessTokenModel) TableName() string {
	return "access_token"
}

func (token *AccessTokenModel) GetAccessToken() string {
	if token.isExpire() {
		err := token.updateAccessToken()
		if err != nil {
			return ""
		}
	}
	return token.AccessToken
}

func (token *AccessTokenModel) updateAccessToken() error {
	wxToken, err := sdk.NewWXClient().GetAccessToken()
	if err != nil {
		return err
	}
	access_token := strings.TrimSpace(wxToken.Access_token)
	expires_in := wxToken.Expires_in
	expires_at := time.Now().Add(time.Duration(expires_in) * time.Second)

	err = db.Model(token).Update("access_token", access_token).
		Update("expires_at", expires_at).
		Update("expires_in", expires_in).Error
	if err != nil {
		log.Errorf("DB_updateAccessToken: %v", err)
		return err
	}

	return nil
}

func (token *AccessTokenModel) isExpire() bool {
	if token.AccessToken == "" {
		return true
	}
	now := time.Now()
	if token.ExpiresAt.Before(now) {
		return true
	}
	return false
}
