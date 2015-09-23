package config

import "os"

const (
	defaultAPI = "https://leanote.com/api"
)

type Config struct {
	API      string `json:"api"`
	Email    string
	Password string
	Editor   string `json:"editor"`
}

func New() (*Config, error) {
	c := &Config{}
	err := c.load()
	return c, err
}

func (c *Config) load() error {
	c.Email = os.Getenv("LEANOTE_EMAIL")
	c.Password = os.Getenv("LEANOTE_PWD")
	c.API = os.Getenv("LEANOTE_APIURL")
	return c.setDefaults()
}

func (c *Config) setDefaults() error {
	if c.API == "" {
		c.API = defaultAPI
	}
	return nil
}
