#!/bin/bash

echo "Iniciando Servicio de Clientes..."
cd ./ServicioClientes/cmd
go run . &
cd ..
cd ..

echo "Iniciando Servicio de Cuentas por Pagar..."
cd ./ServicioCuentasPorPagar/cmd
go run . &
cd ..
cd ..

echo "Iniciando Servicio de Pedidos..."
cd ./ServicioPedidos
dotnet run &
cd ..

echo "Iniciando APIGateway..."
cd ./APIGatewayGo
go run . &
cd ..

wait
