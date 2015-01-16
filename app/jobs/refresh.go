package jobs

import (
	"time"

	"github.com/revel/revel"
	"github.com/revel/revel/modules/jobs/app/jobs"
	"github.com/uzimith/outcastify/app/controllers"
)

type Refresh struct{}

func (c Refresh) Run() {
	revel.INFO.Println("--- refresh ---")
	t := time.Now()
	t = t.AddDate(0, 0, -1)
	controllers.Gdb.Exec("DELETE from users WHERE updated_at < ?", t.Format("2006-01-02 15:04:05"))
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1day", Refresh{})
	})
}
