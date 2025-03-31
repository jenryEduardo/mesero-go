package main

import (
    "log"
    "net/http"
    "FCM/app"
    "FCM/infra/adapters"
    "FCM/infra/controllers"
    "FCM/infra/routes"
)

func main() {
    fcmService := adapters.NewFCMService("BOxRJG6C-O6RRXga_NgwPxDDTwxYHRGmhyo8irHGRAn_i0Vax50i0ynIsn6eanT2wkWQfrCGu_eZGxJXzY4AoFM")
    notificacionService := app.NewNotificacionService(fcmService)
    notifController := controllers.NewNotificacionController(notificacionService)

    routes.InitializeRoutes(notifController)

    log.Println("Servidor iniciado en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
