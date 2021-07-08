package settings

import (
	"bytes"
)

// DatabaseSettings exposes credentials for connecting database
type DatabaseSettings struct {
	Host           string `mapstructure:"host" default:"localhost"`
	Port           string `mapstructure:"port" default:"5432"`
	User           string `mapstructure:"user" validate:"required"`
	Name           string `mapstructure:"name" validate:"required"`
	Password       string `mapstructure:"password" validate:"required"`
	SslMode        string `mapstructure:"sslmode" default:"disable"`
	ConnectionPool string `mapstructure:"connpool" default:"20"`
	Status         bool   `default:"false"`
}

// ConnectionString returns the concatenated string to connect to db
func (s *DatabaseSettings) ConnectionString() string {
	var buffer bytes.Buffer
	buffer.WriteString("host=" + s.Host + " ")
	buffer.WriteString("port=" + s.Port + " ")
	buffer.WriteString("user=" + s.User + " ")
	buffer.WriteString("dbname=" + s.Name + " ")
	buffer.WriteString("password=" + s.Password + " ")
	buffer.WriteString("sslmode=" + s.SslMode)
	return buffer.String()
}

func init() {
	// read the settings from configuration files
	readSettings(DatabasePath, &GetSettings().Database, true)
}
