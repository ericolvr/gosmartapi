package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Log      Log
	App      App
	RabbitMQ RabbitMQ
	Database Database
	Secure   Secure
}

type App struct {
	Name     string `env:"APP_NAME" envDefault:"asappay-collector-trx"`
	Company  string `env:"APP_COMPANY" envDefault:"asappay"`
	Version  string `env:"APP_VERSION" envDefault:"v1.0.0"`
	RestPort string `env:"APP_REST_PORT" envDefault:"8080"`
	RestHost string `env:"APP_REST_HOST" envDefault:"localhost"`
	GRPCPort string `env:"APP_GRPC_PORT" envDefault:"8081"`
	GRPCHost string `env:"APP_GRPC_HOST" envDefault:"localhost"`
}

type Secure struct {
	TLSEnable bool   `env:"SECURE_TLS_ENABLE" envDefault:"false"`
	TLSCert   string `env:"SECURE_TLS_CERT"`
	TLSKey    string `env:"SECURE_TLS_KEY"`
}

type Log struct {
	KindJSON bool   `env:"LOG_JSON_MESSAGES" envDefault:"false"`
	Level    string `env:"LOG_LEVEL" envDefault:"info"`
}

type RabbitMQ struct {
	Enable     bool   `env:"RABBITMQ_ENABLE" envDefault:"true"`
	Url        string `env:"RABBITMQ_URL" envDefault:"amqp://guest:guest@localhost:5672"`
	Host       string `env:"RABBITMQ_HOST" envDefault:"localhost"`
	Port       int    `env:"RABBITMQ_PORT" envDefault:"5672"`
	User       string `env:"RABBITMQ_USER" envDefault:"guest"`
	Pass       string `env:"RABBITMQ_PASS" envDefault:"guest"`
	Consumer   string `env:"RABBITMQ_CONSUMER" envDefault:""`
	Durable    bool   `env:"RABBITMQ_DURABLE" envDefault:"true"`
	AutoAck    bool   `env:"RABBITMQ_AUTOACK" envDefault:"true"`
	Exclusive  bool   `env:"RABBITMQ_EXCLUSIVE" envDefault:"false"`
	NoWait     bool   `env:"RABBITMQ_NOWAIT" envDefault:"false"`
	NoLocal    bool   `env:"RABBITMQ_NOLOCAL" envDefault:"false"`
	Mandatory  bool   `env:"RABBITMQ_MANDATORY" envDefault:"false"`
	Immediate  bool   `env:"RABBITMQ_IMMEDIATE" envDefault:"false"`
	Name       string `env:"RABBITMQ_NAME" envDefault:"name_queue"`
	Exchange   string `env:"RABBITMQ_EXCHANGE" envDefault:"name_exchange"`
	Kind       string `env:"RABBITMQ_KIND" envDefault:"direct"`
	RoutingKey string `env:"RABBITMQ_ROUTING_KEY" envDefault:"name_router"`
	Internal   bool   `env:"RABBITMQ_INTERNAL" envDefault:"false"`
	AutoDelete bool   `env:"RABBITMQ_AUTO_DELETE" envDefault:"false"`
}

type Database struct {
	Postgres Postgres
}

type Postgres struct {
	Host            string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port            string `env:"POSTGRES_PORT" envDefault:"12432"`
	User            string `env:"POSTGRES_USER" envDefault:"adempiere"`
	Pass            string `env:"POSTGRES_PASS" envDefault:"adempiere"`
	Database        string `env:"POSTGRES_DATABASE" envDefault:"idempiere"`
	SSLMode         string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	ConnMaxIdleTime string `env:"POSTGRES_CONN_MAX_IDLE_TIME" envDefault:"5m"`
	ConnMaxLifetime string `env:"POSTGRES_CONN_MAX_LIFETIME" envDefault:"5m"`
	MaxIdleConns    int32  `env:"POSTGRES_MAX_IDLE_CONNS" envDefault:"5"`
}

func NewConfig() (*Config, error) {

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) GetLog() Log {
	return c.Log
}

func (c *Config) GetApp() App {
	return c.App
}

func (c *Config) GetSecure() Secure {
	return c.Secure
}

func (c *Config) GetRabbitMQ() RabbitMQ {
	return c.RabbitMQ
}

func (c *Config) GetPostgres() Postgres {
	return c.Database.Postgres
}
