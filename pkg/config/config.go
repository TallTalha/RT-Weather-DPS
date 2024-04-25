package config

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	MongoDB  MongoDBConfig  `mapstructure:"mongodb"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	GRPC     GRPCConfig     `mapstructure:"grpc"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type MongoDBConfig struct {
	URI string `mapstructure:"uri"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

type GRPCConfig struct {
	Port int `mapstructure:"port"`
}
