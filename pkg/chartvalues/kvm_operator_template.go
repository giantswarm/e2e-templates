package chartvalues

const kvmOperatorTemplate = `clusterRoleBinding: {{ .ClusterRole.BindingName }}
clusterRoleBindingPSP: {{ .ClusterRolePSP.BindingName }}
clusterRoleName: {{ .ClusterRole.Name }}
clusterRoleNamePSP: {{ .ClusterRolePSP.Name }}
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
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"{{ .RegistryPullSecret }}\"}}}"
labelSelector: 'clusterID={{ .ClusterName }}'
`
