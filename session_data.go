package irdata

import "time"

type Session struct {
	AiAvoidPlayers             bool                      `json:"ai_avoid_players"`
	AiMaxSkill                 int                       `json:"ai_max_skill,omitempty"`
	AiMinSkill                 int                       `json:"ai_min_skill,omitempty"`
	AiRosterName               string                    `json:"ai_roster_name,omitempty"`
	AllowedClubs               []interface{}             `json:"allowed_clubs"`
	AllowedLeagues             []interface{}             `json:"allowed_leagues"`
	AllowedTeams               []interface{}             `json:"allowed_teams"`
	CarTypes                   []CarType                 `json:"car_types"`
	Cars                       []SessionCar              `json:"cars"`
	CarsLeft                   int                       `json:"cars_left"`
	ConsecCautionsSingleFile   bool                      `json:"consec_cautions_single_file"`
	CountByCarClassID          map[string]int            `json:"count_by_car_class_id"`
	CountByCarID               map[string]int            `json:"count_by_car_id"`
	DamageModel                int                       `json:"damage_model"`
	DisallowVirtualMirror      bool                      `json:"disallow_virtual_mirror"`
	DoNotCountCautionLaps      bool                      `json:"do_not_count_caution_laps"`
	DoNotPaintCars             bool                      `json:"do_not_paint_cars"`
	DriverChangeRule           int                       `json:"driver_change_rule"`
	DriverChanges              bool                      `json:"driver_changes"`
	Eligibility                SessionEligibility        `json:"elig"`
	EnablePitlaneCollisions    bool                      `json:"enable_pitlane_collisions"`
	EntryCount                 int                       `json:"entry_count"`
	EventTypes                 []EventType               `json:"event_types"`
	Farm                       SessionServerFarm         `json:"farm"`
	FixedSetup                 bool                      `json:"fixed_setup"`
	FullCourseCautions         bool                      `json:"full_course_cautions"`
	GreenWhiteCheckeredLimit   int                       `json:"green_white_checkered_limit"`
	HardcoreLevel              int                       `json:"hardcore_level"`
	Host                       SessionDriver             `json:"host"`
	IncidentLimit              int                       `json:"incident_limit"`
	IncidentWarnMode           int                       `json:"incident_warn_mode"`
	IncidentWarnParam1         int                       `json:"incident_warn_param1"`
	IncidentWarnParam2         int                       `json:"incident_warn_param2"`
	LaunchAt                   time.Time                 `json:"launch_at"`
	LeagueID                   int                       `json:"league_id"`
	LeagueSeasonID             int                       `json:"league_season_id"`
	LicenseGroupTypes          []SessionLicenseGroupType `json:"license_group_types"`
	LoneQualify                bool                      `json:"lone_qualify"`
	LuckyDog                   bool                      `json:"lucky_dog"`
	MaxAiDrivers               int                       `json:"max_ai_drivers"`
	MaxDrivers                 int                       `json:"max_drivers"`
	MaxIr                      int                       `json:"max_ir"`
	MaxLicenseLevel            int                       `json:"max_license_level"`
	MaxTeamDrivers             int                       `json:"max_team_drivers"`
	MinIr                      int                       `json:"min_ir"`
	MinLicenseLevel            int                       `json:"min_license_level"`
	MinTeamDrivers             int                       `json:"min_team_drivers"`
	MulticlassType             int                       `json:"multiclass_type"`
	MustUseDiffTireTypesInRace bool                      `json:"must_use_diff_tire_types_in_race"`
	NoLapperWaveArounds        bool                      `json:"no_lapper_wave_arounds"`
	NumFastTows                int                       `json:"num_fast_tows"`
	NumOptLaps                 int                       `json:"num_opt_laps"`
	OpenRegExpires             time.Time                 `json:"open_reg_expires"`
	OrderID                    int                       `json:"order_id"`
	PaceCarClassID             int                       `json:"pace_car_class_id"`
	PaceCarID                  int                       `json:"pace_car_id"`
	PasswordProtected          bool                      `json:"password_protected"`
	PitsInUse                  int                       `json:"pits_in_use"`
	PracticeLength             int                       `json:"practice_length"`
	PrivateSessionID           int                       `json:"private_session_id"`
	QualifierMustStartRace     bool                      `json:"qualifier_must_start_race"`
	QualifyLaps                int                       `json:"qualify_laps"`
	QualifyLength              int                       `json:"qualify_length"`
	RaceLaps                   int                       `json:"race_laps"`
	RaceLength                 int                       `json:"race_length"`
	Restarts                   int                       `json:"restarts"`
	RestrictResults            bool                      `json:"restrict_results"`
	RestrictViewing            bool                      `json:"restrict_viewing"`
	RollingStarts              bool                      `json:"rolling_starts"`
	SessionFull                bool                      `json:"session_full"`
	SessionID                  int                       `json:"session_id"`
	SessionName                string                    `json:"session_name"`
	SessionType                int                       `json:"session_type"`
	SessionTypes               []SessionType             `json:"session_types"`
	ShortParadeLap             bool                      `json:"short_parade_lap"`
	StartOnQualTire            bool                      `json:"start_on_qual_tire"`
	StartZone                  bool                      `json:"start_zone"`
	Status                     int                       `json:"status"`
	SubsessionID               int                       `json:"subsession_id"`
	TeamEntryCount             int                       `json:"team_entry_count"`
	TelemetryForceToDisk       int                       `json:"telemetry_force_to_disk"`
	TelemetryRestriction       int                       `json:"telemetry_restriction"`
	TimeLimit                  int                       `json:"time_limit"`
	Track                      SessionTrack              `json:"track"`
	TrackState                 SessionTrackState         `json:"track_state"`
	TrackTypes                 []SessionTrackType        `json:"track_types"`
	UnsportConductRuleMode     int                       `json:"unsport_conduct_rule_mode"`
	WarmupLength               int                       `json:"warmup_length"`
	Weather                    SessionWeather            `json:"weather"`
	RegisteredTeams            []int                     `json:"registered_teams,omitempty"`
	SessionDesc                string                    `json:"session_desc,omitempty"`
	HeatSesInfo                SessionHeatInfo           `json:"heat_ses_info,omitempty"`
	AltAssetID                 int                       `json:"alt_asset_id,omitempty"`
}

