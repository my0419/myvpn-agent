package installer

import "time"

type State struct {
	Status          Status `json:"status"`
	TimeStarting    time.Time `json:"time_starting"`
	TimeRunning 	int  `json:"time_running"`
}