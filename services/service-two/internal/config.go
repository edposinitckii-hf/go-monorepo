package internal

// config stores app configuration
type config struct {
	DSN string
}

// NewConfig creates new configuration
func NewConfig() config {
	return config{DSN: ":memory"}
}
