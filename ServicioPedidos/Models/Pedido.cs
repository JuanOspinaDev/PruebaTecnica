// Models/Pedido.cs
using System.Text.Json.Serialization;

namespace ServicioPedidos.Models
{
    public class Pedido
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("cliente_id")]
        public int ClienteId { get; set; }

        [JsonPropertyName("concepto")]
        public string Concepto { get; set; } = string.Empty;

        [JsonPropertyName("plazo")]
        public int Plazo { get; set; }

        [JsonPropertyName("valor_total")]
        public decimal ValorTotal { get; set; }
    }
}