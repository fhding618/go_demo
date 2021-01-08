package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"time"
	"webcrawler/base"
)

type WXModel struct {
	gorm.Model
	AppId       string `gorm:"column:app_id"`
	appSecret   string
	AccessToken string `gorm:"column:access_token"`
	ExpiresIn   uint32
	ExpiresAt   time.Time
}

func NewWXModel() *WXModel {
	return &WXModel{}
}

func (m *WXModel) TableName() string {
	return "wx"
}

func (m *WXModel) GetAccessToken() (string, error) {
	var wx WXModel
	err := db.First(&wx, "app_id = ?", base.WX_APP_ID).Error
	if err != nil {
		log.Errorf("DB_GetAccessToken: %v", err)
		return "", err
	}
	return wx.AccessToken, nil
}

func (m *WXModel) updateAccessToken() error {
	var wx WXModel
	err := db.First(&wx, "app_id = ?", base.WX_APP_ID).Error
	if err != nil {
		log.Errorf("DB_updateAccessToken: %v", err)
		return err
	}

	return nil
}

func (m *WXModel) accessTokenIsExpire() bool {
	if m.AccessToken == "" {
		return true
	}
	now := time.Now()

}
