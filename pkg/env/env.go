package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvLoader struct {
	EnvDir      string
	RootEnvFile string
	ChainEnvKey string
	Overload    bool
}

func DefaultEnvLoader() EnvLoader {
	return EnvLoader{
		EnvDir:      ".",
		RootEnvFile: ".env",
		ChainEnvKey: "chain_env",
		Overload:    true,
	}
}

func (e *EnvLoader) Load() error {
	// Load root .env file
	err := e.loadEnvFile(e.RootEnvFile)
	if err != nil {
		return err
	}

	firstEnvKeyValue := os.Getenv(e.ChainEnvKey)
	chainEnvSet := make(map[string]struct{}, 1)

	for {
		// check if we have chain_env is available
		if v, ok := os.LookupEnv(e.ChainEnvKey); ok && v != "" {

			// check if we have already set this chain_env
			if _, ok = chainEnvSet[v]; ok {
				break
			}

			chainEnvSet[v] = struct{}{}

			if !e.Overload {
				// clean chain env key so that next env could be loaded
				os.Unsetenv(e.ChainEnvKey)
			}

			// Load new chained env file
			err = e.loadEnvFile(v)
			if err != nil {
				break
			}
		} else {
			break
		}
	}

	// for not overload, we should reset chain env key
	if !e.Overload {
		os.Setenv(e.ChainEnvKey, firstEnvKeyValue)
	}

	return err
}

func (e *EnvLoader) loadEnvFile(envFileName string) error {
	envFileName = e.EnvDir + envFileName
	if e.Overload {
		return godotenv.Overload(envFileName)
	} else {
		return godotenv.Load(envFileName)
	}
}
