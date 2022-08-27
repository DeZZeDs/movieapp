package handlers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"movieapp/internal/errors"
	"movieapp/internal/models"
	"movieapp/pkg/database"
	"movieapp/pkg/logger"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("RegistrationHandler - parse json error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		errors.WriteErrorInResponse(w, 5, errors.Errors[5], err.Error())
		return
	}

	if validError := user.ValidationUser(); validError != nil {
		logger.Error("RegistrationHandler - Validation error: %s", errors.PrepareValidationErrorToString(validError))
		w.WriteHeader(http.StatusBadRequest)
		errors.WriteErrorInResponse(w, 6, errors.Errors[6], errors.PrepareValidationErrorToString(validError))
		return
	}

	hash, err := GenerateHash(user.Password)
	if err != nil {
		logger.Error("RegistrationHandler - Could not hash password")
		w.WriteHeader(http.StatusInternalServerError)
		errors.WriteErrorInResponse(w, 7, errors.Errors[7], err.Error())
		return
	}
	user.Password = hash

	if err := database.UserCreate(&user); err != nil {
		logger.Error("RegistrationHandler - Error adding data to database: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errors.WriteErrorInResponse(w, 2, errors.Errors[2], err.Error())
		return
	}
	logger.Trace("RegistrationHandler - Created new user: %s, %s, %s, %s", user.FirstName, user.LastName, user.Email, user.Password)
	w.WriteHeader(http.StatusOK)
}

func GenerateHash(password string) (string, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}
