package FetchRepostory

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Repostory struct {
	Url         string
	Filename    string
	Destination string
	Content     []byte
}

func Fetch(url, savePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	saveFilename := strings.Split(savePath, "/")
	ioutil.WriteFile(saveFilename[len(saveFilename)-1], data, 0666)
	return nil
}

// TODO: Make a FetchGithub function
// func FetchGithub(url, repostoryName string) {}
