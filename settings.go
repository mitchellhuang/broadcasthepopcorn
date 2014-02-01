/*
* settings.go
* Manage settings (settings.json).
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JSONSettings struct {
	CacheDir string `json:"cache_dir"`
	Database string `json:"database"`
	PTP      struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Passkey  string `json:"passkey"`
		Settings struct {
			MovieSource     string `json:"movie_source"`
			MovieResolution string `json:"movie_resolution"`
		} `json:"settings"`
	} `json:"ptp"`
}

func NewJSONSettings() (JSONSettings, error) {
	var settings JSONSettings

	b, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		fmt.Println(err)
		return settings, err
	}

	if err := json.Unmarshal(b, &settings); err != nil {
		fmt.Println(err)
		return settings, err
	}

	return settings, nil
}
