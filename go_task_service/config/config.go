package config

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-microservices-template/pkg/constants"
	kafkaClient "go-microservices-template/pkg/kafka"
	"go-microservices-template/pkg/logger"
	"go-microservices-template/pkg/postgres"
	"go-microservices-template/pkg/probes"
	"go-microservices-template/pkg/tracing"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "name", "", "Task microservice, microservices config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      *logger.Config      `mapstructure:"logger"`
	Postgresql  *postgres.Config    `mapstructure:"postgres"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	GRPC        GRPC                `mapstructure:"grpc"`
	Kafka       *kafkaClient.Config `mapstructure:"kafka"`
	Probes      probes.Config       `mapstructure:"probes"`
	Jaeger      *tracing.Config     `mapstructure:"jaeger"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type KafkaTopics struct {
	TaskCreate  kafkaClient.TopicConfig `mapstructure:"taskCreate"`
	TaskCreated kafkaClient.TopicConfig `mapstructure:"taskCreated"`
	TaskUpdate  kafkaClient.TopicConfig `mapstructure:"taskUpdate"`
	TaskUpdated kafkaClient.TopicConfig `mapstructure:"taskUpdated"`
	TaskDelete  kafkaClient.TopicConfig `mapstructure:"taskDelete"`
	TaskDeleted kafkaClient.TopicConfig `mapstructure:"taskDeleted"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()

			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/go_tast_service/config/config.yml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}

	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}

	return cfg, nil
}
