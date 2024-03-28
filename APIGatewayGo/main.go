package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{

		// Grupo de rutas para Clientes
		api.GET("/clientes/", func(c *gin.Context) {
			resp, err := http.Get("http://localhost:8081/clientes")
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Servicio de Clientes no disponible"})
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta del servicio de Clientes"})
				return
			}
			c.Data(resp.StatusCode, "application/json", body)
		})

		// Rutas para el servicio de Cuentas por Pagar
		cuentas := api.Group("/cuentas")
		{
			cuentas.GET("/", func(c *gin.Context) {
				resp, err := http.Get("http://localhost:8082/cuentas")
				if err != nil {
					c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Servicio de Cuentas por Pagar no disponible"})
					return
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta del servicio de Cuentas por Pagar"})
					return
				}
				c.Data(resp.StatusCode, "application/json", body)
			})

			cuentas.GET("/:clienteId", func(c *gin.Context) {
				clienteId := c.Param("clienteId")
				resp, err := http.Get("http://localhost:8082/cuentas/" + clienteId)
				if err != nil {
					c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Servicio de Cuentas por Pagar no disponible"})
					return
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta del servicio de Cuentas por Pagar"})
					return
				}
				c.Data(resp.StatusCode, "application/json", body)
			})

			cuentas.POST("/create", func(c *gin.Context) {
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
					return
				}
				req, err := http.NewRequest("POST", "http://localhost:8082/cuentas/create", bytes.NewBuffer(bodyBytes))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud para el servicio de Cuentas por Pagar"})
					return
				}
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Servicio de Cuentas por Pagar no disponible"})
					return
				}
				defer resp.Body.Close()
				respBody, err := io.ReadAll(resp.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta del servicio de Cuentas por Pagar"})
					return
				}
				c.Data(resp.StatusCode, "application/json", respBody)
			})
		}

		// Grupo de rutas para Pedidos
		api.POST("/pedidos", func(c *gin.Context) {
			// Capturar la solicitud entrante
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
				return
			}

			// Reenviar la solicitud al servicio de Pedidos
			req, err := http.NewRequest("POST", "http://localhost:8083/pedidos", bytes.NewBuffer(bodyBytes))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud para el servicio de Pedidos"})
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Servicio de Pedidos no disponible"})
				return
			}
			defer resp.Body.Close()

			// Leer la respuesta del servicio de Pedidos y reenviarla
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta del servicio de Pedidos"})
				return
			}
			c.Data(resp.StatusCode, "application/json", respBody)
		})
	}

	router.Run(":8088") // Escucha en el puerto 8080
}
