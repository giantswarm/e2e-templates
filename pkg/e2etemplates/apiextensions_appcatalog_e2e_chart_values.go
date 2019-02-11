package e2etemplates

const ApiextensionsAppCatalogE2EChartValues = `app:
  config:
    configMap:
      name: {{ .ConfigMap.Name }}
      namespace: {{ .ConfigMap.Namespace }}
    secret:
      name: {{ .Secret.Name }}
      namespace: {{ .Secret.Namespace }}
  description: {{ .Description }}
  logoURL: {{ .logoURL }}
  storage:
    name: {{ .Storage.Type }}
    namespace: {{ .Storage.URL }}
  title: {{ .Title }}
`

type ApiextensionsAppCatalogValues struct {
	ConfigMap   ApiextensionsAppCatalogConfigMap
	Secret      ApiextensionsAppCatalogSecret
	Description string
	LogoURL     string
	Storage     ApiextensionsAppCatalogStorage
	Title       string
}

type ApiextensionsAppCatalogConfigMap struct {
	Name      string
	Namespace string
}

type ApiextensionsAppCatalogSecret struct {
	Name      string
	Namespace string
}

type ApiextensionsAppCatalogStorage struct {
	Type string
	URL  string
}
