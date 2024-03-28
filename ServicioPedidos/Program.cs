using Confluent.Kafka;
using ServicioPedidos.Models;
using ServicioPedidos.Repositories;
using ServicioPedidos.Services;

var builder = WebApplication.CreateBuilder(args);

// Configuración del productor de Kafka
var producerConfig = new ProducerConfig { BootstrapServers = builder.Configuration["Kafka:BootstrapServers"] };
builder.Services.AddSingleton<IProducer<Null, string>>(new ProducerBuilder<Null, string>(producerConfig).Build());

// Configuración de servicios y repositorios
builder.Services.AddSingleton<IPedidoRepository, MockPedidoRepository>();
builder.Services.AddScoped<PedidoService>();

var app = builder.Build();

// Endpoint para crear pedidos
app.MapPost("/pedidos", async (Pedido pedido, PedidoService pedidoService) =>
{
    await pedidoService.CrearYNotificarPedidoAsync(pedido);
    return Results.Created($"/pedidos/{pedido.Id}", pedido);
}).WithName("CrearPedido").Produces(201);

app.Run();
