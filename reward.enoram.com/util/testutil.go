package util

import (
	"github.com/spf13/viper"
	"log"
)

func InitTestConfig() {
	viper.SetConfigFile("./../config/app-config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("## Error while reading config file", err)
		return
	}
	viper.Set("FirestoreConfig.ServiceAccountKey", "./../dropshipstorepos-firebase-adminsdk-j8rxe.json")
	viper.Set("FirestoreConfig.DropshipDllServiceAccountKey", "./../dropshipdll-firebase-adminsdk.json")
}
