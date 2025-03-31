package routes

import (
    "net/http"
    "FCM/infra/controllers"
)

// InitializeRoutes configura las rutas de la API.
func InitializeRoutes(notifController *controllers.NotificacionController) {
    http.HandleFunc("/enviarEstado", notifController.EnviarEstado)
}
