package configs

type AppConfig struct {
	DatabaseFile  string `env:"DATABASE_FILE" yaml:"database_file"`
	ListenAddress string `env:"LISTEN_ADDRESS" yaml:"listen_address"`
	AutoMigrate   bool   `env:"AUTOMIGRATE" yaml:"automigrate"`
	JwtKey 		  string `env:"JWT_KEY" yaml:"jwt_key"`
}
