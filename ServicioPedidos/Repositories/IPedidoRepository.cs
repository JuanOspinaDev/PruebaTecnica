// Repositories/IPedidoRepository.cs
using ServicioPedidos.Models;

namespace ServicioPedidos.Repositories
{
    public interface IPedidoRepository
    {
        Task<Pedido> CrearPedidoAsync(Pedido pedido);
        Task<Pedido?> ObtenerPedidoAsync(int id);
    }
}