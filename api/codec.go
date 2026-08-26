package api

import (
	"coffeeware/model"
	"encoding/json"
	"net/http"
)

type recordInput struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Material string  `json:"material"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

func decodeRecord(r *http.Request) (model.Record, error) {
	var in recordInput
	if e := json.NewDecoder(r.Body).Decode(&in); e != nil {
		return model.Record{}, e
	}
	x := model.NewRecord(in.ID, in.Name, in.Category, in.Price, in.Stock)
	x.Material = in.Material
	return x, nil
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
func methodAllowed(w http.ResponseWriter, method string, allowed ...string) bool {
	for _, a := range allowed {
		if method == a {
			return true
		}
	}
	w.Header().Set("Allow", joinMethods(allowed))
	writeError(w, 405, "method not allowed")
	return false
}
func joinMethods(v []string) string {
	out := ""
	for i, x := range v {
		if i > 0 {
			out += ", "
		}
		out += x
	}
	return out
}
