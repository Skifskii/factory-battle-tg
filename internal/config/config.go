package config

import (
	"log"
	"sync"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfighcl"
)

type Config struct {
	TelegramBotToken string `hcl:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN" required:"true"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "FBB",
			Files:     []string{"./config.hcl"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".hcl": aconfighcl.New(),
			},
		})

		if err := loader.Load(); err != nil {
			log.Printf("failed to load config: %v", err)
		}
	})

	return cfg
}
