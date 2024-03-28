package main

import (
	"ServicioCuentasPorPagar/api"
	"ServicioCuentasPorPagar/internal/kafka"
	"ServicioCuentasPorPagar/internal/repository"
	"ServicioCuentasPorPagar/internal/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Configuración de dependencias
	repo := repository.NewMockCuentaPendienteRepository()
	service := services.NewCuentaPendienteService(repo)

	// Iniciar el consumidor de Kafka
	kafkaURL := "localhost:9092"
	topic := "pedidos"
	groupID := "cuentaspendientes"
	kafka.StartConsumer(ctx, kafkaURL, topic, groupID, service)

	// Configurar el router de Gin
	router := api.SetupRouter(repo)

	// Crear un servidor HTTP con Gin
	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	// Iniciar el servidor HTTP en una goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Esperar por señal de interrupción
	<-ctx.Done()

	// Crear un contexto de timeout para el cierre del servidor
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cerrar el servidor HTTP
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Servicio detenido correctamente")
}
