package api

import (
	"fmt"

	"github.com/humboldtux/leanote/config"
)

type AuthService interface {
	Login(*config.Config) error
	Logout() error
	Register(*config.Config) error
}

type AuthServiceOp struct {
	client *Client
}

type AccountAuth struct {
	Ok       bool   `json:"ok"`
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (s AuthServiceOp) Login(conf *config.Config) error {
	url := fmt.Sprintf("%s/auth/login?email=%s&pwd=%s", s.client.APIHost, conf.Email, conf.Password)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, &s.client.Account)
	if err != nil {
		return err
	}

	return nil
}

func (s AuthServiceOp) Logout() error {
	url := fmt.Sprintf("%s/auth/logout?&token=%s", s.client.APIHost, s.client.Account.Token)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	var er response
	_, err = s.client.Do(req, &er)
	if err != nil {
		return err
	}

	return nil
}

func (s AuthServiceOp) Register(conf *config.Config) error {
	url := fmt.Sprintf("%s/auth/register?email=%s&pwd=%s", s.client.APIHost, conf.Email, conf.Password)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	var er response
	_, err = s.client.Do(req, &er)
	if err != nil {
		return err
	}

	return nil
}
