package chainautoload

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/omar391/chain-env-loader/conf"
)

func init() {
	conf.Load("chain_env")
}
