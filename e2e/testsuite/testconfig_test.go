package testsuite

import (
	"testing"

	"github.com/cosmos/ibc-go/e2e/relayer"
)

func TestPopulateDefaultsAddsDefaultRelayersWhenEmpty(t *testing.T) {
	tc := TestConfig{
		ChainConfigs: []ChainConfig{
			{Tag: "main"},
			{Tag: "main"},
			{Tag: "main"},
			{Tag: "main"},
		},
		// [AGORIC] RelayerConfigs: []relayer.Config{},
	}

	tc = populateDefaults(tc)

	if len(tc.RelayerConfigs) != 2 {
		t.Fatalf("expected 2 default relayer configs, got %d", len(tc.RelayerConfigs))
	}

	if tc.RelayerConfigs[0].ID != relayer.Rly {
		t.Fatalf("expected first default relayer to be %q, got %q", relayer.Rly, tc.RelayerConfigs[0].ID)
	}

	if tc.RelayerConfigs[1].ID != relayer.Hermes {
		t.Fatalf("expected second default relayer to be %q, got %q", relayer.Hermes, tc.RelayerConfigs[1].ID)
	}
}

func TestGetConfigAddsDefaultRelayersWithoutConfigFile(t *testing.T) {
	t.Setenv(E2EConfigFilePathEnv, "/path/that/does/not/exist.yaml")

	tc := getConfig()

	if len(tc.RelayerConfigs) != 2 {
		t.Fatalf("expected 2 default relayer configs, got %d", len(tc.RelayerConfigs))
	}

	if tc.ActiveRelayer != relayer.Hermes {
		t.Fatalf("expected default active relayer %q, got %q", relayer.Hermes, tc.ActiveRelayer)
	}
}
