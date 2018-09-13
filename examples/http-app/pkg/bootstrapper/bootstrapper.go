// bootstrapper package bootstraps an application by providing various utilities
package bootstrapper

import (
	"log"

	"github.com/spf13/viper"

	util "github.com/shijuvar/gokit/examples/http-app/pkg/apputil"
	"github.com/shijuvar/gokit/examples/http-app/pkg/auth"
)

// StartUp bootstrapps the application
func StartUp() {
	// Initialize private/public keys for JWT authentication
	auth.InitRSAKeys()
	// Initialize Logger objects with Log Level
	util.SetLogLevel(util.Level(util.AppConfig.LogLevel))

	// Add other one-time bootstrapping activities

}

// loadAppConfig reads the config file app.toml and create AppConfig instance
func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config file not found:", err)
	}

	// Configure app specific config values
	util.AppConfig.Server = viper.GetString("devserver.Server")
	util.AppConfig.LogLevel = viper.GetInt("devserver.LogLevel")

	// Configure Postgres configuration values
	util.AppConfig.DBHost = viper.GetString("postgres.Host")
	util.AppConfig.DBPort = viper.GetString("postgres.Port")
	util.AppConfig.DBUser = viper.GetString("postgres.User")
	util.AppConfig.DBPassword = viper.GetString("postgres.Password")
	util.AppConfig.Database = viper.GetString("postgres.Database")

}
