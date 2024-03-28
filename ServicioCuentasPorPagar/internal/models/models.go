package models

import "time"

// CuentaPorPagar representa una cuenta individual por pagar
type CuentaPorPagar struct {
	ID               int64     `json:"id"`
	ClienteID        int64     `json:"cliente_id"`
	Concepto         string    `json:"concepto"`
	Valor            float64   `json:"valor"`
	FechaVencimiento time.Time `json:"fecha_vencimiento"`
}
