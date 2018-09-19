package chartvalues

import (
	"testing"

	"github.com/giantswarm/e2etemplates/internal/rendertest"
)

func newAPIExtensionsReleaseConfigFromFilled(modifyFunc func(*APIExtensionsReleaseE2EConfig)) APIExtensionsReleaseE2EConfig {
	c := APIExtensionsReleaseE2EConfig{
		Namespace: "default",

		Operator: APIExtensionsReleaseE2EConfigOperator{
			Name:    "test-operator",
			Version: "1.0.0",
		},
		VersionBundle: APIExtensionsReleaseE2EConfigVersionBundle{
			Version: "1.0.0",
		},
	}

	modifyFunc(&c)
	return c
}

func Test_NewAPIExtensionsReleaseE2E(t *testing.T) {
	testCases := []struct {
		name           string
		config         APIExtensionsReleaseE2EConfig
		expectedValues string
		errorMatcher   func(err error) bool
	}{
		{
			name:           "case 0: invalid config",
			config:         APIExtensionsReleaseE2EConfig{},
			expectedValues: ``,
			errorMatcher:   IsInvalidConfig,
		},
		{
			name:   "case 1: all values set",
			config: newAPIExtensionsReleaseConfigFromFilled(func(v *APIExtensionsReleaseE2EConfig) {}),
			expectedValues: `namespace: default
operator:
  name: test-operator
  version: 1.0.0
versionBundle:
  version: 1.0.0
`,
			errorMatcher: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := NewAPIExtensionsReleaseE2E(tc.config)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}

			if tc.errorMatcher != nil {
				return
			}

			line, difference := rendertest.Diff(values, tc.expectedValues)
			if line > 0 {
				t.Fatalf("line == %d, want 0, diff: %s", line, difference)
			}
		})
	}
}

func Test_NewAPIExtensionsReleaseE2E_invalidConfigError(t *testing.T) {
	testCases := []struct {
		name         string
		config       APIExtensionsReleaseE2EConfig
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: invalid .Namespace",
			config: newAPIExtensionsReleaseConfigFromFilled(func(v *APIExtensionsReleaseE2EConfig) {
				v.Namespace = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 1: invalid .Operator.Name",
			config: newAPIExtensionsReleaseConfigFromFilled(func(v *APIExtensionsReleaseE2EConfig) {
				v.Operator.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 2: invalid .Operator.Version",
			config: newAPIExtensionsReleaseConfigFromFilled(func(v *APIExtensionsReleaseE2EConfig) {
				v.Operator.Version = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 3: invalid .VersionBundle.Version",
			config: newAPIExtensionsReleaseConfigFromFilled(func(v *APIExtensionsReleaseE2EConfig) {
				v.VersionBundle.Version = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAPIExtensionsReleaseE2E(tc.config)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}

			if tc.errorMatcher != nil {
				return
			}
		})
	}
}
