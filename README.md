# PruebaTecnica
Aplicación de microservicios para manejo de clientes, cuentas por pagar y pedidos, implementada con Golang y C para Prueba Tecnica en Servireencauche

## Pre-requisitos

- Docker
- Docker Compose

## Estructura de Microservicios

Lista de microservicios y sus funciones:

- `api_gateway`: El punto de entrada a los microservicios que enruta las solicitudes a los servicios correspondientes.
- `ServicioClientes`: Microservicio en Go que permite la interaccion con los clientes permitiendo ver la lista de clientes.
- `ServicioCuentasPorPagar`: Microservicio en Go que permite la interaccion con las cuentas de pago, tiene 2 endpoints para ver las cuentas de pago por ID de cliente y en general, ademas de permitir al recibir un evento kafka la creacion de una cuenta pendiente.
- `ServicioPedidos`: Microservicio en C# que maneja los pedidos de los usuarios, al recibir un endpoint de creacion de pedidos lo publica a un topico kafka.

## Rutas

Descripción de las rutas expuestas por cada servicio:

- API Gateway: `http://localhost:8080`
  - `/api/clientes/`: Rutas para visualizar clientes. (GET)
  - `/api/cuentas/`: Rutas para visualizar cuentas por pagar. (GET)
  - `/api/cuentas/:clienteId`: Rutas para visualizar cuentas por pagar por ID de cliente. (GET)
  - `/api/pedidos/`: Rutas para gestionar la creación de pedidos. (POST)

* `/api/pedidos/` inicia un evento que crea una cuenta de pago

## Instrucciones de Ejecución
Para correr la aplicación completa incluyendo Apache Kafka con Docker Compose, sigue estos pasos:

1. Clona el repositorio:
git clone [URL de tu repositorio]
cd [Nombre de tu repositorio]

2. Ejecuta el siguiente comando para levantar Apache Kafka:
docker compose up

3. Ejecuta este script que levanta todos los microservicios y la puerta de enlaces
chmod +x run_services.sh
./run_services.sh

Para detener y remover los contenedores creados por Docker Compose, puedes usar:
docker-compose down