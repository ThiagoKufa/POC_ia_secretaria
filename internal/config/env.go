package config

import (
	"fmt"
	"os"
)

// Config armazena todas as configurações da aplicação
type Config struct {
	GeminiAPIKey string
	Prompts      *Prompts
}

// LoadConfig carrega todas as configurações do ambiente
func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("a variável de ambiente GEMINI_API_KEY não está definida")
	}

	return &Config{
		GeminiAPIKey: apiKey,
		Prompts:      DefaultPrompts(),
	}, nil
}
