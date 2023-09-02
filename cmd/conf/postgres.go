package conf

// Postgresql configuration for Postgresql
type Postgresql struct {
	Host       string
	User       string
	Password   string
	DBName     string
	Port       string
	DBPoolSize int
}

// LoadPostgresql checks current configured envs for Postgresql
func LoadPostgresql(configs map[string]string) (*Postgresql, error) {
	return &Postgresql{
		Host:     configs[DBHost],
		Port:     configs[DBPort],
		User:     configs[DBUser],
		Password: configs[DBPassword],
		DBName:   configs[DBName],
	}, nil
}