type SessionDriver struct {
	CustID      int          `json:"cust_id"`
	DisplayName string       `json:"display_name"`
	Helmet      DriverHelmet `json:"helmet"`
}

type SessionCar struct {
	CarID           int     `json:"car_id"`
	CarName         string  `json:"car_name"`
	CarClassID      int     `json:"car_class_id"`
	CarClassName    string  `json:"car_class_name"`
	MaxPctFuelFill  int     `json:"max_pct_fuel_fill"`
	WeightPenaltyKg int     `json:"weight_penalty_kg"`
	PowerAdjustPct  float64 `json:"power_adjust_pct"`
	MaxDryTireSets  int     `json:"max_dry_tire_sets"`
	PackageID       int     `json:"package_id"`
}

type SessionEligibility struct {
	SessionFull     bool  `json:"session_full"`
	CanSpot         bool  `json:"can_spot"`
	CanWatch        bool  `json:"can_watch"`
	CanDrive        bool  `json:"can_drive"`
	HasSessPassword bool  `json:"has_sess_password"`
	NeedsPurchase   bool  `json:"needs_purchase"`
	OwnCar          bool  `json:"own_car"`
	OwnTrack        bool  `json:"own_track"`
	PurchaseSkus    []int `json:"purchase_skus"`
	Registered      bool  `json:"registered"`
}

type SessionEventType struct {
	EventType int `json:"event_type"`
}

type SessionServerFarm struct {
	FarmID      int    `json:"farm_id"`
	DisplayName string `json:"display_name"`
	ImagePath   string `json:"image_path"`
	Displayed   bool   `json:"displayed"`
}

type SessionLicenseGroupType struct {
	LicenseGroupType int `json:"license_group_type"`
}

type SessionType struct {
	SessionType int `json:"session_type"`
}

type SessionTrack struct {
	TrackID    int    `json:"track_id"`
	TrackName  string `json:"track_name"`
	ConfigName string `json:"config_name"`
}

type SessionTrackState struct {
	LeaveMarbles         bool `json:"leave_marbles"`
	PracticeRubber       int  `json:"practice_rubber"`
	QualifyRubber        int  `json:"qualify_rubber"`
	WarmupRubber         int  `json:"warmup_rubber"`
	RaceRubber           int  `json:"race_rubber"`
	PracticeGripCompound int  `json:"practice_grip_compound"`
	QualifyGripCompound  int  `json:"qualify_grip_compound"`
	WarmupGripCompound   int  `json:"warmup_grip_compound"`
	RaceGripCompound     int  `json:"race_grip_compound"`
}

