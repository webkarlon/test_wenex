package vast

import "encoding/xml"

type VAST struct {
	Version string  `xml:"version,attr"`
	Ads     []Ad    `xml:"Ad"`
	Errors  []CDATA `xml:"Error,omitempty"`
}

type Ad struct {
	ID       string   `xml:"id,attr,omitempty"`
	Sequence int      `xml:"sequence,attr,omitempty"`
	InLine   *InLine  `xml:",omitempty"`
	Wrapper  *Wrapper `xml:",omitempty"`
}

type CDATA struct {
	CDATA string `xml:",cdata"`
}

type InLine struct {
	AdSystem    *AdSystem
	AdTitle     CDATA
	Impressions []Impression `xml:"Impression,omitempty"`
	Creatives   []Creative   `xml:"Creatives>Creative"`
	Description CDATA        `xml:",omitempty"`
	Advertiser  string       `xml:",omitempty"`
	Survey      CDATA        `xml:",omitempty"`
	Errors      []CDATA      `xml:"Error,omitempty"`
	Pricing     *Pricing     `xml:",omitempty"`
	Extensions  *[]Extension `xml:"Extensions>Extension,omitempty"` // remove with tns counter
}

type Extension struct {
	Type      string     `xml:"type,attr,omitempty"`
	Trackings []Tracking `xml:"Tracking,omitempty"`
}

type Impression struct {
	ID  string `xml:"id,attr,omitempty"`
	URI string `xml:",cdata"`
}

type Pricing struct {
	Model    string `xml:"model,attr"`
	Currency string `xml:"currency,attr"`
	Value    string `xml:",cdata"`
}

type Wrapper struct {
	AdSystem     *AdSystem
	VASTAdTagURI CDATA
	Impressions  []Impression      `xml:"Impression"`
	Errors       []CDATA           `xml:"Error,omitempty"`
	Creatives    []CreativeWrapper `xml:"Creatives>Creative"`
	Extensions   []Extension       `xml:"Extensions>Extension,omitempty"` // remove with tns counter

	FallbackOnNoAd           *bool `xml:"fallbackOnNoAd,attr,omitempty"`
	AllowMultipleAds         *bool `xml:"allowMultipleAds,attr,omitempty"`
	FollowAdditionalWrappers *bool `xml:"followAdditionalWrappers,attr,omitempty"`
}

type AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	Name    string `xml:",cdata"`
}

type Creative struct {
	ID            string         `xml:"id,attr,omitempty"`
	Sequence      int            `xml:"sequence,attr,omitempty"`
	AdID          string         `xml:"AdID,attr,omitempty"`
	APIFramework  string         `xml:"apiFramework,attr,omitempty"`
	Linear        *Linear        `xml:",omitempty"`
	CompanionAds  *CompanionAds  `xml:",omitempty"`
	NonLinearAds  *NonLinearAds  `xml:",omitempty"`
	UniversalAdID *UniversalAdID `xml:"UniversalAdId,omitempty"`
	// CreativeExtensions *[]Extension   `xml:"CreativeExtensions>CreativeExtension,omitempty"`
}

type CompanionAds struct {
	Required   string      `xml:"required,attr,omitempty"`
	Companions []Companion `xml:"Companion,omitempty"`
}

type NonLinearAds struct {
	TrackingEvents []Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears     []NonLinear `xml:"NonLinear,omitempty"`
}

type CreativeWrapper struct {
	ID           string               `xml:"id,attr,omitempty"`
	Sequence     int                  `xml:"sequence,attr,omitempty"`
	AdID         string               `xml:"AdID,attr,omitempty"`
	Linear       *LinearWrapper       `xml:",omitempty"`
	CompanionAds *CompanionAdsWrapper `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty"`
}

type CompanionAdsWrapper struct {
	Required   string             `xml:"required,attr,omitempty"`
	Companions []CompanionWrapper `xml:"Companion,omitempty"`
}

type NonLinearAdsWrapper struct {
	TrackingEvents []Tracking         `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears     []NonLinearWrapper `xml:"NonLinear,omitempty"`
}

type Linear struct {
	// SkipOffset     *Offset `xml:"skipoffset,attr,omitempty"`
	Duration       string
	AdParameters   *AdParameters `xml:",omitempty"`
	Icons          *Icons
	TrackingEvents []Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks    *VideoClicks `xml:",omitempty"`
	MediaFiles     []MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`
}

type LinearWrapper struct {
	Icons          *Icons
	TrackingEvents []Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks    *VideoClicks `xml:",omitempty"`
}

type Companion struct {
	ID                     string          `xml:"id,attr,omitempty"`
	Width                  int             `xml:"width,attr,omitempty"`
	Height                 int             `xml:"height,attr,omitempty"`
	AssetWidth             int             `xml:"assetWidth,attr,omitempty"`
	AssetHeight            int             `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth          int             `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int             `xml:"expandedHeight,attr,omitempty"`
	APIFramework           string          `xml:"apiFramework,attr,omitempty"`
	AdSlotID               string          `xml:"adSlotId,attr,omitempty"`
	CompanionClickThrough  CDATA           `xml:",omitempty"`
	CompanionClickTracking []CDATA         `xml:",omitempty"`
	AltText                string          `xml:",omitempty"`
	TrackingEvents         []Tracking      `xml:"TrackingEvents>Tracking,omitempty"`
	AdParameters           *AdParameters   `xml:",omitempty"`
	StaticResource         *StaticResource `xml:",omitempty"`
	IFrameResource         CDATA           `xml:",omitempty"`
	HTMLResource           *HTMLResource   `xml:",omitempty"`
}

