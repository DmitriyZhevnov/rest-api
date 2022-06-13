package handler

import (
	"net/http"

	"github.com/DmitriyZhevnov/rest-api/pkg/response"
)

const (
	authURL = "/authors"
)

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := h.services.Author.FindAll(r.Context())
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, authors)
	return nil
}
