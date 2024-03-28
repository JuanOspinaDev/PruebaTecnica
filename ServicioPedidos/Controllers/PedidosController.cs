// Controllers/PedidosController.cs
using Microsoft.AspNetCore.Mvc;
using ServicioPedidos.Models;
using ServicioPedidos.Services;
using System.Threading.Tasks;

namespace ServicioPedidos.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class PedidosController : ControllerBase
    {
        private readonly PedidoService _pedidoService;

        public PedidosController(PedidoService pedidoService)
        {
            _pedidoService = pedidoService;
        }

        [HttpPost]
        public async Task<ActionResult> CrearPedido([FromBody] Pedido pedido)
        {
            await _pedidoService.CrearYNotificarPedidoAsync(pedido);
            return CreatedAtAction(nameof(CrearPedido), new { id = pedido.Id }, pedido);
        }
    }
}
