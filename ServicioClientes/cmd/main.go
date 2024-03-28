package main

import "ServicioClientes/api"

func main() {
	router := api.SetUpRouter()
	router.Run(":8081")
}
