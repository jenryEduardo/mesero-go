package controllers

import (
	"api-main/users/application"
	"encoding/json"
	"net/http"
	"fmt"
)

// GetUserByEmailHandler es el controlador que maneja las solicitudes para obtener un usuario por su email
func GetUserByEmailHandler(getUserByEmailUseCase *application.GetUserByEmail) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el email de la URL o parámetros de la consulta
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "El email es requerido", http.StatusBadRequest)
			return
		}

		// Ejecutar el caso de uso de obtener el usuario por email
		user, err := getUserByEmailUseCase.Execute(email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error obteniendo usuario: %v", err), http.StatusNotFound)
			return
		}

		// Responder con los detalles del usuario
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
