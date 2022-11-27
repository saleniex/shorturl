package params

import (
	"os"
	"strconv"
)

// EnvParams implement Params interface for parameters from environment
type EnvParams struct{}

func NewEnvParams() EnvParams {
	return EnvParams{}
}

func (ep EnvParams) Get(name string) string {
	return os.Getenv(name)
}

func (ep EnvParams) GetInt(name string) (int, error) {
	return strconv.Atoi(os.Getenv(name))
}

func (ep EnvParams) GetIntWithDefault(name string, defaultVal int) int {
	val, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		return defaultVal
	}
	return val
}

func (ep EnvParams) GetWithDefault(name, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	} else {
		return val
	}
}
