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
