package gochat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
)

// CgiBin ...
type CgiBin struct {
	appID     string
	appSecret string
	reply     *CgiBinReply
}

// CgiBinReply ...
type CgiBinReply struct {
	AccessToken string `json:"access_token"`
	Ticket      string `json:"ticket"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// GetAccessToken 获取普通AccessToken
func (p *CgiBin) GetAccessToken() error {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", p.appID, p.appSecret))

	if err != nil {
		return err
	}

	reply := new(CgiBinReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	p.reply = reply

	return nil
}

// GetTicket 获取 JSAPI ticket
func (p *CgiBin) GetTicket(accessToken string) error {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken))

	if err != nil {
		return err
	}

	reply := new(CgiBinReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	p.reply = reply

	return nil
}

// AccessToken ...
func (p *CgiBin) AccessToken() string {
	return p.reply.AccessToken
}

// Ticket ...
func (p *CgiBin) Ticket() string {
	return p.reply.Ticket
}

// ExpiresIn ...
func (p *CgiBin) ExpiresIn() int64 {
	return p.reply.ExpiresIn
}

// NewCgiBin returns new cgi-bin
func NewCgiBin(channel WXChannel) *CgiBin {
	setting := GetConfigWithChannel(channel)

	return &CgiBin{
		appID:     setting.AppID,
		appSecret: setting.AppSecret,
	}
}