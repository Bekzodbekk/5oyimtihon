package load

import "github.com/spf13/viper"

type DialServiceConfig struct {
	HotelServiceHost string
	HotelServicePort int

	UserServiceHost string
	UserServicePort int

	BookingServiceHost string
	BookingServicePort int

	NotifServiceHost string
	NotifServicePort int
}

type Config struct {
	DialServiceConfig DialServiceConfig
	ApiGatewayHost    string
	ApiGatewayPort    int

	SecretKey string
	Cert      string
	Key       string
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		DialServiceConfig: DialServiceConfig{
			HotelServiceHost: viper.GetString("hotel_service.host"),
			HotelServicePort: viper.GetInt("hotel_service.port"),

			UserServiceHost: viper.GetString("user_service.host"),
			UserServicePort: viper.GetInt("user_service.port"),

			BookingServiceHost: viper.GetString("booking_service.host"),
			BookingServicePort: viper.GetInt("booking_service.port"),

			NotifServiceHost: viper.GetString("notification_service.host"),
			NotifServicePort: viper.GetInt("notification_service.port"),
		},

		ApiGatewayHost: viper.GetString("server.host"),
		ApiGatewayPort: viper.GetInt("server.port"),

		SecretKey: viper.GetString("secretKey"),
		Cert:      viper.GetString("tls.cert"),
		Key:       viper.GetString("tls.key"),
	}

	return &conf, nil
}
