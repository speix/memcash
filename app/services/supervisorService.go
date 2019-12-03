package services

import (
	"time"
)

func (e *Engine) Supervise() {

	ticker := time.NewTicker(e.Supervisor.Interval)

	for {
		select {
		case <-ticker.C:
			e.Purge()
		case <-e.Cache.Supervisor.Stop:
			ticker.Stop()
			return
		}
	}

}
