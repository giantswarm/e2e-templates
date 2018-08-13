package chartvalues

import (
	"testing"

	"github.com/giantswarm/e2etemplates/internal/rendertest"
)

func newAWSOperatorConfigFromFilled(modifyFunc func(*AWSOperatorConfig)) AWSOperatorConfig {
	c := AWSOperatorConfig{
		Guest: AWSOperatorConfigGuest{
			Update: AWSOperatorConfigGuestUpdate{
				Enabled: true,
			},
		},
		Provider: AWSOperatorConfigProvider{
			AWS: AWSOperatorConfigProviderAWS{
				Encrypter: "vault",
				Region:    "eu-central-1",
			},
		},
		Secret: AWSOperatorConfigSecret{
			AWSOperator: AWSOperatorConfigSecretAWSOperator{
				IDRSAPub: "test-idrsa-pub",
				SecretYaml: AWSOperatorConfigSecretAWSOperatorSecretYaml{
					Service: AWSOperatorConfigSecretAWSOperatorSecretYamlService{
						AWS: AWSOperatorConfigSecretAWSOperatorSecretYamlServiceAWS{
							AccessKey: AWSOperatorConfigSecretAWSOperatorSecretYamlServiceAWSAccessKey{
								ID:     "test-access-key-id",
								Secret: "test-access-key-secret",
								Token:  "test-access-key-token",
							},
							HostAccessKey: AWSOperatorConfigSecretAWSOperatorSecretYamlServiceAWSAccessKey{
								ID:     "test-host-access-key-id",
								Secret: "test-host-access-key-secret",
								Token:  "test-host-access-key-token",
							},
						},
					},
				},
			},
		},
		RegistryPullSecret: "test-registry-pull-secret",
	}

	modifyFunc(&c)
	return c
}

func Test_NewAWSOperator(t *testing.T) {
	testCases := []struct {
		name           string
		config         AWSOperatorConfig
		expectedValues string
		errorMatcher   func(err error) bool
	}{
		{
			name:           "case 0: invalid config",
			config:         AWSOperatorConfig{},
			expectedValues: ``,
			errorMatcher:   IsInvalidConfig,
		},
		{
			name:   "case 1: all values set",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {}),
			expectedValues: `Installation:
  V1:
    Auth:
      Vault:
        Address: http://vault.default.svc.cluster.local:8200
    Guest:
      Kubernetes:
        API:
          Auth:
            Provider:
              OIDC:
                ClientID: ""
                IssueURL: ""
                UsernameClaim: ""
                GroupsClaim: ""
      SSH:
        SSOPublicKey: 'test'
      Update:
        Enabled: true
    Name: ci-aws-operator
    Provider:
      AWS:
        Region: 'eu-central-1'
        DeleteLoggingBucket: true
        IncludeTags: true
        Route53:
          Enabled: true
        Encrypter: 'vault'
    Registry:
      Domain: quay.io
    Secret:
      AWSOperator:
        IDRSAPub: test-idrsa-pub
        SecretYaml: |
          service:
            aws:
              accesskey:
                id: 'test-access-key-id'
                secret: 'test-access-key-secret'
                token: 'test-access-key-token'
              hostaccesskey:
                id: 'test-host-access-key-id'
                secret: 'test-host-access-key-secret'
                token: 'test-host-access-key-token'

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"test-registry-pull-secret\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`,
			errorMatcher: nil,
		},
		{
			name: "case 2: all optional values left",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Guest.Update.Enabled = false
				v.Provider.AWS.Encrypter = ""
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token = ""
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token = ""
			}),
			expectedValues: `Installation:
  V1:
    Auth:
      Vault:
        Address: http://vault.default.svc.cluster.local:8200
    Guest:
      Kubernetes:
        API:
          Auth:
            Provider:
              OIDC:
                ClientID: ""
                IssueURL: ""
                UsernameClaim: ""
                GroupsClaim: ""
      SSH:
        SSOPublicKey: 'test'
      Update:
        Enabled: false
    Name: ci-aws-operator
    Provider:
      AWS:
        Region: 'eu-central-1'
        DeleteLoggingBucket: true
        IncludeTags: true
        Route53:
          Enabled: true
        Encrypter: 'kms'
    Registry:
      Domain: quay.io
    Secret:
      AWSOperator:
        IDRSAPub: test-idrsa-pub
        SecretYaml: |
          service:
            aws:
              accesskey:
                id: 'test-access-key-id'
                secret: 'test-access-key-secret'
                token: ''
              hostaccesskey:
                id: 'test-host-access-key-id'
                secret: 'test-host-access-key-secret'
                token: ''

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"test-registry-pull-secret\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`,
			errorMatcher: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := NewAWSOperator(tc.config)

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

func Test_NewAWSOperator_invalidConfigError(t *testing.T) {
	testCases := []struct {
		name         string
		config       AWSOperatorConfig
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: invalid .Provider.AWS.Region",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Provider.AWS.Region = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 1: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 2: invalid .Secret.AWSOperator.IDRSAPub",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Secret.AWSOperator.IDRSAPub = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 3: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 4: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 5: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 6: invalid .RegistryPullSecret",
			config: newAWSOperatorConfigFromFilled(func(v *AWSOperatorConfig) {
				v.RegistryPullSecret = ""
			}),
			errorMatcher: IsInvalidConfig,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAWSOperator(tc.config)

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
