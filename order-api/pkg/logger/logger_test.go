package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInit_SetsLogLevelAndFormatter(t *testing.T) {
	// reseta para garantir controle
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	Init() // executa o que queremos testar

	assert.Equal(t, logrus.InfoLevel, logrus.GetLevel(), "deve configurar o nível para Info")

	formatter, ok := logrus.StandardLogger().Formatter.(*logrus.TextFormatter)
	assert.True(t, ok, "deve usar TextFormatter")
	assert.True(t, formatter.FullTimestamp, "deve habilitar FullTimestamp")
}
