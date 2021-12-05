package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		chain_env_key string
		overload      bool
		expectErr     bool
		expected      string
	}{
		{
			name:          "chain-env overload test",
			chain_env_key: "chain_env",
			overload:      true,
			expectErr:     false,
			expected:      "local3.env",
		},
		{
			name:          "chain-env load test",
			chain_env_key: "chain_env",
			overload:      false,
			expectErr:     false,
			expected:      "local.env",
		},
		{
			name:          "chain-env error test",
			chain_env_key: "chain_env",
			expectErr:     true,
		},
	}

	el := &EnvLoader{
		EnvDir:      "../../",
		RootEnvFile: ".env",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el.ChainEnvKey = tt.chain_env_key
			el.Overload = tt.overload
			err := el.Load()

			if tt.expectErr {
				if err == nil {
					t.Errorf("Load() error = %v, wantErr %v", err, tt.expectErr)
					return
				}

			} else {
				if got := os.Getenv(tt.chain_env_key); tt.expected != got {
					t.Errorf("Expected %s, got %s", tt.expected, got)
				}
			}

		})

		// clean env
		os.Unsetenv(tt.chain_env_key)
	}
}
