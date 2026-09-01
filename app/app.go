package app

import (
	"context"

	"github.com/fduenkel/gorag/chat"
	"github.com/fduenkel/gorag/config"
	"github.com/fduenkel/gorag/llm"
)

func Run(ctx context.Context, cfg config.Config) error {
	client := llm.New(cfg)
	return chat.RunREPL(ctx, client, chat.Options{
		SystemPromptFile: cfg.SystemPropmptFile,
	})
}
