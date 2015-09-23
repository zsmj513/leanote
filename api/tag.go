package api

import (
	"fmt"
	"strconv"
)

type TagsService interface {
	GetSyncTags(int, int) ([]Tag, error)
	AddTag(string) (*Tag, error)
	DeleteTag(string) (*ReUpdate, error)
}

type TagsServiceOp struct {
	client *Client
}

type Tag struct {
	TagId       string `json:"TagId"`
	UserId      string `json:"UserId"`
	Title       string `json:"Tag"`
	CreatedTime string `json:"CreatedTime"`
	UpdatedTime string `json:"UpdatedTime"`
	IsDeleted   bool   `json:"IsDeleted"`
	Usn         int    `json:"Usn"`
}

func (s TagsServiceOp) GetSyncTags(afterUsn, maxEntry int) ([]Tag, error) {
	url := fmt.Sprintf("%s/tag/getSyncTags?token=%s&afterUsn=%s&maxEntry=%s", s.client.APIHost, s.client.Account.Token, strconv.Itoa(afterUsn), strconv.Itoa(maxEntry))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r []Tag
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s TagsServiceOp) AddTag(tag string) (*Tag, error) {
	url := fmt.Sprintf("%s/tag/addTag?token=%s&tag=%s", s.client.APIHost, s.client.Account.Token, tag)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r := &Tag{}
	_, err = s.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s TagsServiceOp) DeleteTag(tag string) (*ReUpdate, error) {
	//TODO:
	//// Get all tags
	//// Find the Tag to delete
	//// Get usn of tag
	u := "100000"
	url := fmt.Sprintf("%s/tag/deleteTag?token=%s&tag=%s&usn=%s", s.client.APIHost, s.client.Account.Token, tag, u)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r := &ReUpdate{}
	_, err = s.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
