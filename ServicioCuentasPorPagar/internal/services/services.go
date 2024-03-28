package services

import (
	"ServicioCuentasPorPagar/internal/models"
	"ServicioCuentasPorPagar/internal/repository"
	"fmt"
	"log"
	"math"
	"time"
)

// CuentaPendienteService define las funcionalidades para los clientes
type CuentaPendienteService interface {
	GetCuentas() ([]models.CuentaPorPagar, error)
	GetCuentasPorCliente(clienteID int64) ([]models.CuentaPorPagar, error)
	CrearCuentasAPartirDePedido(evento models.PedidoEvento) error
}

type cuentaPendienteService struct {
	cuentaRepo repository.CuentaPendienteRepository
}

func NewCuentaPendienteService(cuentaRepo repository.CuentaPendienteRepository) *cuentaPendienteService {
	return &cuentaPendienteService{
		cuentaRepo: cuentaRepo,
	}
}

// GetClientes llama al metodo de repositorio correspondiente
func (cp *cuentaPendienteService) GetCuentas() ([]models.CuentaPorPagar, error) {
	return cp.cuentaRepo.GetCuentas()
}

// GetClientes llama al metodo de repositorio correspondiente
func (cp *cuentaPendienteService) GetCuentasPorCliente(clienteID int64) ([]models.CuentaPorPagar, error) {
	return cp.cuentaRepo.GetCuentasPorClienteID(clienteID)
}

// CrearCuentasAPartirDePedido toma un evento de pedido y crea cuentas por pagar
func (cp *cuentaPendienteService) CrearCuentasAPartirDePedido(evento models.PedidoEvento) error {
	fechaActual := time.Now()
	plazo := evento.Plazo

	numCuotas := plazo / 30

	if numCuotas < 1 {
		log.Println("La fecha de vencimiento debe ser al menos 30 días después de la fecha actual.")
		return fmt.Errorf("la fecha de vencimiento debe ser al menos 30 días después de la fecha actual")
	}

	valorCuota := evento.ValorTotal / float64(numCuotas)
	valorCuota = math.Round(valorCuota*100) / 100

	for i := 1; i <= int(numCuotas); i++ {
		fechaCuota := fechaActual.AddDate(0, 0, i*30)
		// Formatear la fechaCuota como un string en el formato año-mes-día
		fechaCuotaStr := fechaCuota.Format("2006-01-02")

		cuenta := models.CuentaPorPagar{
			ClienteID:        evento.ClienteID,
			Concepto:         evento.Concepto,
			PedidoID:         evento.ID,
			Valor:            valorCuota,
			FechaVencimiento: fechaCuotaStr,
			NumeroDeCuota:    int64(i),
		}
		log.Printf("cuenta: %v", cuenta)
		if err := cp.cuentaRepo.CreateCuentasPendientes(cuenta); err != nil {
			log.Printf("error al crear cuenta por pagar: %v", err)
			return err
		}
	}

	return nil
}
