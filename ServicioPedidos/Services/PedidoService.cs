// Services/PedidoService.cs
using ServicioPedidos.Models;
using ServicioPedidos.Repositories;
using Confluent.Kafka;
using System.Text.Json;
using System.Threading.Tasks;

namespace ServicioPedidos.Services
{
    public class PedidoService
    {
        private readonly IPedidoRepository _pedidoRepository;
        private readonly IProducer<Null, string> _producer;

        public PedidoService(IPedidoRepository pedidoRepository, IProducer<Null, string> producer)
        {
            _pedidoRepository = pedidoRepository;
            _producer = producer;
        }

        public async Task CrearYNotificarPedidoAsync(Pedido pedido)
        {
            // Crear el pedido en la base de datos
            pedido = await _pedidoRepository.CrearPedidoAsync(pedido);

            // Serializar el pedido a JSON para el evento
            var pedidoEvent = JsonSerializer.Serialize(pedido);

            // Publicar el evento en el topic de Kafka
            await _producer.ProduceAsync("pedidos", new Message<Null, string> { Value = pedidoEvent });
        }
    }
}
