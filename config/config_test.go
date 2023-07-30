package config_test

import (
	"os"
	"qdrant/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetEnviron(t *testing.T) {
	err := godotenv.Load("test.env")
	assert.Nil(t, err)

	expectedValue := "TestValue"
	os.Setenv("TEST_VARIABLE", expectedValue)

	result := config.GetEnviron("TEST_VARIABLE")
	assert.Equal(t, expectedValue, result)

	result = config.GetEnviron("NON_EXISTING_VARIABLE")
	assert.Empty(t, result)

	os.Unsetenv("TEST_VARIABLE")
}
