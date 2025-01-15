package config

type Config struct {
	DB     Postgres `mapstructure:"psql"`
	Server Server   `mapstructure:"server"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Postgres struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	DBName   string `mapstructure:"database_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}