type SessionTrackType struct {
	TrackType string `json:"track_type"`
}

type SessionWeather struct {
	Version                 int    `json:"version"`
	Type                    int    `json:"type"`
	TempUnits               int    `json:"temp_units"`
	TempValue               int    `json:"temp_value"`
	RelHumidity             int    `json:"rel_humidity"`
	Fog                     int    `json:"fog"`
	WindDir                 int    `json:"wind_dir"`
	WindUnits               int    `json:"wind_units"`
	WindValue               int    `json:"wind_value"`
	Skies                   int    `json:"skies"`
	WeatherVarInitial       int    `json:"weather_var_initial"`
	WeatherVarOngoing       int    `json:"weather_var_ongoing"`
	TimeOfDay               int    `json:"time_of_day"`
	SimulatedStartTime      string `json:"simulated_start_time"`
	SimulatedTimeOffsets    []int  `json:"simulated_time_offsets"`
	SimulatedTimeMultiplier int    `json:"simulated_time_multiplier"`
}

type SessionHeatInfo struct {
	HeatInfoID                           int       `json:"heat_info_id"`
	CustID                               int       `json:"cust_id"`
	Hidden                               bool      `json:"hidden"`
	Created                              time.Time `json:"created"`
	HeatInfoName                         string    `json:"heat_info_name"`
	MaxEntrants                          int       `json:"max_entrants"`
	RaceStyle                            int       `json:"race_style"`
	OpenPractice                         bool      `json:"open_practice"`
	PreQualPracticeLengthMinutes         int       `json:"pre_qual_practice_length_minutes"`
	PreQualNumToMain                     int       `json:"pre_qual_num_to_main"`
	QualStyle                            int       `json:"qual_style"`
	QualLengthMinutes                    int       `json:"qual_length_minutes"`
	QualLaps                             int       `json:"qual_laps"`
	QualNumToMain                        int       `json:"qual_num_to_main"`
	QualScoring                          int       `json:"qual_scoring"`
	QualCautionType                      int       `json:"qual_caution_type"`
	QualOpenDelaySeconds                 int       `json:"qual_open_delay_seconds"`
	QualScoresChampPoints                bool      `json:"qual_scores_champ_points"`
	HeatLengthMinutes                    int       `json:"heat_length_minutes"`
	HeatLaps                             int       `json:"heat_laps"`
	HeatMaxFieldSize                     int       `json:"heat_max_field_size"`
	HeatNumPositionToInvert              int       `json:"heat_num_position_to_invert"`
	HeatCautionType                      int       `json:"heat_caution_type"`
	HeatNumFromEachToMain                int       `json:"heat_num_from_each_to_main"`
	HeatScoresChampPoints                bool      `json:"heat_scores_champ_points"`
	ConsolationNumToConsolation          int       `json:"consolation_num_to_consolation"`
	ConsolationNumToMain                 int       `json:"consolation_num_to_main"`
	ConsolationFirstMaxFieldSize         int       `json:"consolation_first_max_field_size"`
	ConsolationFirstSessionLengthMinutes int       `json:"consolation_first_session_length_minutes"`
	ConsolationFirstSessionLaps          int       `json:"consolation_first_session_laps"`
	ConsolationDeltaMaxFieldSize         int       `json:"consolation_delta_max_field_size"`
	ConsolationDeltaSessionLengthMinutes int       `json:"consolation_delta_session_length_minutes"`
	ConsolationDeltaSessionLaps          int       `json:"consolation_delta_session_laps"`
	ConsolationNumPositionToInvert       int       `json:"consolation_num_position_to_invert"`
	ConsolationScoresChampPoints         bool      `json:"consolation_scores_champ_points"`
	ConsolationRunAlways                 bool      `json:"consolation_run_always"`
	PreMainPracticeLengthMinutes         int       `json:"pre_main_practice_length_minutes"`
	MainLengthMinutes                    int       `json:"main_length_minutes"`
	MainLaps                             int       `json:"main_laps"`
	MainMaxFieldSize                     int       `json:"main_max_field_size"`
	MainNumPositionToInvert              int       `json:"main_num_position_to_invert"`
	HeatSessionMinutesEstimate           int       `json:"heat_session_minutes_estimate"`
}
