package user

import (
	"net/http"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/handlers"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{logger: logger}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.MiddleWare(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.MiddleWare(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.MiddleWare(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.MiddleWare(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.MiddleWare(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.MiddleWare(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))

	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(200)
	// w.Write([]byte("this is user with uuid"))

	return apperror.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this update user"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))

	return nil
}
