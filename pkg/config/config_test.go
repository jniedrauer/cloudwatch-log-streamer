package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVarFromProperty(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "LogLevel",
			expect: "LOG_LEVEL",
		},
		{
			input:  "Foo",
			expect: "FOO",
		},
	}

	for _, test := range tests {
		result := getEnvVarFromProperty(test.input)

		assert.Equal(t, test.expect, result)
	}
}
