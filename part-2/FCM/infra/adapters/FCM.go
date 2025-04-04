package adapters

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"
    "bytes"
)

// FCMService implementa NotificacionPort y se comunica con Firebase Cloud Messaging.
type FCMService struct {
    ServerKey string
}

// NewFCMService crea una nueva instancia del servicio FCM.
func NewFCMService(serverKey string) *FCMService {
    return &FCMService{ServerKey: serverKey}
}

// EnviarNotificacion envía un mensaje a un cliente FCM.
func (f *FCMService) EnviarNotificacion(token string, titulo string, mensaje string) error {
    fcmURL := "https://fcm.googleapis.com/fcm/send"

    // Estructura del mensaje FCM
    notification := map[string]interface{}{
        "to": token,
        "notification": map[string]string{
            "title": titulo,
            "body":  mensaje,
        },
    }

    jsonData, err := json.Marshal(notification)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "key="+f.ServerKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("error al enviar notificación FCM")
    }

    log.Println("Notificación enviada con éxito")
    return nil
}
