//Package conf provides settings of app
package conf

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(chain_env_key string) {
	chainEnvSet := make(map[string]struct{})

	for {
		// check if we have chain_env is available
		if v, ok := os.LookupEnv(chain_env_key); ok && v != "" {

			// check if we have already set this chain_env
			_, ok = chainEnvSet[v]
			if !ok {
				chainEnvSet[v] = struct{}{}
			}
			godotenv.Load(v)

		} else {
			break
		}
	}
}
