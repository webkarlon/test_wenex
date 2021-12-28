package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"git.jetbrains.space/yabbi/dsp/bidder/api"
	"git.jetbrains.space/yabbi/dsp/bidder/api/types"
	"git.jetbrains.space/yabbi/dsp/bidder/dbtype"
	"git.jetbrains.space/yabbi/dsp/components/extrarpc"
	"github.com/globalsign/mgo/bson"
	"github.com/nexcode/wenex"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"test_wenex/api/test"
	"test_wenex/vast"
	"test_wenex/vast/vast_2"
)

const (
	bucket           = "mytestbkt3"
	scopeStorageFull = "https://www.googleapis.com/auth/devstorage.full_control"

	// scope for call speechToText API
	scopeSpeechToText = "https://www.googleapis.com/auth/cloud-platform"
	statusWait        = "wait"

	// statusCalc mean that document start recognize but
	// but no result have been recieved yet
	statusCalc = "calc"

	// statusReady mean the audio file recognize is finished
	// and we can construct vtt subtitles
	statusReady = "ready"

	// statusErr an error was received during
	// upload or recognize
	statusErr = "error"
)

type StorageResp struct {
	Name string `json:"name"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := path.Base(wd) + "/jsonapi"

	if path.Dir(wd)[0] != '/' {
		configPath = "jsonapi"
	}

	wnx, err := wenex.New(configPath, wenex.DefaultConfig(), nil)
	if err != nil {
		panic(err)
	}

	rpcBidder := extrarpc.NewClient(wnx.Config.Get("RPC-HOST-BIDDER").(string))
	testapi := test.New(rpcBidder)

	wnx.Router.StrictRoute("/test", "GET").MustChain(testapi.TestQuery)
	wnx.Router.StrictRoute("/bl/wrapper.xml", "GET").MustChain(func(w http.ResponseWriter, r *http.Request) {
		type paramType struct {
			Overlay  []string `json:"overlay"`
			Tracking string   `json:"tracking"`
			Videos   struct {
				URL      string `json:"url"`
				MimeType string `json:"mimetype"`
			} `json:"videos"`
		}

		param := paramType{
			Overlay:  []string{r.FormValue("overlay")},
			Tracking: r.FormValue("tracking"),
			Videos: struct {
				URL      string `json:"url"`
				MimeType string `json:"mimetype"`
			}{
				MimeType: "video/mp4",
			},
		}
		paramJSON, _ := json.Marshal(&param)

		vastObj := vast.VAST{
			Version: "3.0",
			Ads: []vast.Ad{{
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{Name: "Plazkart"},
					Creatives: []vast.Creative{{
						Linear: &vast.Linear{
							AdParameters: &vast.AdParameters{
								Parameters: string(paramJSON),
							},
							Duration: "00:00:15",
							MediaFiles: []vast.MediaFile{{
								Type:         "application/javascript",
								APIFramework: "VPAID",
								URI:          "https://cdn.adx.com.ru/vpaid-logs-2.js",
							}},
							VideoClicks: &vast.VideoClicks{},
						},
					}},
				},
			}},
		}

		vastBytes, err := xml.Marshal(vastObj)
		if err != nil {
			return
		}

		w.Write([]byte(xml.Header))
		w.Write(vastBytes)
	})

	wnx.Router.StrictRoute("/v/wrapper.xml", "GET").MustChain(func(w http.ResponseWriter, r *http.Request) {

		vastObj := vast_2.VAST20{}
		vastObj.New()

		vastBytes, err := xml.Marshal(vastObj)
		if err != nil {
			return
		}

		w.Write([]byte(xml.Header))
		w.Write(vastBytes)
	})

	wnx.Router.StrictRoute("/test/p", "POST").MustChain(func(w http.ResponseWriter, r *http.Request) {

		var banners types.GetBannerApproved
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(banners); err != nil {
			panic(err)
		}
	})

	wnx.Router.StrictRoute("/test/file", "POST").MustChain(func(w http.ResponseWriter, r *http.Request) {

		formFile, _, err := r.FormFile("fileBanner")
		if err != nil {
			fmt.Println("Err 1", err)
			return
		}

		var buf bytes.Buffer
		_, err = io.Copy(&buf, formFile)
		formFile.Seek(0, io.SeekStart)

		call := api.Call{
			Query: api.SubRequest{
				BannerID:  bson.ObjectIdHex("61c41883d41e06e9b70b7dc6"),
				SoundData: buf.Bytes(),
			},
		}

		fmt.Println(&call.Query)

		if err := getSubtitles(call); err != nil {
			fmt.Println("SUB:", err)
			return
		}

	})

	fmt.Println("start")
	if err = wnx.Run(); err != nil {
		panic(err)
	}
}

func getSubtitles(call api.Call) error {

	subRequest, ok := call.Query.(api.SubRequest)
	if !ok {
		return errors.New("Type assertion error")
	}

	audioData, err := ConvertVideoToAudio(subRequest.SoundData)
	if err != nil {
		return err
	}

	fmt.Println(len(audioData))

	err = Upload(audioData, subRequest.BannerID)
	if err != nil {
		return err
	}

	return nil
}

func Upload(soundData []byte, bannerID bson.ObjectId) error {

	jsonConf, err := ioutil.ReadFile("gca/google_auth.json")
	if err != nil {
		return err
	}

	jwt, err := google.JWTConfigFromJSON(jsonConf, scopeStorageFull, scopeSpeechToText)
	if err != nil {
		return err
	}

	client := jwt.Client(context.Background())

	u := fmt.Sprintf("https://storage.googleapis.com/upload/storage/v1/b/%s/o?uploadType=media&name=%s.flac", bucket, bannerID.Hex())
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(soundData))
	if err != nil {
		fmt.Println("NewRequest:", err)
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Do:", err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errStr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(errStr))
	}

	//ses := g.dbConfig.Ses.Copy()
	//defer ses.Close()
	//
	var storageResp StorageResp
	err = json.NewDecoder(resp.Body).Decode(&storageResp)

	//cs := ses.DB(g.dbConfig.DB).C(g.dbConfig.CollectionName)
	bannerSubtitles := dbtype.Subtitles{
		ID:     bannerID,
		Name:   storageResp.Name,
		Status: statusWait,
	}

	fmt.Println(bannerSubtitles)
	//
	//err = cs.Insert(bannerSubtitles)
	fmt.Println("Finish")
	return err
}

func ConvertVideoToAudio(data []byte) ([]byte, error) {
	videoFileName, err := saveVideoFile(data)
	if err != nil {
		removeFiles(videoFileName)
		return nil, err
	}

	audioFileName, err := convertToAudio(videoFileName)
	if err != nil {
		removeFiles(videoFileName, audioFileName)
		return nil, err
	}

	soundData, err := getSoundData(audioFileName)
	if err != nil {
		removeFiles(videoFileName, audioFileName)
		return nil, err
	}

	//removeFiles(videoFileName, audioFileName)

	return soundData, nil
}

func saveVideoFile(data []byte) (string, error) {
	fileName := bson.NewObjectId().Hex()

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return "", err
	}

	return fileName, err
}

func removeFiles(fileNames ...string) {
	for _, name := range fileNames {
		_ = os.Remove(name)
	}
}

func convertToAudio(fileName string) (string, error) {
	soundFileName := fileName + ".flac"

	cmd := exec.Command("ffmpeg", "-i", fileName, soundFileName)

	err := cmd.Run()
	if err != nil {
		return "", errors.New("audio conversion err")
	}

	return soundFileName, nil
}

func getSoundData(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
