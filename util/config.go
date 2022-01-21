
package util

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT 				string
	JWTAccessSecret		string
	JWTRefreshSecret	string
	Redis				[3]string
}

func LoadConfig (path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
        log.Fatal("viper cannot ReadInConfig: ", err)
    }
	err = viper.Unmarshal(&config)
	if err != nil {
        log.Fatal("viper cannot Unmarshal config: ", err)
    }
	return

}