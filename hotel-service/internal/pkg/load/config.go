package load

import "github.com/spf13/viper"

type Mongosh struct {
	MongoshHost       string
	MongoshPort       int
	MongoshDatabase   string
	MongoshCollection string
}

type Config struct {
	Mongosh Mongosh

	HotelServiceHost string
	HotelServicePort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Mongosh: Mongosh{
			MongoshHost:       viper.GetString("mongosh.host"),
			MongoshPort:       viper.GetInt("mongosh.port"),
			MongoshDatabase:   viper.GetString("mongosh.database"),
			MongoshCollection: viper.GetString("mongosh.collection"),
		},

		HotelServiceHost: viper.GetString("server.host"),
		HotelServicePort: viper.GetInt("server.port"),
	}

	return &conf, nil
}
