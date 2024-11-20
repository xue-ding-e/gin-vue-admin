package config

// TODO 合并到插件中
type Wxpay struct {
	MchID                      string `mapstructure:"mch-id" json:"mch-id" yaml:"mch-id"`
	AppID                      string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	Secret                     string `mapstructure:"secret" json:"secret" yaml:"secret"`
	MchCertificateSerialNumber string `mapstructure:"mch-certificate-serial-number" json:"mch-certificate-serial-number" yaml:"mch-certificate-serial-number"`
	MchAPIv3Key                string `mapstructure:"mch-api-v3-key" json:"mch-api-v3-key" yaml:"mch-api-v3-key"`
	PemPath                    string `mapstructure:"pem-path" json:"pem-path" yaml:"pem-path"`
	NotifyUrl                  string `mapstructure:"notify-url" json:"notify-url" yaml:"notify-url"`
}
