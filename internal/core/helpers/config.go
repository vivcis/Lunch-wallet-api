package helpers

// Instance is the global configuration
var Instance *Config

type Config struct {
	DBUser     string `env:"DB_USER"`
	DBPass     string `env:"DB_PASS"`
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"`
	DBTimeZone string `env:"DB_TIMEZONE"`
	DBMode     string `env:"DB_MODE"`
	JWTSecret  string `env:"JWT_SECRET"`
	Env        string `env:"ENV"`
	Port       string `env:"PORT"`
}
