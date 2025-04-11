package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort_Default(t *testing.T) {
	_ = os.Unsetenv("PORT")
	assert.Equal(t, "8080", GetPort())
}

func TestGetPort_FromEnv(t *testing.T) {
	_ = os.Setenv("PORT", "9090")
	assert.Equal(t, "9090", GetPort())
}

func TestGetDistributionCenterURL_Default(t *testing.T) {
	_ = os.Unsetenv("DISTRIBUTION_CENTER_URL")
	assert.Equal(t, "http://localhost:8001/distribuitioncenters", GetDistributionCenterURL())
}

func TestGetDistributionCenterURL_FromEnv(t *testing.T) {
	_ = os.Setenv("DISTRIBUTION_CENTER_URL", "http://fake-url.com/api")
	assert.Equal(t, "http://fake-url.com/api", GetDistributionCenterURL())
}
