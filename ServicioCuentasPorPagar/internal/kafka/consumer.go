package kafka

import (
	"ServicioCuentasPorPagar/internal/models"
	"ServicioCuentasPorPagar/internal/services"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// Consumer contiene la configuración del consumidor de Kafka
type Consumer struct {
	reader   *kafka.Reader
	services services.CuentaPendienteService
}

// NewConsumer crea un nuevo consumidor de Kafka
func NewConsumer(kafkaURL, topic, groupID string, services services.CuentaPendienteService) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		Topic:   topic,
		GroupID: groupID,
	})
	return &Consumer{
		reader:   reader,
		services: services,
	}
}

func (c *Consumer) Start(ctx context.Context) {
    for {
        select {
        case <-ctx.Done(): 
            log.Println("Contexto cancelado, cerrando el consumidor.")
            c.reader.Close() 
            return 

        default:
            m, err := c.reader.ReadMessage(ctx)
            if err != nil {
                log.Printf("could not read message: %v", err)
                if err == context.Canceled {
                    log.Println("Contexto cancelado, cerrando el consumidor.")
                    c.reader.Close() 
                    return
                }
                continue
            }
            log.Println("Mensaje recibido:", string(m.Value))

            var evento models.PedidoEvento
            if err := json.Unmarshal(m.Value, &evento); err != nil {
                log.Printf("error unmarshalling message: %v", err)
                continue 
            }
            log.Println("Evento deserializado:", evento)

            if err := c.services.CrearCuentasAPartirDePedido(evento); err != nil {
                log.Printf("error handling pedido event: %v", err)
            } else {
                log.Println("Cuenta creada a partir del evento:", evento)
            }
        }
    }
}

func StartConsumer(ctx context.Context, kafkaURL, topic, groupID string, service services.CuentaPendienteService) {
	consumer := NewConsumer(kafkaURL, topic, groupID, service)
	go consumer.Start(ctx)
}
