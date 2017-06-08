package envs

import (
	"os"

	"github.com/jonboulle/clockwork"
)

//환경변수들
var (
	Clock             = clockwork.NewRealClock()
	GoEnv             = os.Getenv("GOENV")
	ServiceHost       = os.Getenv("SERVICE_HOST")
	DBHost            = os.Getenv("DB_HOST")
	DBPort            = os.Getenv("DB_PORT")
	DBDatabase        = os.Getenv("DB_DATABASE")
	DBUser            = os.Getenv("DB_USER")
	DBPassword        = os.Getenv("DB_PASSWORD")
	DBSync            = os.Getenv("DB_SYNC")
	SecretKey         = os.Getenv("SECRET_KEY")
	SystemAdmin       = os.Getenv("SYSTEM_ADMIN")
	SystemPassword    = os.Getenv("SYSTEM_PASSWORD")
	AuthRedisURL      = os.Getenv("AUTH_REDIS_URL")
	AuthRedisPassword = os.Getenv("AUTH_REDIS_PASSWORD")
)

//GetDBSync ...
func GetDBSync() bool {
	if DBSync == "true" {
		return true
	}
	return false
}
