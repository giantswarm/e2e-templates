package e2etemplates

import "testing"

func newAWSOperatorChartAllSetFromFullySet(modifyFunc func(*AWSOperatorChart)) AWSOperatorChart {
	fullySet := AWSOperatorChart{
		Guest: AWSOperatorChartGuest{
			Update: AWSOperatorChartGuestUpdate{
				Enabled: true,
			},
		},
		Provider: AWSOperatorChartProvider{
			AWS: AWSOperatorChartProviderAWS{
				Encrypter: "vault",
				Region:    "eu-central-1",
			},
		},
		Secret: AWSOperatorChartSecret{
			AWSOperator: AWSOperatorChartSecretAWSOperator{
				IDRSAPub: "test-idrsa-pub",
				SecretYaml: AWSOperatorChartSecretAWSOperatorSecretYaml{
					Service: AWSOperatorChartSecretAWSOperatorSecretYamlService{
						AWS: AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWS{
							AccessKey: AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWSAccessKey{
								ID:     "test-access-key-id",
								Secret: "test-access-key-secret",
								Token:  "test-access-key-token",
							},
							HostAccessKey: AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWSAccessKey{
								ID:     "test-host-access-key-id",
								Secret: "test-host-access-key-secret",
								Token:  "test-host-access-key-token",
							},
						},
					},
				},
			},
		},

		RegistryPullSecret: "registry-pull-secret",
	}

	modifyFunc(&fullySet)
	return fullySet
}

func Test_AWSOperatorChartValues(t *testing.T) {
	testCases := []struct {
		name           string
		data           Data
		expectedValues string
		errorMatcher   func(err error) bool
	}{
		{
			name:           "case 0: invalid data",
			data:           AWSOperatorChart{},
			expectedValues: ``,
			errorMatcher:   IsInvalidData,
		},
		{
			name: "case 1: all values set",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {}),
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
                id: test-access-key-id
                secret: test-access-key-secret
                token: test-access-key-token
              hostaccesskey:
                id: test-host-access-key-id
                secret: test-host-access-key-secret
                token: test-host-access-key-token

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"registry-pull-secret\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`,
			errorMatcher: nil,
		},
		{
			name: "case 2: all optional values left",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Guest.Update.Enabled = false
				v.Provider.AWS.Encrypter = ""
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
                id: test-access-key-id
                secret: test-access-key-secret
                token: test-access-key-token
              hostaccesskey:
                id: test-host-access-key-id
                secret: test-host-access-key-secret
                token: test-host-access-key-token

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"registry-pull-secret\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`,
			errorMatcher: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := Render(AWSOperatorChartValues, tc.data)

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

			line, difference := diff(values, tc.expectedValues)
			if line > 0 {
				t.Fatalf("line == %d, want 0, diff: %s", line, difference)
			}
		})
	}
}

func Test_AWSOperatorChartValues_Validate(t *testing.T) {
	testCases := []struct {
		name         string
		data         AWSOperatorChart
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: invalid .Provider.AWS.Region",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Provider.AWS.Region = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 1: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 2: invalid .Secret.AWSOperator.IDRSAPub",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.IDRSAPub = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 3: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 4: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 5: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 6: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 7: invalid .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token = ""
			}),
			errorMatcher: IsInvalidData,
		},
		{
			name: "case 8: invalid .RegistryPullSecret",
			data: newAWSOperatorChartAllSetFromFullySet(func(v *AWSOperatorChart) {
				v.RegistryPullSecret = ""
			}),
			errorMatcher: IsInvalidData,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Render(AWSOperatorChartValues, tc.data)

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
