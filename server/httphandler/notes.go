package httphandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	art "github.com/standielpls/articulate/server"
)

type request struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	note, err := parseCreateNotes(r.Body)
	if err != nil {
		h.writeError(w, err.Error(), 400)
		return
	}

	id, err := h.NoteStore.CreateNote(ctx, note)
	if err != nil {
		h.writeError(w, err.Error(), 500)
		return
	}

	data := response{
		ID: id,
		Ok: true,
	}

	h.writeData(w, data, 200)
}

type listRes struct {
	Articles []art.Article `json:"articles"`
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := r.URL.Query()

	userID := v.Get("user_id")
	if userID == "" {
		h.writeError(w, "user_id must be provided", 400)
		return
	}

	arts, err := h.NoteStore.ListByUserID(ctx, userID)
	if err != nil {
		h.writeError(w, err.Error(), 500)
		return
	}
	if len(arts) < 1 {
		arts = []art.Article{}
	}

	data := listRes{
		Articles: arts,
	}

	h.writeData(w, data, 200)
}

func parseCreateNotes(r io.Reader) (art.Article, error) {
	var a art.Article
	err := json.NewDecoder(r).Decode(&a)
	if err != nil {
		return a, err
	}

	if a.URL == "" {
		return a, errors.New("url must be provided")
	}
	if a.Article == "" {
		return a, errors.New("article must be provided")
	}

	if a.Comment == "" {
		return a, errors.New("comment must be provided")
	}

	if a.UserID == "" {
		return a, errors.New("user_id must be provided")
	}

	return a, nil
}
