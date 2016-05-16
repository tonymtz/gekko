package services

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

const (
	GET_DATA_URL = "https://www.googleapis.com/plus/v1/people/%v?access_token=%v"
)

type GoogleAPI interface {
	GetProfile(string) (*GoogleProfile, error)
}

type googleAPI struct {
	getDataURL string
	Debug      bool
}

func NewGoogleAPI() GoogleAPI {
	return &googleAPI{
		getDataURL: GET_DATA_URL,
	}
}

func (this *googleAPI) GetProfile(token string) (*GoogleProfile, error) {
	getUrl := fmt.Sprintf(this.getDataURL, "me", token)

	fmt.Println(token)

	resp, err := http.Get(getUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	profile := &GoogleProfile{}

	dumpData, _ := ioutil.ReadAll(resp.Body)

	if this.Debug {
		fmt.Printf("%s\n", string(dumpData))
	}

	err = json.Unmarshal(dumpData, &profile)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

type GoogleProfile struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Emails      []*GoogleProfileEmail `json:"emails"`
	Image       *GoogleProfileImage `json:"image"`
}

type GoogleProfileEmail struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type GoogleProfileImage struct {
	Url string `json:"url"`
}
