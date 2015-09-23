package api

import "fmt"

type FilesService interface {
	GetImage(string) error
	GetAttach(string) error
	GetAllAttachs(string) error
}

type FilesServiceOp struct {
	client *Client
}

func (s FilesServiceOp) GetAllAttachs(noteId string) error {
	return fmt.Errorf("GetAllAttachs not implemented.")
}

func (s FilesServiceOp) GetAttach(fileId string) error {
	return fmt.Errorf("GetAttach not implemented.")
}

func (s FilesServiceOp) GetImage(fileId string) error {
	return fmt.Errorf("GetImage not implemented.")
}
