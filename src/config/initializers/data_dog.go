package initializers

type DataDog struct {
	Environment    string `env:"datadog_environment"`
	AgentAddr      string `env:"datadog_agentAddr"`
	ServiceName    string `env:"datadog_serviceName"`
	ServiceVersion string `env:"datadog_serviceVersion"`
}
