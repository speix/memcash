package app

import "time"

type Supervisor struct {
	Interval time.Duration
	Stop     chan bool
}

type SupervisorService interface {
	Supervise()
}
