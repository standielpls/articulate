package httphandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	art "github.com/standielpls/articulate/server"
)

type Handler struct {
	NoteStore art.NoteStore
	Version   string
}

func (h *Handler) Handler() http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").Name("root").HandlerFunc(h.VersionHandler)
	r.Methods("GET").Path("/version").Name("version").HandlerFunc(h.VersionHandler)

	v1 := r.NewRoute().PathPrefix("/v1").Subrouter()
	v1.Methods("GET").Path("/notes").Name("list_notes").HandlerFunc(h.ListNotes)
	v1.Methods("POST").Path("/notes").Name("create_note").HandlerFunc(h.CreateNote)

	return r
}

func (h *Handler) VersionHandler(w http.ResponseWriter, r *http.Request) {
	ver := struct {
		Version   string    `json:"version"`
		Timestamp time.Time `json:"timestamp"`
		Service   string    `json:"service"`
	}{
		Version:   h.Version,
		Timestamp: time.Now(),
		Service:   "articulate",
	}
	h.writeData(w, ver, 200)
}

type dataResp struct {
	Data interface{} `json:"data"`
}

func (h *Handler) writeData(w http.ResponseWriter, data interface{}, code int) {
	d := dataResp{
		Data: data,
	}

	h.write(w, d, code)
}

type errMsg struct {
	Message string `json:"message"`
}
type errResp struct {
	Error errMsg `json:"error"`
}

func (h *Handler) write(w http.ResponseWriter, data interface{}, code int) {
	b, _ := json.Marshal(data)
	w.Write(b)
	w.WriteHeader(code)
}

func (h *Handler) writeError(w http.ResponseWriter, msg string, code int) {
	e := errResp{
		Error: errMsg{
			Message: msg,
		},
	}

	h.write(w, e, code)
}
