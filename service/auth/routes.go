package auth

import (
	"fmt"
	"github.com/Darthex/ink-golang/types/auth"
	"github.com/Darthex/ink-golang/utils"
	"net/http"
)

func Router(router *http.ServeMux, store *Store) {

	router.HandleFunc("POST /auth/register",
		func(w http.ResponseWriter, r *http.Request) {
			var payload auth.RegisterUserPayload
			if err := utils.ParseAndValidate(w, r, &payload); err != nil {
				return
			}
			if _, err := store.GetUserByEmail(payload.Email); err == nil {
				utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists"))
				return
			}
			hashedPassword, err := HashPassword(payload.Password)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to hash password: %v", err))
				return
			}
			if createErr := store.CreateNewUser(auth.User{
				Username: payload.Username,
				Email:    payload.Email,
				Password: hashedPassword,
			}); createErr != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusCreated, nil)
			return
		})

	router.HandleFunc("POST /auth/login",
		func(w http.ResponseWriter, r *http.Request) {
			var payload auth.LoginUserPayload
			if err := utils.ParseAndValidate(w, r, &payload); err != nil {
				return
			}
			u, err := store.GetUserByEmail(payload.Email)
			if err != nil {
				utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%v", err))
				return
			}
			passMatch := ComparePassword(u.Password, []byte(payload.Password))
			if !passMatch {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
				return
			}
			newToken, err := CreateJWT(u.ID)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create token: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, map[string]interface{}{"access_token": newToken, "token_type": "bearer", "user": u})
			return
		})

	router.HandleFunc("POST /auth/update",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context().Value("claims").(*auth.CustomClaims)
			var payload auth.UpdateUserPayload
			if err := utils.ParseAndValidate(w, r, &payload); err != nil {
				return
			}
			u, err := store.GetUserByID(ctx.UserID)
			if err != nil {
				utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%v", err))
				return
			}
			if err := store.UpdateUser(payload.Username, u.ID); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update user: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, nil)
			return
		})
}
