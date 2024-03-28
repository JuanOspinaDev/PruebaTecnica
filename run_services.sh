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

echo "Iniciando API Gateway..."
cd ./APIGateway
dotnet run &
cd ..

wait
