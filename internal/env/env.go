package env

type Config struct {
	Port     int    `env:"PORT,default=50001"`
	RedisURL string `env:"REDIS_URL,default=localhost:6379"`
}
