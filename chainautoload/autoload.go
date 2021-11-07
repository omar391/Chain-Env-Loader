package chainautoload

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/omar391/chain-env-loader/env"
)

func init() {
	env.Load("chain_env", true)
}
