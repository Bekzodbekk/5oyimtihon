package load

import "github.com/spf13/viper"

type Redis struct {
	RedisHost string
	RedisPort int
}

type Config struct {
	Redis Redis

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
		Redis: Redis{
			RedisHost: viper.GetString("redis.host"),
			RedisPort: viper.GetInt("redis.port"),
		},
		NotifHost: viper.GetString("service.host"),
		NotifPort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
