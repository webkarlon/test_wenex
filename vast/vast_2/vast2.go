package vast_2

import (
	"encoding/json"
	"encoding/xml"
)

//Todo name change

type VAST20 struct {
	XMLName xml.Name `xml:"VAST"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns:MM,attr"`
	Script  xml.Name `xml:"script"`
	Ad      Ad       `xml:"Ad"`
}

type Ad struct {
	InLine InLine `xml:"InLine"`
}

type InLine struct {
	AdSystem  AdSystem  `xml:"AdSystem"`
	Creatives Creatives `xml:"Creatives"`
}

type AdSystem struct {
	CDATA string `xml:",cdata"`
}

type Creatives struct {
	Creative Creative `xml:"Creative"`
}

type Creative struct {
	Sequence string `xml:"sequence,attr"`
	Linear   Linear `xml:"Linear"`
}

type Linear struct {
	Duration       string        `xml:"Duration"`
	TrackingEvents string        `xml:"TrackingEvents"`
	AddParameters  AddParameters `xml:"AddParameters"`
	VideoClicks    VideoClicks   `xml:"VideoClicks"`
	MediaFiles     MediaFiles    `xml:"MediaFiles"`
}

type AddParameters struct {
	CDATA string `xml:",cdata"`
}

type VideoClicks struct {
	ClickThrough ClickThrough `xml:"ClickThrough"`
}

type ClickThrough struct {
	ID   string `xml:"id,attr"`
	Link string `xml:",chardata"`
}

type MediaFiles struct {
	MediaFile MediaFile `xml:"MediaFile"`
}

type MediaFile struct {
	Type         string `xml:"type,attr"`
	ApiFramework string `xml:"apiFramework,attr"`
	Link         string `xml:",chardata"`
}

type OverlayAd struct {
	Enabled             bool   `json:"enabled"`
	PackshotVisibleTime int    `json:"packshotVisibleTime"`
	MainText            string `json:"mainText"`
	ButtonText          string `json:"buttonText"`
}

type CDATA struct {
	ButtonAd       ButtonAd  `json:"buttonAd,omitempty"`
	OverlayAd      OverlayAd `json:"overlayAd,omitempty"`
	Sound          bool      `json:"sound,omitempty"`
	ProgressBar    bool      `json:"progressBar,omitempty"`
	Button         string    `json:"button,omitempty"`
	NeedSkipButton bool      `json:"needSkipButton,omitempty"`
	UnskipableTime int       `json:"unskipableTime,omitempty"`
	Logo           string    `json:"logo,omitempty"`
	AdMessage      string    `json:"adMessage,omitempty"`
	Videos         []Video   `json:"videos,omitempty"`
}

type Video struct {
	Url      string `json:"url"`
	MimeType string `json:"mimetype"`
}

type ButtonAd struct {
	Enabled          bool   `json:"enabled"`
	ButtonLayoutText string `json:"buttonLayoutText"`
}

func (v *VAST20) New() {

	mediaFile := MediaFile{
		Type:         "application/javascript",
		ApiFramework: "VPAID",
		Link:         "https://",
	}

	clickThrough := ClickThrough{
		ID:   "123",
		Link: "https://plazkart.ru/",
	}

	cdata := CDATA{
		ButtonAd:       ButtonAd{Enabled: false, ButtonLayoutText: "Доступная видео реклама"},
		OverlayAd:      OverlayAd{Enabled: false, PackshotVisibleTime: 5000, MainText: "ПЛАЦКАРТ - Доступная видео реклама", ButtonText: "Ознакомится подробнее"},
		Sound:          true,
		ProgressBar:    true,
		Button:         "Подробнее",
		NeedSkipButton: true,
		UnskipableTime: 10,
		Logo:           "dsafsdfsdfsdfsf",
		AdMessage:      "Тест тест",
		Videos: []Video{
			{
				Url:      "https://cdn.adx.com.ru/banner/sdfdfsdfsdf/",
				MimeType: "video/mp4",
			},
		},
	}

	paramJSON, _ := json.Marshal(&cdata)

	v.Version = "2.0"
	v.Xmlns = "https://"
	v.Ad.InLine.AdSystem.CDATA = "MEDIAMIND"
	v.Ad.InLine.Creatives.Creative.Sequence = "1"
	v.Ad.InLine.Creatives.Creative.Linear.Duration = "00:00:13"
	v.Ad.InLine.Creatives.Creative.Linear.MediaFiles.MediaFile = mediaFile
	v.Ad.InLine.Creatives.Creative.Linear.VideoClicks.ClickThrough = clickThrough
	v.Ad.InLine.Creatives.Creative.Linear.AddParameters.CDATA = string(paramJSON)
}
