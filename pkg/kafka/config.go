package kafka

// Config kafka config
type Config struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupID    string   `mapstructure:"groupID"`
	InitTopics bool     `mapstructure:"topic"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"brokers"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"relicationFactor"`
}
