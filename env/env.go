package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(chainEnvKey string, overload bool) error {
	// Load root .env file
	err := envFileLoad(".env", overload)
	if err != nil {
		return err
	}

	firstEnvKeyValue := os.Getenv(chainEnvKey)
	chainEnvSet := make(map[string]struct{}, 1)

	for {
		// check if we have chain_env is available
		if v, ok := os.LookupEnv(chainEnvKey); ok && v != "" {

			// check if we have already set this chain_env
			if _, ok = chainEnvSet[v]; ok {
				break
			}

			chainEnvSet[v] = struct{}{}

			if !overload {
				// clean chain env key so that next env could be loaded
				os.Unsetenv(chainEnvKey)
			}

			// Load new chained env file
			err = envFileLoad(v, overload)
			if err != nil {
				break
			}
		} else {
			break
		}
	}

	// for not overload, we should reset chain env key
	if !overload {
		os.Setenv(chainEnvKey, firstEnvKeyValue)
	}

	return err
}

func envFileLoad(name string, overload bool) error {
	isPathExist := func(name string) bool {
		if _, err := os.Stat(name); err == nil {
			return true
		}
		return false
	}

	if !isPathExist(name) {
		name = "../" + name
	}

	if overload {
		return godotenv.Overload(name)
	} else {
		return godotenv.Load(name)
	}
}
