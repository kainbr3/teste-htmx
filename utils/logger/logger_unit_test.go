package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected zap.AtomicLevel
	}{
		{"Debug level", "DEBUG", zap.NewAtomicLevelAt(zap.DebugLevel)},
		{"Warn level", "WARN", zap.NewAtomicLevelAt(zap.WarnLevel)},
		{"Error level", "ERROR", zap.NewAtomicLevelAt(zap.ErrorLevel)},
		{"Info level", "INFO", zap.NewAtomicLevelAt(zap.InfoLevel)},
		{"Empty level", "", zap.NewAtomicLevelAt(zap.InfoLevel)},
		{"Default level", "UNKNOWN", zap.NewAtomicLevelAt(zap.InfoLevel)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getLevel(tt.level)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
