package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/humboldtux/leanote/config"
)

const (
	libraryVersion = "0.1.0"
	userAgent      = "leanote-cli/" + libraryVersion
	mediaType      = "application/json"
)

type Client struct {
	APIHost   string
	Username  string
	client    *http.Client
	Account   AccountAuth
	Auth      AuthService
	Notes     NotesService
	Notebooks NotebooksService
	Tags      TagsService
	User      UserService
	Files     FilesService
}

type response struct {
	Ok  bool   `json:"Ok"`
	Msg string `json:"Msg"`
}

func NewClient(conf *config.Config) *Client {

	c := &Client{
		APIHost: conf.API,
		client:  http.DefaultClient,
		Account: AccountAuth{},
	}

	c.Notes = &NotesServiceOp{client: c}
	c.Notebooks = &NotebooksServiceOp{client: c}
	c.Tags = &TagsServiceOp{client: c}
	c.User = &UserServiceOp{client: c}
	c.Auth = &AuthServiceOp{client: c}
	c.Files = &FilesServiceOp{client: c}

	c.Auth.Login(conf)
	return c
}

func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", mediaType)

	return req, nil
}

// Do performs an http.Request and optionally parses the response body into the given interface
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusNoContent:
		return res, nil
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("an internal server error was received.")
	default:
		if v != nil {
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(v); err != nil {
				return nil, fmt.Errorf("error parsing API response - %s", err)
			}
		}
	}

	return res, nil
}
