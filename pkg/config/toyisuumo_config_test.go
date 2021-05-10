package config_test

import (
	"os"
	"testing"

	"github.com/44smkn/sqlc-sample/pkg/config"
	"github.com/google/go-cmp/cmp"
)

func TestReadFromToyIsuumoConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		envs map[string]string
		want config.ToyIsuumoConfig
	}{
		{
			name: "simple",
			envs: map[string]string{
				"LOG_LEVEL": "DEBUG",
				"PORT":      "8080",
			},
			want: config.ToyIsuumoConfig{
				LogLevel: "DEBUG",
				Port:     8080,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resetFn := setToyIsuumoConfigs(t, tt.envs)
			defer resetFn()

			got, _ := config.ReadFromEnv()
			if diff := cmp.Diff(tt.want, *got); diff != "" {
				t.Errorf("ReadFromToyIsuumoConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func setToyIsuumoConfig(t *testing.T, key, value string) func() {
	original := os.Getenv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatal(err)
	}
	return func() {
		if original == "" {
			os.Unsetenv(key)
		} else {
			if err := os.Setenv(key, original); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func setToyIsuumoConfigs(t *testing.T, envs map[string]string) func() {
	t.Helper()

	var resetFuncs []func()
	for k, v := range envs {
		r := setToyIsuumoConfig(t, k, v)
		resetFuncs = append(resetFuncs, r)
	}
	return func() {
		for _, f := range resetFuncs {
			f()
		}
	}
}
