package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Address  string `yaml:"address"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

func loadEnvVar(key string) string {
	address := os.Getenv(key)
	if address == "" {
		return ""
	}
	return address
}

// 从环境变量加载配置
func (c *Config) LoadFromEnv() {
	address := loadEnvVar("FI_ADDRESS")
	if address != "" {
		c.Address = address
	}

	account := loadEnvVar("FI_ACCOUNT")
	if account != "" {
		c.Account = account
	}

	password := loadEnvVar("FI_PASSWORD")
	if password != "" {
		c.Password = password
	}
}

// 从本地文件加载配置
func (c *Config) LoadFromFile(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var cfg Config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return err
	}
	*c = cfg
	return nil
}

// 保存配置到本地文件
func (c *Config) StoreToFile(file string) error {
	marshal, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(file, marshal, 0644)
}
