package main

import "ServicioCuentasPorPagar/api"

func main() {
	router := api.SetUpRouter()
	router.Run(":8080")
}
