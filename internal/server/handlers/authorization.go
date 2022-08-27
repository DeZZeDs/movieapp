package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"movieapp/internal/errors"
	"movieapp/pkg/database"
	"movieapp/pkg/logger"
	"net/http"
)

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")

	logger.Trace("userEmail:%s, userPassword: %s", userEmail, userPassword)

	user, err := database.GetUser(userEmail)
	if err != nil {
		logger.Error("Authorization - Error getting data from database: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		errors.WriteErrorInResponse(w, 1, errors.Errors[1], err.Error())
		return
	}

	if err = CompareHashAndPassword(user.Password, userPassword); err != nil {
		logger.Error("Authorization - Error compare hash and password: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errors.WriteErrorInResponse(w, 7, errors.Errors[7], err.Error())
		return
	}

	//Выдаю токен

	logger.Trace("Чел авторизован: %s", user.Email)
	w.WriteHeader(http.StatusOK)
}

func CompareHashAndPassword(hash string, password string) error {
	incoming := []byte(password)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
