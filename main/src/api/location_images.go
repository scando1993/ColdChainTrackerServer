package api

import (
	"bytes"
	"encoding/json"
	"github.com/mr-tron/base58/base58"
	"io/ioutil"
	"net/http"
	"path"
)

var MainPort string

func GetImage(family, location string) (img []byte, err error) {
	//reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_location := reg.ReplaceAllString(location, "")
	//_location := url.QueryEscape(location)
	imagePath := path.Join(DataFolder, "images", base58.FastBase58Encoding([]byte(family)), location + ".png")
	logger.Log.Debugf("loading %s", imagePath)
	img, err = ioutil.ReadFile(imagePath)
	return
}

func GenerateImages(family string) {
	logger.Log.Debugf("generating images for %s", family)
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	type Payload struct {
		URL        string `json:"url"`
		DataFolder string `json:"data_folder"`
	}

	data := Payload{
		URL:        "http://localhost:" + MainPort + "/api/v1/data/" + family,
		DataFolder: path.Join(DataFolder, "images", base58.FastBase58Encoding([]byte(family))),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "http://localhost:" + AIPort + "/plot", body)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	defer resp.Body.Close()

	type responseStruct struct {
		Message string
		Success bool
	}
	var target responseStruct
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	logger.Log.Debug(target)
}
