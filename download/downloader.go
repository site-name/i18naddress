package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	VALIDATION_DATA_DIR  string
	MAIN_URL             = "https://chromium-i18n.appspot.com/ssl-address/data"
	VALIDATION_DATA_PATH string
	client               http.Client
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	parent := path.Dir(wd)

	VALIDATION_DATA_DIR = path.Join(parent, "data")
	VALIDATION_DATA_PATH = path.Join(VALIDATION_DATA_DIR, "%s.json")

	client = http.Client{
		Timeout: 10 * time.Second,
	}
}

func fetch(url string) io.ReadCloser {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// defer res.Body.Close()

	return res.Body
}

func getCountries() []string {
	rc := fetch(MAIN_URL)

	var data struct {
		Countries string `json:"countries,omitempty"`
		Id        string `json:"id,omitempty"`
	}
	err := json.NewDecoder(rc).Decode(&data)
	if err != nil {
		panic(err)
	}

	return append(strings.Split(data.Countries, "~"), "ZZ")
}

func main() {
	fmt.Println(getCountries())
}

func download(countryCode string) {
	_, err := os.Stat(VALIDATION_DATA_DIR)
	if err != nil {
		panic(err)
	}

}
