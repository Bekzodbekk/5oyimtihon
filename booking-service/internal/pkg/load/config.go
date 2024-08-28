package load

import "github.com/spf13/viper"

type Mongo struct {
	MongoHost       string
	MongoPort       int
	MongoDatabase   string
	MongoCollection string
}

type Config struct {
	Mongo              Mongo
	BookingServiceHost string
	BookingServicePort int
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
		BookingServiceHost: viper.GetString("service.host"),
		BookingServicePort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
