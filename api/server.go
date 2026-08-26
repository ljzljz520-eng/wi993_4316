package api

import (
	"coffeeware/model"
	"coffeeware/service"
	"context"
	"encoding/json"
	"net/http"
)

type Server struct{ svc *service.Service }

func New(s *service.Service) *Server { return &Server{svc: s} }
func (s *Server) Routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/records", s.records)
	m.HandleFunc("/tasks", s.tasks)
	return m
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	if e := s.svc.Health(); e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	w.WriteHeader(204)
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var in model.Record
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			http.Error(w, "bad json", 400)
			return
		}
		if e := s.svc.Register(in); e != nil {
			http.Error(w, e.Error(), 422)
			return
		}
		w.WriteHeader(201)
	case "GET":
		out, e := s.svc.Query(r.URL.Query().Get("q"))
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(out)
	default:
		http.Error(w, "method", 405)
	}
}
func (s *Server) tasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method == "POST" {
		e := s.svc.Submit(r.Context(), id)
		if e != nil {
			http.Error(w, e.Error(), 409)
			return
		}
		w.WriteHeader(202)
		return
	}
	status, e := s.svc.Status(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": status})
}
func Run(ctx context.Context, addr string, svc *service.Service) error {
	srv := &http.Server{Addr: addr, Handler: New(svc).Routes()}
	go func() { <-ctx.Done(); srv.Shutdown(context.Background()) }()
	return srv.ListenAndServe()
}
