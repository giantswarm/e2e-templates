package chartvalues

const awsOperatorTemplate = `Installation:
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
        Encrypter: '{{ .Provider.AWS.Encrypter }}'
    Registry:
      Domain: quay.io
    Secret:
      AWSOperator:
        IDRSAPub: {{ .Secret.AWSOperator.IDRSAPub }}
        SecretYaml: |
          service:
            aws:
              accesskey:
                id: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.ID }}'
                secret: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Secret }}'
                token: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.AccessKey.Token }}'
              hostaccesskey:
                id: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.ID }}'
                secret: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Secret }}'
                token: '{{ .Secret.AWSOperator.SecretYaml.Service.AWS.HostAccessKey.Token }}'

      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"{{ .RegistryPullSecret }}\"}}}"
    Security:
      RestrictAccess:
        Enabled: false
`
