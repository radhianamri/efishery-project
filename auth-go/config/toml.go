package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	CorsURL       []string `toml:"cors_url"`
	APIRouteURL   string   `toml:"api_route_url"`
	RedisURL      string   `toml:"redis_url"`
	RedisPass     string   `toml:"redis_pass"`
	SwaggerHost   string   `toml:"swagger_host"`
	SwaggerSchema []string `toml:"swagger_schema"`
	JWTSecret     string   `toml:"jwt_secret"`
	Mode          string   `toml:"mode"`
}

func LoadTomlConfig() (err error) {
	var (
		dep           = os.Getenv("deployment_type")
		deplymentType = []string{"LOCALHOST", "PRODUCTION"}
		loadToml      struct {
			Config map[string]config `toml:"config"`
		}
	)

	if dep == "" {
		dep = "LOCALHOST"
	} else {
		dep = strings.ToUpper(dep)
	}

	for _, dt := range deplymentType {
		if dep == dt {
			if _, err = toml.DecodeFile("config/config.toml", &loadToml); err != nil {
				return
			}
			conf = loadToml.Config[dep]
			return
		}
	}

	return fmt.Errorf("Invalid Deployment Type : %s", dep)

}
