package workflow

import (
	"coffeeware/model"
	"sort"
	"time"
)

type TaskReport struct {
	ID, State         string
	Started, Finished time.Time
	Children          int
}

func BuildReport(id, state string, children int) TaskReport {
	return TaskReport{ID: id, State: state, Started: time.Now().UTC(), Children: children}
}
func SortReports(rs []TaskReport) []TaskReport {
	out := append([]TaskReport(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].Started.Before(out[j].Started) })
	return out
}
func IsSuccessful(r TaskReport) bool { return r.State == "completed" && r.Children > 0 }
func FailureReason(r TaskReport) string {
	if r.State == "cancelled" {
		return "operator cancelled"
	}
	if r.Children == 0 {
		return "no child inspections"
	}
	return ""
}
func EventForReport(r TaskReport) model.Event {
	return model.NewEvent(r.ID, r.ID, "report", "system", r.State)
}
