package config

type Minio struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	UseSSL   bool   `mapstructure:"use-ssl" json:"use-ssl" yaml:"use-ssl"`
	Region           string `mapstructure:"region" json:"region" yaml:"region"`
	DurationSeconds  int    `mapstructure:"duration-seconds" json:"duration-seconds" yaml:"duration-seconds"`
}