type CompanionWrapper struct {
	ID                     string          `xml:"id,attr,omitempty"`
	Width                  int             `xml:"width,attr"`
	Height                 int             `xml:"height,attr"`
	AssetWidth             int             `xml:"assetWidth,attr"`
	AssetHeight            int             `xml:"assetHeight,attr"`
	ExpandedWidth          int             `xml:"expandedWidth,attr"`
	ExpandedHeight         int             `xml:"expandedHeight,attr"`
	APIFramework           string          `xml:"apiFramework,attr,omitempty"`
	AdSlotID               string          `xml:"adSlotId,attr,omitempty"`
	CompanionClickThrough  CDATA           `xml:",omitempty"`
	CompanionClickTracking []CDATA         `xml:",omitempty"`
	AltText                string          `xml:",omitempty"`
	TrackingEvents         []Tracking      `xml:"TrackingEvents>Tracking,omitempty"`
	AdParameters           *AdParameters   `xml:",omitempty"`
	StaticResource         *StaticResource `xml:",omitempty"`
	IFrameResource         CDATA           `xml:",omitempty"`
	HTMLResource           *HTMLResource   `xml:",omitempty"`
}

type NonLinear struct {
	ID                     string          `xml:"id,attr,omitempty"`
	Width                  int             `xml:"width,attr"`
	Height                 int             `xml:"height,attr"`
	ExpandedWidth          int             `xml:"expandedWidth,attr"`
	ExpandedHeight         int             `xml:"expandedHeight,attr"`
	Scalable               bool            `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    bool            `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   string          `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework           string          `xml:"apiFramework,attr,omitempty"`
	NonLinearClickTracking []CDATA         `xml:",omitempty"`
	NonLinearClickThrough  CDATA           `xml:",omitempty"`
	AdParameters           *AdParameters   `xml:",omitempty"`
	StaticResource         *StaticResource `xml:",omitempty"`
	IFrameResource         CDATA           `xml:",omitempty"`
	HTMLResource           *HTMLResource   `xml:",omitempty"`
}

type NonLinearWrapper struct {
	ID                     string     `xml:"id,attr,omitempty"`
	Width                  int        `xml:"width,attr"`
	Height                 int        `xml:"height,attr"`
	ExpandedWidth          int        `xml:"expandedWidth,attr"`
	ExpandedHeight         int        `xml:"expandedHeight,attr"`
	Scalable               bool       `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    bool       `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   string     `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework           string     `xml:"apiFramework,attr,omitempty"`
	TrackingEvents         []Tracking `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinearClickTracking []CDATA    `xml:",omitempty"`
}

type Icons struct {
	XMLName xml.Name `xml:"Icons,omitempty"`
	Icon    []Icon   `xml:"Icon,omitempty"`
}

type Icon struct {
	Program   string `xml:"program,attr"`
	Width     int    `xml:"width,attr"`
	Height    int    `xml:"height,attr"`
	XPosition string `xml:"xPosition,attr"`
	YPosition string `xml:"yPosition,attr"`
	// Offset             Offset          `xml:"offset,attr"`
	Duration           string          `xml:"duration,attr"`
	APIFramework       string          `xml:"apiFramework,attr,omitempty"`
	IconClickThrough   CDATA           `xml:"IconClicks>IconClickThrough,omitempty"`
	IconClickTrackings []CDATA         `xml:"IconClicks>IconClickTracking,omitempty"`
	StaticResource     *StaticResource `xml:",omitempty"`
	IFrameResource     CDATA           `xml:",omitempty"`
	HTMLResource       *HTMLResource   `xml:",omitempty"`
}

type Tracking struct {
	Event string `xml:"event,attr"`
	// Offset *Offset `xml:"offset,attr,omitempty"`
	URI string `xml:",cdata"`
}

type StaticResource struct {
	CreativeType string `xml:"creativeType,attr,omitempty"`
	URI          string `xml:",cdata"`
}

type HTMLResource struct {
	XMLEncoded bool   `xml:"xmlEncoded,attr,omitempty"`
	HTML       string `xml:",cdata"`
}

type AdParameters struct {
	XMLEncoded bool   `xml:"xmlEncoded,attr,omitempty"`
	Parameters string `xml:",cdata"`
}

type VideoClicks struct {
	ClickThroughs  []VideoClick `xml:"ClickThrough,omitempty"`
	ClickTrackings []VideoClick `xml:"ClickTracking,omitempty"`
	CustomClicks   []VideoClick `xml:"CustomClick,omitempty"`
}

type VideoClick struct {
	ID  string `xml:"id,attr,omitempty"`
	URI string `xml:",cdata"`
}

type MediaFile struct {
	ID                  string `xml:"id,attr,omitempty"`
	Delivery            string `xml:"delivery,attr"`
	Type                string `xml:"type,attr"`
	Codec               string `xml:"codec,attr,omitempty"`
	Bitrate             int    `xml:"bitrate,attr,omitempty"`
	MinBitrate          int    `xml:"minBitrate,attr,omitempty"`
	MaxBitrate          int    `xml:"maxBitrate,attr,omitempty"`
	Width               int    `xml:"width,attr"`
	Height              int    `xml:"height,attr"`
	Scalable            bool   `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio bool   `xml:"maintainAspectRatio,attr,omitempty"`
	APIFramework        string `xml:"apiFramework,attr,omitempty"`
	URI                 string `xml:",cdata"`
}

type UniversalAdID struct {
	IDRegistry string `xml:"idRegistry,attr"`
	IDValue    string `xml:"idValue,attr"`
	ID         string `xml:",cdata"`
}
