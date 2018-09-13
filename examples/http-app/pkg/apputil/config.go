package apputil

// configuration for app
type Configuration struct {
	Server   string // WebServer Host
	LogLevel int    // Log Level: 0 - 4
	// Config for DataBase
	DBHost, DBPort, DBUser, DBPassword, Database string
}

// AppConfig holds the configurations used for web app
var AppConfig Configuration

func init() {
	AppConfig = Configuration{}
}
