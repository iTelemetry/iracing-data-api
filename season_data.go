package irdata

import "time"

type Seasons struct {
	SeasonQuarter int      `json:"season_quarter"`
	Seasons       []Season `json:"seasons"`
	SeasonYear    int      `json:"season_year"`
}

type Season struct {
	SeasonID      int    `json:"season_id"`
	SeriesID      int    `json:"series_id"`
	SeasonName    string `json:"season_name"`
	SeriesName    string `json:"series_name"`
	Official      bool   `json:"official"`
	SeasonYear    int    `json:"season_year"`
	SeasonQuarter int    `json:"season_quarter"`
	LicenseGroup  int    `json:"license_group"`
	FixedSetup    bool   `json:"fixed_setup"`
	DriverChanges bool   `json:"driver_changes"`
	RookieSeason  string `json:"rookie_season,omitempty"`
}

type RaceGuide struct {
	Subscribed     bool               `json:"subscribed"`
	Sessions       []RaceGuideSession `json:"sessions"`
	BlockBeginTime time.Time          `json:"block_begin_time"`
	BlockEndTime   time.Time          `json:"block_end_time"`
	Success        bool               `json:"success"`
}

type RaceGuideSession struct {
	SeasonID     int       `json:"season_id"`
	StartTime    time.Time `json:"start_time"`
	SuperSession bool      `json:"super_session"`
	SeriesID     int       `json:"series_id"`
	RaceWeekNum  int       `json:"race_week_num"`
	EndTime      time.Time `json:"end_time"`
	SessionID    int       `json:"session_id,omitempty"`
	EntryCount   int       `json:"entry_count"`
}

type SpectatorSubSessionIDs struct {
	EventTypes    []int `json:"event_types"`
	Success       bool  `json:"success"`
	SubsessionIds []int `json:"subsession_ids"`
}
