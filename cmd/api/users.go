package main

import (
	"net/http"

	"github.com/rizbo-dev/social-api/internal/store"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseIdFromUrl("userID", r)

	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.badRequestError(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}

}
