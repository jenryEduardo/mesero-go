package infraestructure

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MySQLRepository struct {
	db         *sql.DB
	MQTTClient mqtt.Client
}

func NewMySQLRepository(db *sql.DB, mqttClient mqtt.Client) *MySQLRepository {
	return &MySQLRepository{db: db, MQTTClient: mqttClient}
}

// FindColor busca el color a partir del idPedido y lo publica en MQTT
func (r *MySQLRepository) FindColor(idPedido int) (int, error) {
	var idMesa int
	var idCircuito int

	err := r.db.QueryRow("SELECT idMesa FROM pedido WHERE idPedido=?", idPedido).Scan(&idMesa)
	if err != nil {
		log.Println("Error al obtener idMesa:", err)
		return 0, err
	}

	err = r.db.QueryRow("SELECT color FROM circuito WHERE idMesa=?", idMesa).Scan(&idCircuito)
	if err != nil {
		log.Println("Error al obtener idCircuito:", err)
		return 0, err
	}

	// Publicar en MQTT
	topic := "esp32/circuito"
	message := fmt.Sprintf(`{"idPedido": %d, "idCircuito": %d}`, idPedido, idCircuito)
	r.publishToMQTT(topic, message)

	return idCircuito, nil
}

// GetStatus escucha los mensajes desde MQTT y procesa los estados
func (r *MySQLRepository) GetStatus() error {
	topic := "status/update"

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		var data map[string]interface{}
		err := json.Unmarshal(msg.Payload(), &data)
		if err != nil {
			log.Println("Error al procesar mensaje MQTT:", err)
			return
		}

		idPedido := int(data["idPedido"].(float64))
		status := data["status"].(string)

		// Llamar a UpdateStatus para actualizar el estado en la BD
		err = r.UpdateStatus(idPedido, status)
		if err != nil {
			log.Println("Error al actualizar estado:", err)
			return
		}
	}

	if token := r.MQTTClient.Subscribe(topic, 1, messageHandler); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error al suscribirse a MQTT: %v", token.Error())
	}

	log.Println("Escuchando status en MQTT:", topic)
	return nil
}

// UpdateStatus actualiza el status en la BD
func (r *MySQLRepository) UpdateStatus(idPedido int, nuevoStatus string) error {
	var idRobot int

	err := r.db.QueryRow("SELECT idRobot FROM historial_entrega WHERE idPedido = ?", idPedido).Scan(&idRobot)
	if err != nil {
		log.Println("Error al obtener idRobot:", err)
		return err
	}

	_, err = r.db.Exec("UPDATE estado_robot SET status = ? WHERE idRobot = ?", nuevoStatus, idRobot)
	if err != nil {
		log.Println("Error al actualizar status en BD:", err)
		return err
	}

	return nil
}

// publishToMQTT publica un mensaje en un tópico MQTT
func (r *MySQLRepository) publishToMQTT(topic string, message string) {
	token := r.MQTTClient.Publish(topic, 0, false, message)
	token.Wait()
	log.Printf("Publicado en MQTT [%s]: %s", topic, message)
}
