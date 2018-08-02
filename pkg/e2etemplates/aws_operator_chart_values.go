package e2etemplates

import "github.com/giantswarm/microerror"

// AWSOperatorChartValues values required by aws-operator-chart, the environment
// variables will be expanded before writing the contents to a file.
const AWSOperatorChartValues = `Installation:
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
        Enabled: {{ .Guest.Update.Enabled }}
    Name: ci-aws-operator
    Provider:
      AWS:
        Region: '{{ .Provider.AWS.Region }}'
        DeleteLoggingBucket: true
        IncludeTags: true
        Route53:
          Enabled: true
	{{- if .Provider.AWS.Encrypter }}
        Encrypter: '{{ .Provider.AWS.Encrypter }}'
	{{- else }}
        Encrypter: 'kms'
	{{- end }}
    Registry:
      Domain: quay.io
    Secret:
      AWSOperator:
        IDRSAPub: {{ .Secret.AWSOperator.IDRSAPub }}
        SecretYaml: |
          service:
            aws:
              accesskey:
                id: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID }}
                secret: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret }}
                token: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token }}
              hostaccesskey:
                id: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID }}
                secret: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret }}
                token: {{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token }}

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"{{ .RegistryPullSecret }}\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`

type AWSOperatorChart struct {
	Guest    AWSOperatorChartGuest
	Provider AWSOperatorChartProvider
	Secret   AWSOperatorChartSecret

	RegistryPullSecret string
}

func (a AWSOperatorChart) Validate() error {
	if a.Provider.AWS.Region == "" {
		return microerror.Maskf(invalidDataError, "%T.Provider.AWS.Region must not be empty", a)
	}
	if a.Secret.AWSOperator.IDRSAPub == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.IDRSAPub must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret must not be empty", a)
	}
	if a.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token == "" {
		return microerror.Maskf(invalidDataError, "%T.Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token must not be empty", a)
	}

	if a.RegistryPullSecret == "" {
		return microerror.Maskf(invalidDataError, "%T.RegistryPullSecret must not be empty", a)
	}

	return nil
}

type AWSOperatorChartGuest struct {
	Update AWSOperatorChartGuestUpdate
}

type AWSOperatorChartGuestUpdate struct {
	Enabled bool
}

type AWSOperatorChartProvider struct {
	AWS AWSOperatorChartProviderAWS
}

type AWSOperatorChartProviderAWS struct {
	Encrypter string
	Region    string
}

type AWSOperatorChartSecret struct {
	AWSOperator AWSOperatorChartSecretAWSOperator
}

type AWSOperatorChartSecretAWSOperator struct {
	IDRSAPub   string
	SecretYaml AWSOperatorChartSecretAWSOperatorSecretYaml
}

type AWSOperatorChartSecretAWSOperatorSecretYaml struct {
	Service AWSOperatorChartSecretAWSOperatorSecretYamlService
}

type AWSOperatorChartSecretAWSOperatorSecretYamlService struct {
	AWS AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWS
}

type AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWS struct {
	AccessKey     AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWSAccessKey
	HostAccessKey AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWSAccessKey
}

type AWSOperatorChartSecretAWSOperatorSecretYamlServiceAWSAccessKey struct {
	ID     string
	Secret string
	Token  string
}
