package controllers

import (
    "encoding/json"
    "net/http"
    "FCM/app"
)

// NotificacionController maneja las notificaciones.
type NotificacionController struct {
    NotificacionService *app.NotificacionService
}

// NewNotificacionController crea un nuevo controlador.
func NewNotificacionController(service *app.NotificacionService) *NotificacionController {
    return &NotificacionController{NotificacionService: service}
}

// EnviarEstado recibe el estado de la API y lo envía a FCM.
func (c *NotificacionController) EnviarEstado(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var data struct {
        Token  string `json:"token"`
        Status string `json:"status"`
    }

    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Error en el JSON recibido", http.StatusBadRequest)
        return
    }

    err = c.NotificacionService.EnviarEstado(data.Token, data.Status)
    if err != nil {
        http.Error(w, "Error al enviar notificación", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Notificación enviada exitosamente"))
}
