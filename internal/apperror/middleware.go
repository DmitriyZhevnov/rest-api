package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func MiddleWare(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *appError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appErr) {
				if errors.Is(err, errNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(appErr.Marshal())
					return
				}
				if errors.Is(err, internalServerError) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(appErr.Marshal())
					return
				}
				if errors.Is(err, unprocessableEntity) {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write(appErr.Marshal())
					return
				}
				if errors.Is(err, badRequestError) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(appErr.Marshal())
					return
				}

				err = err.(*appError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errNotFound.Marshal())
			}

			w.WriteHeader(http.StatusTeapot)
			w.Write(NewInternalServerError(err.Error(), "0000").Marshal())
		}
	}
}
