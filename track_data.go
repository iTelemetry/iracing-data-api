package irdata

import "time"

type Track struct {
	AiEnabled              bool        `json:"ai_enabled"`
	AllowPitlaneCollisions bool        `json:"allow_pitlane_collisions"`
	AllowRollingStart      bool        `json:"allow_rolling_start"`
	AllowStandingStart     bool        `json:"allow_standing_start"`
	AwardExempt            bool        `json:"award_exempt"`
	Category               string      `json:"category"`
	CategoryID             int         `json:"category_id"`
	Closes                 string      `json:"closes"`
	ConfigName             string      `json:"config_name,omitempty"`
	CornersPerLap          int         `json:"corners_per_lap"`
	Created                time.Time   `json:"created"`
	FirstSale              time.Time   `json:"first_sale"`
	FreeWithSubscription   bool        `json:"free_with_subscription"`
	FullyLit               bool        `json:"fully_lit"`
	GridStalls             int         `json:"grid_stalls"`
	HasOptPath             bool        `json:"has_opt_path"`
	HasShortParadeLap      bool        `json:"has_short_parade_lap"`
	HasStartZone           bool        `json:"has_start_zone"`
	HasSvgMap              bool        `json:"has_svg_map"`
	IsDirt                 bool        `json:"is_dirt"`
	IsOval                 bool        `json:"is_oval"`
	IsPsPurchasable        bool        `json:"is_ps_purchasable"`
	LapScoring             int         `json:"lap_scoring"`
	Latitude               float64     `json:"latitude"`
	Location               string      `json:"location"`
	Longitude              float64     `json:"longitude"`
	MaxCars                int         `json:"max_cars"`
	NightLighting          bool        `json:"night_lighting"`
	NominalLapTime         float64     `json:"nominal_lap_time"`
	NumberPitStalls        int         `json:"number_pitstalls"`
	Opens                  string      `json:"opens"`
	PackageID              int         `json:"package_id"`
	PitRoadSpeedLimit      int         `json:"pit_road_speed_limit,omitempty"`
	Price                  float64     `json:"price"`
	PriceDisplay           string      `json:"price_display,omitempty"`
	Priority               int         `json:"priority"`
	Purchasable            bool        `json:"purchasable"`
	QualifyLaps            int         `json:"qualify_laps"`
	RestartOnLeft          bool        `json:"restart_on_left"`
	Retired                bool        `json:"retired"`
	SearchFilters          string      `json:"search_filters"`
	SiteURL                string      `json:"site_url,omitempty"`
	Sku                    int         `json:"sku"`
	SoloLaps               int         `json:"solo_laps"`
	StartOnLeft            bool        `json:"start_on_left"`
	SupportsGripCompound   bool        `json:"supports_grip_compound"`
	TechTrack              bool        `json:"tech_track"`
	TimeZone               string      `json:"time_zone"`
	TrackConfigLength      float64     `json:"track_config_length"`
	TrackDirPath           string      `json:"track_dirpath"`
	TrackID                int         `json:"track_id"`
	TrackName              string      `json:"track_name"`
	TrackTypes             []TrackType `json:"track_types"`
	Banking                string      `json:"banking,omitempty"`
}

type TrackType struct {
	TrackType string `json:"track_type"`
}

type TrackAssets struct {
	Coordinates     string            `json:"coordinates"`
	Detail          *string           `json:"detail_copy"`
	DetailTechSpecs *string           `json:"detail_tech_specs_copy"`
	DetailVideo     *string           `json:"detail_video"`
	Folder          string            `json:"folder"`
	GalleryImages   string            `json:"gallery_images"`
	GalleryPrefix   string            `json:"gallery_prefix"`
	LargeImage      string            `json:"large_image"`
	SmallImage      string            `json:"small_image"`
	Logo            string            `json:"logo"`
	North           string            `json:"north"`
	NumSvgImages    int               `json:"num_svg_images"`
	TrackID         int               `json:"track_id"`
	TrackMap        string            `json:"track_map"`
	TrackMapLayers  map[string]string `json:"track_map_layers"`
}
