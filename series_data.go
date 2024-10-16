package irdata

import "time"

type SeriesAssets map[string]SeriesAsset

type SeriesAsset struct {
	LargeImage string `json:"large_image"`
	Logo       string `json:"logo"`
	SeriesCopy string `json:"series_copy"`
	SeriesID   int    `json:"series_id"`
	SmallImage string `json:"small_image"`
}

type Series []struct {
	AllowedLicenses []AllowedLicense `json:"allowed_licenses"`
	Category        string           `json:"category"`
	CategoryID      int              `json:"category_id"`
	Eligible        bool             `json:"eligible"`
	ForumURL        string           `json:"forum_url,omitempty"`
	MaxStarters     int              `json:"max_starters"`
	MinStarters     int              `json:"min_starters"`
	OvalCautionType int              `json:"oval_caution_type"`
	RoadCautionType int              `json:"road_caution_type"`
	SeriesID        int              `json:"series_id"`
	SeriesName      string           `json:"series_name"`
	SeriesShortName string           `json:"series_short_name"`
	SearchFilters   string           `json:"search_filters,omitempty"`
}

type AllowedLicense struct {
	LicenseGroup    int    `json:"license_group"`
	MinLicenseLevel int    `json:"min_license_level"`
	MaxLicenseLevel int    `json:"max_license_level"`
	GroupName       string `json:"group_name"`
}

type SeriesSeasons []SeriesSeason

type SeriesSeason struct {
	SeasonID                   int                    `json:"season_id"`
	SeasonName                 string                 `json:"season_name"`
	Active                     bool                   `json:"active"`
	AllowedSeasonMembers       any                    `json:"allowed_season_members"`
	CarClassIds                []int                  `json:"car_class_ids"`
	CarSwitching               bool                   `json:"car_switching"`
	CarTypes                   []CarType              `json:"car_types"`
	CautionLapsDoNotCount      bool                   `json:"caution_laps_do_not_count"`
	Complete                   bool                   `json:"complete"`
	CrossLicense               bool                   `json:"cross_license"`
	DriverChangeRule           int                    `json:"driver_change_rule"`
	DriverChanges              bool                   `json:"driver_changes"`
	Drops                      int                    `json:"drops"`
	EnablePitlaneCollisions    bool                   `json:"enable_pitlane_collisions"`
	FixedSetup                 bool                   `json:"fixed_setup"`
	GreenWhiteCheckeredLimit   int                    `json:"green_white_checkered_limit"`
	GridByClass                bool                   `json:"grid_by_class"`
	HardcoreLevel              int                    `json:"hardcore_level"`
	HasSupersessions           bool                   `json:"has_supersessions"`
	IgnoreLicenseForPractice   bool                   `json:"ignore_license_for_practice"`
	IncidentLimit              int                    `json:"incident_limit"`
	IncidentWarnMode           int                    `json:"incident_warn_mode"`
	IncidentWarnParam1         int                    `json:"incident_warn_param1"`
	IncidentWarnParam2         int                    `json:"incident_warn_param2"`
	IsHeatRacing               bool                   `json:"is_heat_racing"`
	LicenseGroup               int                    `json:"license_group"`
	LicenseGroupTypes          []LicenseGroupType     `json:"license_group_types"`
	LuckyDog                   bool                   `json:"lucky_dog"`
	MaxTeamDrivers             int                    `json:"max_team_drivers"`
	MaxWeeks                   int                    `json:"max_weeks"`
	MinTeamDrivers             int                    `json:"min_team_drivers"`
	Multiclass                 bool                   `json:"multiclass"`
	MustUseDiffTireTypesInRace bool                   `json:"must_use_diff_tire_types_in_race"`
	NextRaceSession            any                    `json:"next_race_session"`
	NumOptLaps                 int                    `json:"num_opt_laps"`
	Official                   bool                   `json:"official"`
	OpDuration                 int                    `json:"op_duration"`
	OpenPracticeSessionTypeID  int                    `json:"open_practice_session_type_id"`
	QualifierMustStartRace     bool                   `json:"qualifier_must_start_race"`
	RaceWeek                   int                    `json:"race_week"`
	RaceWeekToMakeDivisions    int                    `json:"race_week_to_make_divisions"`
	RegOpenMinutes             int                    `json:"reg_open_minutes,omitempty"`
	RegUserCount               int                    `json:"reg_user_count"`
	RegionCompetition          bool                   `json:"region_competition"`
	RestrictByMember           bool                   `json:"restrict_by_member"`
	RestrictToCar              bool                   `json:"restrict_to_car"`
	RestrictViewing            bool                   `json:"restrict_viewing"`
	ScheduleDescription        string                 `json:"schedule_description"`
	Schedules                  []SeriesSeasonSchedule `json:"schedules"`
	SeasonQuarter              int                    `json:"season_quarter"`
	SeasonShortName            string                 `json:"season_short_name"`
	SeasonYear                 int                    `json:"season_year"`
	SendToOpenPractice         bool                   `json:"send_to_open_practice"`
	SeriesID                   int                    `json:"series_id"`
	ShortParadeLap             bool                   `json:"short_parade_lap"`
	StartDate                  time.Time              `json:"start_date"`
	StartOnQualTire            bool                   `json:"start_on_qual_tire"`
	StartZone                  bool                   `json:"start_zone"`
	TrackTypes                 []TrackType            `json:"track_types"`
	UnsportConductRuleMode     int                    `json:"unsport_conduct_rule_mode"`
	RacePoints                 int                    `json:"race_points,omitempty"`
	RookieSeason               string                 `json:"rookie_season,omitempty"`
	HeatSesInfo                HeatSessionInfo        `json:"heat_ses_info,omitempty"`
}

