package bootstrapper

import (
	"log"

	"github.com/spf13/viper"

	util "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
)

// StartUp bootstrapps the application
func StartUp() {
	// Initialize private/public keys for JWT authentication
	util.InitRSAKeys()
	// Initialize Logger objects with Log Level
	util.SetLogLevel(util.Level(AppConfig.LogLevel))
	// Start a MongoDB session
	createDBSession()
	// Add indexes into MongoDB
	addIndexes()
}

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
	LogLevel                                     int
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// loadAppConfig reads config file app.toml and create AppConfig
func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found:", err)
	}
	AppConfig = configuration{}
	AppConfig.Server = viper.GetString("development.Server")
	AppConfig.MongoDBHost = viper.GetString("development.MongoDBHost")
	AppConfig.DBUser = viper.GetString("development.DBUser")
	AppConfig.DBPwd = viper.GetString("development.DBPwd")
	AppConfig.Database = viper.GetString("development.Database")
	AppConfig.LogLevel = viper.GetInt("development.LogLevel")

}
