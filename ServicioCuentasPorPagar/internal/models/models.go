package models

// CuentaPorPagar representa una cuenta individual por pagar
type CuentaPorPagar struct {
	ID               int64   `json:"id"`
	ClienteID        int64   `json:"cliente_id"`
	PedidoID         int64   `json:"pedido_id"`
	Concepto         string  `json:"concepto"`
	NumeroDeCuota    int64   `json:"numero_cuota"`
	Valor            float64 `json:"valor"`
	FechaVencimiento string  `json:"fecha_vencimiento"`
}

// PedidoEvento representa la estructura de datos de un pedido recibido de Kafka
type PedidoEvento struct {
	ID         int64   `json:"id"`
	ClienteID  int64   `json:"cliente_id"`
	Concepto   string  `json:"concepto"`
	ValorTotal float64 `json:"valor_total"`
	Plazo      int64   `json:"plazo"`
}
