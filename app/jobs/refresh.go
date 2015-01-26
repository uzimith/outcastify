package jobs

import (
	"time"

	"github.com/revel/revel"
	"github.com/revel/revel/modules/jobs/app/jobs"
	"github.com/uzimith/outcastify/app/controllers"
	"github.com/uzimith/outcastify/app/models"
)

type Refresh struct{}

func (c Refresh) Run() {
	revel.INFO.Println("--- refresh ---")
	t := time.Now()
	t = t.AddDate(0, 0, -1)
	controllers.Gdb.Unscoped().Where("updated_at < ?", t.Format("2006-01-02 15:04:05")).Delete(models.User{})
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@midnight", Refresh{})
	})
}
