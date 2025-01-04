package config

type Config struct {
	DB Postgres `mapstructure:"psql"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Password string	`mapstructure:"password"`
	Port     string	`mapstructure:"port"`
}
