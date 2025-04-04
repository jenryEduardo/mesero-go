package domain

// NotificacionPort define la interfaz para el servicio de notificaciones.
type NotificacionPort interface {
    EnviarNotificacion(token string, titulo string, mensaje string) error
}
