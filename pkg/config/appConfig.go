package config

import "github.com/alexedwards/scs/v2"

// AppConfig struct..
type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
}
