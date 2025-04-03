package dependencies

import (
	"iot/historial/application"
	"iot/core"
	"iot/historial/infraestructure"
	"iot/historial/infraestructure/controllers"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/streadway/amqp"
)

var (
	mySQL       infraestructure.MySQLRepository
	mqttClient  mqtt.Client
	rabbitConn  *amqp.Connection
)

// Init inicializa la base de datos, MQTT y RabbitMQ
func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		log.Fatalf("Error inicializando MySQL: %v", err)
	}

	mqttClient = initMQTT()
	rabbitConn = initRabbitMQ()

	mySQL = *infraestructure.NewMySQLRepository(db, mqttClient, rabbitConn)
}

// Inicializa el cliente MQTT
func initMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // Ajusta la URL según tu configuración
	opts.SetClientID("go-mqtt")
	opts.SetUsername("user")
	opts.SetPassword("password")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error conectando a MQTT: %v", token.Error())
	}
	return client
}

// Inicializa la conexión a RabbitMQ
func initRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	return conn
}

// FindColor retorna el controlador de FindColor
func FindColor() *controllers.FindColorCtrl {
	ucFindColor := application.NewFindColor(&mySQL)
	return controllers.NewFindIdCircuitoCtrl(ucFindColor)
}

// GetStatus retorna el controlador para escuchar MQTT
func GetStatus() *controllers.GetStatusCtrl {
	return controllers.NewGetStatusCtrl(&mySQL)
}

// UpdateStatus retorna el controlador para actualizar estados
func UpdateStatus() *controllers.UpdateStatusCtrl {
	return controllers.NewUpdateStatusCtrl(&mySQL)
}
