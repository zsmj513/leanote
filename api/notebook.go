package api

import (
	"fmt"
	"strconv"
)

type NotebooksService interface {
	AddNotebook(string, string, int) (*Notebook, error)
	DeleteNotebook(string, int) error
	UpdateNotebook(string, string, string, int) (*Notebook, error)
	GetNotebooks() ([]Notebook, error)
	GetNotebook(string) (*Notebook, error)
	GetSyncNotebooks(int, int) ([]Notebook, error)
}

type NotebooksServiceOp struct {
	client *Client
}

type Notebook struct {
	CreatedTime      string `json:"CreatedTime"`
	IsBlog           bool   `json:"IsBlog"`
	IsDeleted        bool   `json:"IsDeleted"`
	NotebookId       string `json:"NotebookId"`
	ParentNotebookId string `json:"ParentNotebookId"`
	Seq              int    `json:"Seq"`
	Title            string `json:"Title"`
	UpdatedTime      string `json:"UpdatedTime"`
	UrlTitle         string `json:"UrlTitle"`
	UserId           string `json:"UserId"`
	Usn              int    `json:"Usn"`
}

func (s NotebooksServiceOp) AddNotebook(title string, parentId string, seq int) (*Notebook, error) {
	url := fmt.Sprintf("%s/notebook/addNotebook?token=%s&title=%s&parentNotebookId=%s&seq=%s", s.client.APIHost, s.client.Account.Token, title, parentId, strconv.Itoa(seq))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r *Notebook
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotebooksServiceOp) DeleteNotebook(notebookId string, usn int) error {
	url := fmt.Sprintf("%s/notebook/deleteNotebook?token=%s&notebookId=%s&usn=%s", s.client.APIHost, s.client.Account.Token, notebookId, strconv.Itoa(usn))
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

func (s NotebooksServiceOp) UpdateNotebook(notebookId string, title string, parentId string, seq int) (*Notebook, error) {
	url := fmt.Sprintf("%s/notebook/updateNotebook?token=%s&notebookId=%s&title=%s&parentNotebookId=%s&seq=%s", s.client.APIHost, s.client.Account.Token, notebookId, title, parentId, strconv.Itoa(seq))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var r *Notebook
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotebooksServiceOp) GetNotebooks() ([]Notebook, error) {
	url := fmt.Sprintf("%s/notebook/getNotebooks?token=%s", s.client.APIHost, s.client.Account.Token)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r []Notebook
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotebooksServiceOp) GetNotebook(notebookId string) (*Notebook, error) {
	notebooks, err := s.GetNotebooks()
	if err != nil {
		return nil, err
	} else {
		return &notebooks[0], nil
	}
}

func (s NotebooksServiceOp) GetSyncNotebooks(afterUsn, maxEntry int) ([]Notebook, error) {
	url := fmt.Sprintf("%s/notebook/getSyncNotebooks?token=%s&afterUsn=%s&maxEntry=%s", s.client.APIHost, s.client.Account.Token, strconv.Itoa(afterUsn), strconv.Itoa(maxEntry))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r []Notebook
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil

	if err != nil {
		return nil, err
	}
	return r, nil
}
