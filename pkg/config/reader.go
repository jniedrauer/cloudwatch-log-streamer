package config

import "path/filepath"

// globConfigFiles returns a list of config files from the configuration key
// LogstashConfig
func globConfigFiles() (result []string, err error) {
	return filepath.Glob(Properties["LogstashConfig"])
}
