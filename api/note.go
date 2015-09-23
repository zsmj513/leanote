package api

import (
	"fmt"
	"strconv"
	"time"
)

type NotesService interface {
	GetNotes() ([]Note, error)
	GetNoteAndContent(string) (*Note, error)
	NotesFromNotebook(string) ([]Note, error)
	GetSyncNotes(int, int) ([]Note, error)
}

type NotesServiceOp struct {
	client *Client
}

type NoteFile struct {
	FileId      string `json:"FileId"`
	LocalFileId string `json:"LocalFileId"`
	Type        string `json:"Type"`
	Title       string `json:"Title"`
	HasBody     bool   `json:"HasBody"`
	IsAttach    bool   `json:"IsAttach"`
}

type Note struct {
	NoteId      string     `json:"NoteId"`
	NotebookId  string     `json:"NotebookId"`
	UserId      string     `json:"UserId"`
	Title       string     `json:"Title"`
	Tags        []string   `json:"Tags"`
	Content     string     `json:"Content"`
	IsMarkdown  bool       `json:"IsMarkdown"`
	IsBlog      bool       `json:"IsBlog"`
	IsTrash     bool       `json:"IsBlog"`
	Files       []NoteFile `json:"Files"`
	CreatedTime time.Time  `json:"CreatedTime"`
	UpdatedTime time.Time  `json:"UpdateTime"`
	PublicTime  time.Time  `json:"PublicTime"`
	Usn         int        `json:"Usn"`
}

func (s NotesServiceOp) GetNotes() ([]Note, error) {
	url := fmt.Sprintf("%s/note/getNotes?token=%s", s.client.APIHost, s.client.Account.Token)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r []Note
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotesServiceOp) GetNoteAndContent(noteId string) (*Note, error) {
	url := fmt.Sprintf("%s/note/getNoteAndContent?token=%s&noteId=%s", s.client.APIHost, s.client.Account.Token, noteId)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r := &Note{}
	_, err = s.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotesServiceOp) NotesFromNotebook(notebookId string) ([]Note, error) {
	url := fmt.Sprintf("%s/note/getNotes?token=%s&notebookId=%s", s.client.APIHost, s.client.Account.Token, notebookId)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var r []Note
	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s NotesServiceOp) GetSyncNotes(afterUsn, maxEntry int) ([]Note, error) {
	url := fmt.Sprintf("%s/note/getSyncNotes?token=%s&notebookId=%safterUsn=%s&maxEntry=", s.client.APIHost, s.client.Account.Token, strconv.Itoa(afterUsn), strconv.Itoa(maxEntry))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var r []Note

	_, err = s.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
