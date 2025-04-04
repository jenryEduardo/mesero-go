package app

import "FCM/domain"

// NotificacionService implementa la lógica de negocio de las notificaciones.
type NotificacionService struct {
    NotificacionRepo domain.NotificacionPort
}

// NewNotificacionService crea una nueva instancia del servicio.
func NewNotificacionService(repo domain.NotificacionPort) *NotificacionService {
    return &NotificacionService{NotificacionRepo: repo}
}

// EnviarEstado envía una notificación con el estado del pedido.
func (s *NotificacionService) EnviarEstado(token string, status string) error {
    titulo := "Estado del Pedido"
    return s.NotificacionRepo.EnviarNotificacion(token, titulo, status)
}
