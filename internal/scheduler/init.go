package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	servicependinguser "github.com/hend41234/startchat/internal/service/pending_user"
)

var Scheduler *gocron.Scheduler

func init() {
	var loc = "Asia/Jakarta" // edit the time location that you want
	timeloc, err := time.LoadLocation(loc)
	if err != nil {
		return
	}
	Scheduler = gocron.NewScheduler(timeloc)
	Scheduler.Every(2).Minute().Do(servicependinguser.Cleaner)
	fmt.Println("scheduler running")
	Scheduler.StartAsync()
}
