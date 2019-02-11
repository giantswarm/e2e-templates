package e2etemplates

const ApiextensionsAppE2EChartValues = `app:
  config:
	configMap:
	  name: {{ .ConfigMap.Name }}
	  namespace: {{ .ConfigMap.Namespace }}
    secret:
      name: {{ .Secret.Name }}
      namespace: {{ .Secret.Namespace }}
  kubeConfig:
    secret: 
      name: {{ .KubeConfig.Name }}
      namespace: {{ .KubeConfig.Namespace }}
  name: {{ .Name }}
  namespace: {{ .Namespace }}
`

type ApiextensionsAppValues struct {
	ConfigMap  ApiextensionsAppConfigMap
	Secret     ApiextensionsAppSecret
	KubeConfig ApiextensionsAppKubeConfig
	Name       string
	Namespace  string
}

type ApiextensionsAppConfigMap struct {
	Name      string
	Namespace string
}

type ApiextensionsAppSecret struct {
	Name      string
	Namespace string
}

type ApiextensionsAppKubeConfig struct {
	Name      string
	Namespace string
}
