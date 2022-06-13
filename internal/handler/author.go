package handler

import (
	"net/http"

	"github.com/DmitriyZhevnov/rest-api/pkg/response"
	"github.com/julienschmidt/httprouter"
)

const (
	authorsURL = "/authors"
	authorURL  = "/authors/:uuid"
)

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := h.services.Author.FindAll(r.Context())
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, authors)
	return nil
}

func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	authorID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	author, err := h.services.FindAuthor(r.Context(), authorID)
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, author)
	return nil
}
