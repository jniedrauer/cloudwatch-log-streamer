package config

import (
	"os"
	"regexp"
	"strings"
)

// Properties contains the application configuration. It is initially populated
// with default values.
var Properties = map[string]string{
	"LogLevel":       "INFO",
	"LogstashConfig": "/etc/logstash/conf.d/*.conf",
}

var matchAllCaps = regexp.MustCompile("([a-z0-9])([A-Z])")

func init() {
	for k := range Properties {
		envVar := getEnvVarFromProperty(k)
		envVal, envSet := os.LookupEnv(envVar)
		if !envSet {
			continue
		}
		Properties[k] = envVal
	}
}

// getEnvVarFromProperty excepts a PascalCase string
// and casts it to ENV_VAR_CASE
func getEnvVarFromProperty(propertyName string) string {
	snek := matchAllCaps.ReplaceAllString(propertyName, "${1}_${2}")
	return strings.ToUpper(snek)
}
