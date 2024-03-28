// Infrastructure/KafkaProducerConfigurator.cs
using Confluent.Kafka;
using Microsoft.Extensions.Configuration;

namespace ServicioPedidos.Infrastructure
{
    public static class KafkaProducerConfigurator
    {
        public static IProducer<Null, string> ConfigureProducer(IConfiguration configuration)
        {
            var producerConfig = new ProducerConfig
            {
                BootstrapServers = configuration["Kafka:BootstrapServers"]
            };

            return new ProducerBuilder<Null, string>(producerConfig).Build();
        }
    }
}
