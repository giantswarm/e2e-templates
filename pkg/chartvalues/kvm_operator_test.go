package chartvalues

import (
	"testing"

	"github.com/giantswarm/e2etemplates/internal/rendertest"
)

func newKVMOperatorConfigFromFilled(modifyFunc func(*KVMOperatorConfig)) KVMOperatorConfig {
	c := KVMOperatorConfig{
		ClusterName: "test-cluster",
		ClusterRole: KVMOperatorClusterRole{
			BindingName: "kvm-operator",
			Name:        "kvm-operator",
		},
		ClusterRolePSP: KVMOperatorClusterRole{
			BindingName: "kvm-operator-psp",
			Name:        "kvm-operator-psp",
		},
		RegistryPullSecret: "test-registry-pull-secret",
	}

	modifyFunc(&c)
	return c
}

func Test_NewKVMOperator(t *testing.T) {
	testCases := []struct {
		name           string
		config         KVMOperatorConfig
		expectedValues string
		errorMatcher   func(err error) bool
	}{
		{
			name:           "case 0: invalid config",
			config:         KVMOperatorConfig{},
			expectedValues: ``,
			errorMatcher:   IsInvalidConfig,
		},
		{
			name:   "case 1: all values set",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {}),
			expectedValues: `clusterRoleBinding: kvm-operator
clusterRoleBindingPSP: kvm-operator-psp
clusterRoleName: kvm-operator
clusterRoleNamePSP: kvm-operator-psp
Installation:
  V1:
    Guest:
      SSH:
        SSOPublicKey: 'test'
      Kubernetes:
        API:
          Auth:
            Provider:
              OIDC:
                ClientID: ""
                IssueURL: ""
                UsernameClaim: ""
                GroupsClaim: ""
      Update:
        Enabled: true
    Secret:
      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"test-registry-pull-secret\"}}}"
labelSelector: 'clusterID=test-cluster'
`,
			errorMatcher: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := NewKVMOperator(tc.config)

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

func Test_NewKVMOperator_invalidConfigError(t *testing.T) {
	testCases := []struct {
		name         string
		config       KVMOperatorConfig
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: invalid .ClusterName",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.ClusterName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 1: invalid .ClusterRole.BindingName",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.ClusterRole.BindingName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 2: invalid .ClusterRole.Name",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.ClusterRole.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 3: invalid .ClusterRolePSP.BindingName",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.ClusterRolePSP.BindingName = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 4: invalid .ClusterRolePSP.Name",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.ClusterRolePSP.Name = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 5: invalid .RegistryPullSecret",
			config: newKVMOperatorConfigFromFilled(func(v *KVMOperatorConfig) {
				v.RegistryPullSecret = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewKVMOperator(tc.config)

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
