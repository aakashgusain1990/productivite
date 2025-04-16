// env:
//
//	local:
//	  port: 8080
//	  database:
//	    host: localhost
//	    port: 5432
//	    user: postgres
//	    password: admin
//	    db: productivite
package inits

import (
	"sync"

	"github.com/aakashgusain1990/productivite/pkg/entities"
	"github.com/spf13/viper"
) 


var (
	once sync.Once
	configInstance *entities.Config
)

func GetConfig() *entities.Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/root")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}