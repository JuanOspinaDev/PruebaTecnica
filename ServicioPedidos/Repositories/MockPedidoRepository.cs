// Repositories/MockPedidoRepository.cs
using ServicioPedidos.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace ServicioPedidos.Repositories
{
    public class MockPedidoRepository : IPedidoRepository
    {
        private readonly Dictionary<int, Pedido> _pedidos = new Dictionary<int, Pedido>();
        private int _nextId = 1;

        public async Task<Pedido> CrearPedidoAsync(Pedido pedido)
        {
            // Simula la inserción en la base de datos asignando un ID y guardando en el diccionario
            pedido.Id = _nextId++;
            _pedidos[pedido.Id] = pedido;
            
            return await Task.FromResult(pedido);
        }

        public async Task<Pedido?> ObtenerPedidoAsync(int id)
        {
            _pedidos.TryGetValue(id, out var pedido);
            return await Task.FromResult(pedido);
        }
    }
}