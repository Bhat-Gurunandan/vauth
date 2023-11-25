package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	defaultConfigFileName = ".vauth.json"
)

type Config struct {
	JWTSecret string `json:"jwtSecret,omitempty"`
	Layout    struct {
		Root string `json:"root,omitempty"`
	} `json:"layout,omitempty"`
	Db struct {
		Driver   string `json:"driver,omitempty"`
		Host     string `json:"host,omitempty"`
		Name     string `json:"name,omitempty"`
		User     string `json:"user,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"db,omitempty"`
	Remote struct {
		Host      string `json:"host,omitempty"`
		Port      int    `json:"port,omitempty"`
		User      string `json:"user,omitempty"`
		ImageRoot string `json:"remote_image_root,omitempty"`
		KeyPath   string `json:"key_path,omitempty"`
		HostKey   string `json:"host_key,omitempty"`
	} `json:"remote,omitempty"`
	SMTP struct {
		Server    string `json:"server,omitempty"`
		Port      int    `json:"port,omitempty"`
		User      string `json:"user,omitempty"`
		Password  string `json:"password,omitempty"`
		TestEmail string `json:"test_email,omitempty"`
	} `json:"smtp,omitempty"`
}

var c = Config{}

func Configuration(configFileName ...string) (Config, error) {

	if (c == Config{}) {

		var cfname string
		switch len(configFileName) {
		case 0:
			dirname, err := os.UserHomeDir()
			if err != nil {
				return c, err
			}
			cfname = fmt.Sprintf("%s/%s", dirname, defaultConfigFileName)
		case 1:
			cfname = configFileName[0]
		default:
			return c, fmt.Errorf("incorrect arguments for configuration file name")
		}

		configFile, err := os.Open(cfname)
		if err != nil {
			return c, err
		}
		defer configFile.Close()

		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&c)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}