type CarTypes struct {
	CarType string `json:"car_type"`
}

type LicenseGroupType struct {
	LicenseGroupType int `json:"license_group_type"`
}

type RaceTimeDescriptor struct {
	Repeating      bool        `json:"repeating"`
	SessionMinutes int         `json:"session_minutes"`
	SessionTimes   []time.Time `json:"session_times"`
	SuperSession   bool        `json:"super_session"`
}

type SeriesSeasonTrack struct {
	Category   string `json:"category"`
	CategoryID int    `json:"category_id"`
	ConfigName string `json:"config_name"`
	TrackID    int    `json:"track_id"`
	TrackName  string `json:"track_name"`
}

type TrackState struct {
	LeaveMarbles bool `json:"leave_marbles"`
}

type ForecastOptions struct {
	ForecastType  int   `json:"forecast_type"`
	Precipitation int   `json:"precipitation"`
	Skies         int   `json:"skies"`
	StopPrecip    int   `json:"stop_precip"`
	Temperature   int   `json:"temperature"`
	WeatherSeed   int64 `json:"weather_seed"`
	WindDir       int   `json:"wind_dir"`
	WindSpeed     int   `json:"wind_speed"`
}

type WeatherSummary struct {
	MaxPrecipRate     float64 `json:"max_precip_rate"`
	MaxPrecipRateDesc string  `json:"max_precip_rate_desc"`
	PrecipChance      float64 `json:"precip_chance"`
	SkiesHigh         int     `json:"skies_high"`
	SkiesLow          int     `json:"skies_low"`
	TempHigh          float64 `json:"temp_high"`
	TempLow           float64 `json:"temp_low"`
	TempUnits         int     `json:"temp_units"`
	WindHigh          float64 `json:"wind_high"`
	WindLow           float64 `json:"wind_low"`
	WindUnits         int     `json:"wind_units"`
}

type Weather struct {
	AllowFog                bool            `json:"allow_fog"`
	Fog                     int             `json:"fog"`
	ForecastOptions         ForecastOptions `json:"forecast_options"`
	PrecipOption            int             `json:"precip_option"`
	RelHumidity             int             `json:"rel_humidity"`
	SimulatedStartTime      string          `json:"simulated_start_time"`
	SimulatedStartUtcTime   time.Time       `json:"simulated_start_utc_time"`
	SimulatedTimeMultiplier int             `json:"simulated_time_multiplier"`
	SimulatedTimeOffsets    []int           `json:"simulated_time_offsets"`
	Skies                   int             `json:"skies"`
	TempUnits               int             `json:"temp_units"`
	TempValue               int             `json:"temp_value"`
	TimeOfDay               int             `json:"time_of_day"`
	Type                    int             `json:"type"`
	Version                 int             `json:"version"`
	WeatherSummary          WeatherSummary  `json:"weather_summary"`
	WeatherURL              string          `json:"weather_url"`
	WeatherVarInitial       int             `json:"weather_var_initial"`
	WeatherVarOngoing       int             `json:"weather_var_ongoing"`
	WindDir                 int             `json:"wind_dir"`
	WindUnits               int             `json:"wind_units"`
	WindValue               int             `json:"wind_value"`
}

