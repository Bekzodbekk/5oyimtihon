package load

import "github.com/spf13/viper"

type Config struct {
	NotifHost string
	NotifPort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		NotifHost: viper.GetString("service.host"),
		NotifPort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
