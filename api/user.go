package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UserService interface {
	Info() (*User, error)
	UpdatePwd(string, string) error
	UpdateUsername(string) error
	UpdateLogo(string) error
	GetSyncState() (*SyncState, error)
}

type UserServiceOp struct {
	client *Client
}

type User struct {
	Id       string `json:"UserId"`
	Name     string `json:"Username"`
	Email    string `json:"Email"`
	Verified bool   `json:"Verified"`
	Logo     string `json:"Logo"`
}

func (s UserServiceOp) Info() (*User, error) {
	url := fmt.Sprintf("%s/user/info?token=%s", s.client.APIHost, s.client.Account.Token)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r := &User{}
	_, err = s.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s UserServiceOp) UpdatePwd(oldPwd, newPwd string) error {
	url := fmt.Sprintf("%s/user/updatePwd?token=%s&oldPwd=%s&pwd=%s", s.client.APIHost, s.client.Account.Token, oldPwd, newPwd)
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

func (s UserServiceOp) UpdateLogo(path string) error {
	//Not working
	url := fmt.Sprintf("%s/user/updatePwd?token=%s", s.client.APIHost, s.client.Account.Token)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest("POST", url, body)
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

func (s UserServiceOp) UpdateUsername(name string) error {
	url := fmt.Sprintf("%s/user/updateUsername?token=%s&username=%s", s.client.APIHost, s.client.Account.Token, name)
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

func (s UserServiceOp) GetSyncState() (*SyncState, error) {
	url := fmt.Sprintf("%s/user/getSyncState?token=%s", s.client.APIHost, s.client.Account.Token)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r := &SyncState{}
	_, err = s.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
