package load

import "github.com/spf13/viper"

type Mongo struct {
	MongoHost       string
	MongoPort       int
	MongoDatabase   string
	MongoCollection string
}

type Redis struct {
	RedisHost string
	RedisPort int
}

type Config struct {
	Mongo           Mongo
	Redis           Redis
	UserServiceHost string
	UserServicePort int

	SecretKey string
	Email     string
	PassKey   string
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Mongo: Mongo{
			MongoHost:       viper.GetString("mongosh.host"),
			MongoPort:       viper.GetInt("mongosh.port"),
			MongoDatabase:   viper.GetString("mongosh.database"),
			MongoCollection: viper.GetString("mongosh.collection"),
		},
		Redis: Redis{
			RedisHost: viper.GetString("redis.host"),
			RedisPort: viper.GetInt("redis.port"),
		},
		UserServiceHost: viper.GetString("service.host"),
		UserServicePort: viper.GetInt("service.port"),

		SecretKey: viper.GetString("secretKey"),
		Email:     viper.GetString("email"),
		PassKey:   viper.GetString("passKey"),
	}

	return &conf, nil
}
