package model

import (
	"fmt"
	"regexp"
	"strings"
)

var idPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

func ValidateID(id string) error {
	if len(id) < 3 || len(id) > 64 {
		return fmt.Errorf("id length")
	}
	if !idPattern.MatchString(id) {
		return fmt.Errorf("id format")
	}
	return nil
}
func NormalizeName(v string) string { return strings.Join(strings.Fields(strings.TrimSpace(v)), " ") }
func Statuses() []string            { return []string{"draft", "published", "archived"} }
func IsTerminal(status string) bool { return status == "archived" }
func CanTransition(from, to string) bool {
	switch from {
	case "draft":
		return to == "published"
	case "published":
		return to == "archived"
	case "archived":
		return false
	default:
		return false
	}
}
func EnsureTransition(from, to string) error {
	if !CanTransition(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	return nil
}
func ClampStock(v int) int {
	if v < 0 {
		return 0
	}
	if v > 100000 {
		return 100000
	}
	return v
}
func PriceBand(v float64) string {
	if v < 20 {
		return "entry"
	}
	if v < 100 {
		return "standard"
	}
	if v < 500 {
		return "premium"
	}
	return "professional"
}
func Searchable(r Record) string {
	return strings.ToLower(strings.Join([]string{r.Name, r.Category, r.Material, r.Status}, " "))
}
