package config

import (
	"os"
	"strings"
)

//configs.env там смотреть конфиги тут я их в основном подготавливаю к экспорту и добавляю некоторые СПЕЦИАЛИЗИРОВАННЫЕ
// конфиги
type LogConfig struct {
    Level         string `yaml:"level"`         
    Encoding      string `yaml:"encoding"`     
    OutputPaths   []string `yaml:"outputPaths"`  
    EnableCaller  bool   `yaml:"enableCaller"`  
    StacktraceLevel string `yaml:"stacktraceLevel"` 
}
func LoadFromEnv() *LogConfig {
	return &LogConfig{
		Level:         getEnv("LOG_LEVEL", "info"), 
		Encoding:      getEnv("LOG_ENCODING", "json"),
		OutputPaths:   parseOutputPaths(getEnv("OUTPUT_PATHS", "stdout")),
		EnableCaller:  getEnvBool("ENABLE_CALLER", true),
		StacktraceLevel: getEnv("STACKTRACE_LEVEL", "error"),
	}
}
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return strings.ToLower(value) == "true"
	}
	return defaultValue
}

func parseOutputPaths(paths string) []string {
	rawPaths := strings.Split(paths, ",")
	outputPaths := make([]string, 0, len(rawPaths))
	
	for _, path := range rawPaths {	
		trimmed := strings.TrimSpace(path)
		if trimmed != "" {
			outputPaths = append(outputPaths, trimmed)
		}
	}
	
	if len(outputPaths) == 0 {
		return []string{"stdout"}
	}
	
	return outputPaths
}

// Конфиг контролирует путь до env этот конфиг нужно менять в случае если env переносится
// Он помогает запускать приложение в докере (т.к у приложения в докере и вне разные пути)
func GetConfigPath() string {
	if os.Getenv("IS_DOCKER") == "TRUE" {
		return "/app/config/configs.env" 
	}
	return "./config/configs.env"
	     
}
