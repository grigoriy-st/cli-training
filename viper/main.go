package main
import (
	"fmt"
	"log"
	"strings"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
	NS string `mapstructure:"namespace"`
	Owner string `mapstructure:"owner"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.BindEnv("app.name", "APP_NAME")

	// appName := viper.GetString("app.name")
	// namespace := viper.GetString("namespace")
	// appPort := viper.GetInt("app.port")

	// fmt.Printf("App Name: %s", appName)
	// fmt.Printf("App Namespace: %s", namespace)
	// fmt.Printf("App Port: %d", appPort)

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error when parsing, %s", err)
	}

	fmt.Printf("Config: %v", config)
}
