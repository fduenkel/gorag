package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL           string
	APIKey            string
	Model             string
	SystemPropmptFile string
}

func Load() Config {
	_ = godotenv.Load()
	cfg := Config{
		BaseURL:           os.Getenv("OPENAI_BASE_URL"),
		APIKey:            os.Getenv("OPENAI_API_KEY"),
		Model:             os.Getenv("OPENAI_MODEL"),
		SystemPropmptFile: os.Getenv("SYSTEM_PROMPT_FILE"),
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.openai.com/v1"
	}

	if cfg.Model == "" {
		cfg.Model = "gpt-4o-mini"
	}

	return cfg
}
