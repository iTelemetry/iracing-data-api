package irdata

type HostedSessions struct {
	Subscribed bool            `json:"subscribed"`
	Sessions   []HostedSession `json:"sessions"`
	Success    bool            `json:"success"`
}

type HostedSession struct {
	Session
	Admins []SessionDriver `json:"admins"`
}
