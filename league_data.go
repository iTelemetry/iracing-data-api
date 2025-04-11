package irdata

import "time"

type League struct {
	LeagueID           int            `json:"league_id"`
	OwnerID            int            `json:"owner_id"`
	LeagueName         string         `json:"league_name"`
	Created            time.Time      `json:"created"`
	Hidden             bool           `json:"hidden"`
	Message            string         `json:"message"`
	About              string         `json:"about"`
	URL                string         `json:"url"`
	Recruiting         bool           `json:"recruiting"`
	PrivateWall        bool           `json:"private_wall"`
	PrivateRoster      bool           `json:"private_roster"`
	PrivateSchedule    bool           `json:"private_schedule"`
	PrivateResults     bool           `json:"private_results"`
	IsOwner            bool           `json:"is_owner"`
	IsAdmin            bool           `json:"is_admin"`
	RosterCount        int            `json:"roster_count"`
	Owner              LeagueMember   `json:"owner"`
	Image              LeagueImage    `json:"image"`
	Tags               LeagueTags     `json:"tags"`
	LeagueApplications []any          `json:"league_applications"`
	PendingRequests    []any          `json:"pending_requests"`
	IsMember           bool           `json:"is_member"`
	IsApplicant        bool           `json:"is_applicant"`
	IsInvite           bool           `json:"is_invite"`
	IsIgnored          bool           `json:"is_ignored"`
	Roster             []LeagueMember `json:"roster"`
}

type LeagueSeasons struct {
	Subscribed bool           `json:"subscribed"`
	Seasons    []LeagueSeason `json:"seasons"`
	Success    bool           `json:"success"`
	Retired    bool           `json:"retired"`
	LeagueID   int            `json:"league_id"`
}

type LeagueSeason struct {
	LeagueID                int                    `json:"league_id"`
	SeasonID                int                    `json:"season_id"`
	PointsSystemID          int                    `json:"points_system_id"`
	SeasonName              string                 `json:"season_name"`
	Active                  bool                   `json:"active"`
	Hidden                  bool                   `json:"hidden"`
	NumDrops                int                    `json:"num_drops"`
	NoDropsOnOrAfterRaceNum int                    `json:"no_drops_on_or_after_race_num"`
	PointsCars              []LeagueSeasonCar      `json:"points_cars"`
	DriverPointsCarClasses  []LeagueSeasonCarClass `json:"driver_points_car_classes"`
	TeamPointsCarClasses    []LeagueSeasonCarClass `json:"team_points_car_classes"`
	PointsSystemName        string                 `json:"points_system_name"`
	PointsSystemDesc        string                 `json:"points_system_desc"`
}

type LeagueSeasonCar struct {
	CarID   int    `json:"car_id"`
	CarName string `json:"car_name"`
}

type LeagueSeasonCarClass struct {
	CarClassID  int               `json:"car_class_id"`
	Name        string            `json:"name"`
	CarsInClass []LeagueSeasonCar `json:"cars_in_class"`
}

type LeagueMember struct {
	CustID            int       `json:"cust_id"`
	DisplayName       string    `json:"display_name"`
	Helmet            Helmet    `json:"helmet"`
	Owner             bool      `json:"owner,omitempty"`
	Admin             bool      `json:"admin,omitempty"`
	LeagueMailOptOut  bool      `json:"league_mail_opt_out,omitempty"`
	LeaguePmOptOut    bool      `json:"league_pm_opt_out,omitempty"`
	LeagueMemberSince time.Time `json:"league_member_since,omitempty"`
	CarNumber         any       `json:"car_number"`
	NickName          any       `json:"nick_name"`
}

type Helmet struct {
	Pattern    int    `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int    `json:"face_type"`
	HelmetType int    `json:"helmet_type"`
}

type LeagueImage struct {
	SmallLogo string `json:"small_logo"`
	LargeLogo string `json:"large_logo"`
}

type LeagueTags struct {
	Categorized    []any `json:"categorized"`
	NotCategorized []any `json:"not_categorized"`
}

type LeagueSeasonSessions struct {
	Sessions []LeagueSeasonSession `json:"sessions"`
}

type LeagueSeasonSession struct {
	Cars              []LeagueSeasonSessionCar `json:"cars"`
	DriverChanges     bool                     `json:"driver_changes"`
	EntryCount        int                      `json:"entry_count"`
	HasResults        bool                     `json:"has_results"`
	LaunchAt          time.Time                `json:"launch_at"`
	LeagueID          int                      `json:"league_id"`
	LeagueSeasonID    int                      `json:"league_season_id"`
	LoneQualify       bool                     `json:"lone_qualify"`
	PaceCarClassID    any                      `json:"pace_car_class_id"`
	PaceCarID         any                      `json:"pace_car_id"`
	PasswordProtected bool                     `json:"password_protected"`
	PracticeLength    int                      `json:"practice_length"`
	PrivateSessionID  int                      `json:"private_session_id"`
	QualifyLaps       int                      `json:"qualify_laps"`
	QualifyLength     int                      `json:"qualify_length"`
	RaceLaps          int                      `json:"race_laps"`
	RaceLength        int                      `json:"race_length"`
	SessionID         int                      `json:"session_id"`
	SubsessionID      int                      `json:"subsession_id"`
	Status            int                      `json:"status"`
	TrackID           int                      `json:"track_id"`
	TrackName         string                   `json:"track_name"`
}

type LeagueSeasonSessionCar struct {
	CarID        int    `json:"car_id"`
	CarName      string `json:"car_name"`
	CarClassID   int    `json:"car_class_id"`
	CarClassName string `json:"car_class_name"`
}
