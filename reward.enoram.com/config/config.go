package config

import "github.com/spf13/viper"

type AppConfig struct {
	PORT            string
	FirestoreConfig struct {
		ServiceAccountKey string
	}
	JiraConfig struct {
		SecretKey string
		Endpoint  string
	}
}

func InitConfig(basePath string) (AppConfig, error) {
	viper.SetConfigType("json")
	viper.SetConfigName("app-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(basePath)
	err := viper.ReadInConfig()

	if err != nil {
		return AppConfig{}, err
	}

	viper.Set("FirestoreConfig.ServiceAccountKey", basePath+viper.GetString("FirestoreConfig.ServiceAccountKey"))
	viper.Set("FirestoreConfig.DropshipDllServiceAccountKey", basePath+viper.GetString("FirestoreConfig.DropshipDllServiceAccountKey"))
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)

	return appConfig, err
}