type SeriesSeasonSchedule struct {
	SeasonID                int                  `json:"season_id"`
	RaceWeekNum             int                  `json:"race_week_num"`
	CarRestrictions         []any                `json:"car_restrictions"`
	Category                string               `json:"category"`
	CategoryID              int                  `json:"category_id"`
	EnablePitlaneCollisions bool                 `json:"enable_pitlane_collisions"`
	FullCourseCautions      bool                 `json:"full_course_cautions"`
	QualAttached            bool                 `json:"qual_attached"`
	RaceLapLimit            int                  `json:"race_lap_limit"`
	RaceTimeLimit           int                  `json:"race_time_limit"` // minutes if > 0
	RaceTimeDescriptors     []RaceTimeDescriptor `json:"race_time_descriptors"`
	RaceWeekCars            []any                `json:"race_week_cars"`
	RestartType             string               `json:"restart_type"`
	ScheduleName            string               `json:"schedule_name"`
	SeasonName              string               `json:"season_name"`
	SeriesID                int                  `json:"series_id"`
	SeriesName              string               `json:"series_name"`
	ShortParadeLap          bool                 `json:"short_parade_lap"`
	SimulatedTimeMultiplier int                  `json:"simulated_time_multiplier"`
	SpecialEventType        any                  `json:"special_event_type"`
	StartDate               string               `json:"start_date"`
	StartType               string               `json:"start_type"`
	StartZone               bool                 `json:"start_zone"`
	Track                   Track                `json:"track"`
	TrackState              TrackState           `json:"track_state"`
	Weather                 Weather              `json:"weather"`
}

type HeatSessionInfo struct {
	ConsolationDeltaMaxFieldSize         int       `json:"consolation_delta_max_field_size"`
	ConsolationDeltaSessionLaps          int       `json:"consolation_delta_session_laps"`
	ConsolationDeltaSessionLengthMinutes int       `json:"consolation_delta_session_length_minutes"`
	ConsolationFirstMaxFieldSize         int       `json:"consolation_first_max_field_size"`
	ConsolationFirstSessionLaps          int       `json:"consolation_first_session_laps"`
	ConsolationFirstSessionLengthMinutes int       `json:"consolation_first_session_length_minutes"`
	ConsolationNumPositionToInvert       int       `json:"consolation_num_position_to_invert"`
	ConsolationNumToConsolation          int       `json:"consolation_num_to_consolation"`
	ConsolationNumToMain                 int       `json:"consolation_num_to_main"`
	ConsolationRunAlways                 bool      `json:"consolation_run_always"`
	ConsolationScoresChampPoints         bool      `json:"consolation_scores_champ_points"`
	Created                              time.Time `json:"created"`
	CustID                               int       `json:"cust_id"`
	Description                          string    `json:"description"`
	HeatCautionType                      int       `json:"heat_caution_type"`
	HeatInfoID                           int       `json:"heat_info_id"`
	HeatInfoName                         string    `json:"heat_info_name"`
	HeatLaps                             int       `json:"heat_laps"`
	HeatLengthMinutes                    int       `json:"heat_length_minutes"`
	HeatMaxFieldSize                     int       `json:"heat_max_field_size"`
	HeatNumFromEachToMain                int       `json:"heat_num_from_each_to_main"`
	HeatNumPositionToInvert              int       `json:"heat_num_position_to_invert"`
	HeatScoresChampPoints                bool      `json:"heat_scores_champ_points"`
	HeatSessionMinutesEstimate           int       `json:"heat_session_minutes_estimate"`
	Hidden                               bool      `json:"hidden"`
	MainLaps                             int       `json:"main_laps"`
	MainLengthMinutes                    int       `json:"main_length_minutes"`
	MainMaxFieldSize                     int       `json:"main_max_field_size"`
	MainNumPositionToInvert              int       `json:"main_num_position_to_invert"`
	MaxEntrants                          int       `json:"max_entrants"`
	OpenPractice                         bool      `json:"open_practice"`
	PreMainPracticeLengthMinutes         int       `json:"pre_main_practice_length_minutes"`
	PreQualNumToMain                     int       `json:"pre_qual_num_to_main"`
	PreQualPracticeLengthMinutes         int       `json:"pre_qual_practice_length_minutes"`
	QualCautionType                      int       `json:"qual_caution_type"`
	QualLaps                             int       `json:"qual_laps"`
	QualLengthMinutes                    int       `json:"qual_length_minutes"`
	QualNumToMain                        int       `json:"qual_num_to_main"`
	QualOpenDelaySeconds                 int       `json:"qual_open_delay_seconds"`
	QualScoresChampPoints                bool      `json:"qual_scores_champ_points"`
	QualScoring                          int       `json:"qual_scoring"`
	QualStyle                            int       `json:"qual_style"`
	RaceStyle                            int       `json:"race_style"`
}
