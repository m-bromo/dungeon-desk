package config

const (
	Development = "development"
	Stagin      = "staging"
	Production  = "production"
)

type Environment struct {
	Environment string `env:"ENV,default=development"`
	Api         Server
	PostgresDB  PostgresDatabase
}

type PostgresDatabase struct {
	User     string `env:"DB_PORT,default=admin"`
	Password string `env:"DB_PASSWORD,default=password"`
	Name     string `env:"DB_NAME,default=chardb"`
	Port     string `env:"DB_PORT,default=5432"`
	Host     string `env:"DB_HOST,default=localhost"`
}

type Server struct {
	Host string `env:"API_HOST,default=localhost"`
	Port string `env:"API_PORT,default=8080"`
}
