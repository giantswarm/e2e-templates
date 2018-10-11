package chartvalues

import (
	"testing"

	"github.com/giantswarm/e2etemplates/internal/rendertest"
)

func newCertOperatorConfigFromFilled(modifyFunc func(*CertOperatorConfig)) CertOperatorConfig {
	c := CertOperatorConfig{
		ClusterName: "test-cluster",
		ClusterRole: CertOperatorClusterRole{
			BindingName: "cert-operator",
			Name:        "cert-operator",
		},
		ClusterRolePSP: CertOperatorClusterRole{
			BindingName: "cert-operator-psp",
			Name:        "cert-operator-psp",
		},
		CommonDomain:       "test-domain",
		RegistryPullSecret: "test-registry-pull-secret",
		PSP: CertOperatorPSP{
			Name: "cert-test-psp",
		},
		Vault: CertOperatorVault{
			Token: "test-token",
		},
	}

	modifyFunc(&c)
	return c
}

func Test_NewCertOperator(t *testing.T) {
	testCases := []struct {
		name           string
		config         CertOperatorConfig
		expectedValues string
		errorMatcher   func(err error) bool
	}{
		{
			name:           "case 0: invalid config",
			config:         CertOperatorConfig{},
			expectedValues: ``,
			errorMatcher:   IsInvalidConfig,
		},
		{
			name:   "case 1: all values set",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {}),
			expectedValues: `clusterRoleBindingName: cert-operator
clusterRoleBindingNamePSP: cert-operator-psp
clusterRoleName: cert-operator
clusterRoleNamePSP: cert-operator-psp
Installation:
  V1:
    Auth:
      Vault:
        Address: http://vault.default.svc.cluster.local:8200
        CA:
          TTL: 1440h
    GiantSwarm:
      CertOperator:
        CRD:
          LabelSelector: 'giantswarm.io/cluster=test-cluster'
    Guest:
      Kubernetes:
        API:
          EndpointBase: test-domain
    Secret:
      CertOperator:
        SecretYaml: |
          service:
            vault:
              config:
                token: test-token
      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"test-registry-pull-secret\"}}}"
pspName: cert-test-psp
`,
			errorMatcher: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := NewCertOperator(tc.config)

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

func Test_NewCertOperator_invalidConfigError(t *testing.T) {
	testCases := []struct {
		name         string
		config       CertOperatorConfig
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: invalid .ClusterName",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.ClusterName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 1: invalid .ClusterRole.BindingName",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.ClusterRole.BindingName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 2: invalid .ClusterRole.Name",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.ClusterRole.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 3: invalid .ClusterRolePSP.BindingName",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.ClusterRolePSP.BindingName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 4: invalid .ClusterRolePSP.Name",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.ClusterRolePSP.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 5: invalid .CommonDomain",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.CommonDomain = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 6: invalid .PSP.Name",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.PSP.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 7: invalid .RegistryPullSecret",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.RegistryPullSecret = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 8: invalid .Vault.Token",
			config: newCertOperatorConfigFromFilled(func(v *CertOperatorConfig) {
				v.Vault.Token = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCertOperator(tc.config)

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
