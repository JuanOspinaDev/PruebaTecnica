package repository

import (
	"ServicioCuentasPorPagar/internal/models"
	"errors"
	"sync"
	"time"
)

var (
	ErrCuentaNoEncontrada = errors.New("cuenta por pagar no encontrada")
)

// CuentaPendienteRepository define la interfaz para el repositorio
type CuentaPendienteRepository interface {
	GetCuentas() ([]models.CuentaPorPagar, error)
	GetCuentasPorClienteID(clienteID int64) ([]models.CuentaPorPagar, error)
}

// mockCuentasPendientesRepository es una implementación de la interfaz
type mockCuentaPendienteRepository struct {
	sync.Mutex
	cuentas map[int64]models.CuentaPorPagar
	nextID  int64
}

// NewMockCuentasPendienteRepository crea un nuevo repositorio mock
func NewMockCuentaPendienteRepository() CuentaPendienteRepository {
	repo := &mockCuentaPendienteRepository{
		cuentas: make(map[int64]models.CuentaPorPagar),
		nextID:  1,
	}

	// Inicializa el repositorio con datos mockeados
	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		Concepto:         "Servicio de reencauche",
		Valor:            100000,
		FechaVencimiento: time.Now().Add(30 * 24 * time.Hour),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        2,
		Concepto:         "Reparación de llanta",
		Valor:            25000,
		FechaVencimiento: time.Now().Add(45 * 24 * time.Hour),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        3,
		Concepto:         "Alineación y balanceo",
		Valor:            80000,
		FechaVencimiento: time.Now().AddDate(0, 1, 0), // Un mes a partir de hoy
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        4,
		Concepto:         "Cambio de aceite",
		Valor:            150000,
		FechaVencimiento: time.Now().AddDate(0, 2, 15), // Dos meses y quince días a partir de hoy
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        5,
		Concepto:         "Mantenimiento preventivo",
		Valor:            200000,
		FechaVencimiento: time.Now().Add(7 * 24 * time.Hour), // Una semana a partir de hoy
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		Concepto:         "Cambio de llantas",
		Valor:            300000,
		FechaVencimiento: time.Now().AddDate(0, 0, 10), // Diez días a partir de hoy
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        2,
		Concepto:         "Pintura de vehículo",
		Valor:            500000,
		FechaVencimiento: time.Now().AddDate(0, 3, 0), // Tres meses a partir de hoy
	}
	repo.nextID++

	return repo
}

// GetCuentas retorna todas las cuentas por pagar
func (r *mockCuentaPendienteRepository) GetCuentas() ([]models.CuentaPorPagar, error) {
	r.Lock()
	defer r.Unlock()
	cuentas := make([]models.CuentaPorPagar, 0, len(r.cuentas))
	for _, cuenta := range r.cuentas {
		cuentas = append(cuentas, cuenta)
	}
	return cuentas, nil
}

// GetCuentasPorClienteID retorna todas las cuentas por pagar
func (r *mockCuentaPendienteRepository) GetCuentasPorClienteID(clienteID int64) ([]models.CuentaPorPagar, error) {
	r.Lock()
	defer r.Unlock()
	var cuentas []models.CuentaPorPagar
	for _, cuenta := range r.cuentas {
		if cuenta.ClienteID == clienteID {
			cuentas = append(cuentas, cuenta)
		}
	}
	return cuentas, nil
}
