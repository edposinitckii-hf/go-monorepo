package config

// config stores echo configuration
type config struct {
	Port int
	DSN  string
}

// NewConfig creates new configuration
func NewConfig() config {
	return config{DSN: ":memory", Port: 9001}
}
