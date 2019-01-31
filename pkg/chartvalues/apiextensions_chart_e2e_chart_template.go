package chartvalues

const ApiextensionsChartE2EChartValues = `chart:
  name: "{{ .Name }}"
  namespace: "{{ .Namespace }}"
  config:
    configMap:
      name: "{{ .ConfigMap.Name }}"
      namespace: "{{ .ConfigMap.Namespace }}"
    secret:
      name: "{{ .Secret.Name }}"
      namespace: "{{ .Secret.Namespace }}"
  tarballURL: "{{ .TarballURL }}"

  chartOperator:
	version: "{{ .ChartOperator.Version }}"

  namespace: "{{ .Namespace }} "
`

type ApiextensionsChartValues struct {
	ChartOperator ApiextensionsChartChartOperator
	ConfigMap     ApiextensionsChartConfigMap
	Name          string
	Namespace     string
	Secret        ApiextensionsChartSecret
	TarballURL    string
}

type ApiextensionsChartChartOperator struct {
	Version string
}

type ApiextensionsChartConfigMap struct {
	Name      string
	Namespace string
}

type ApiextensionsChartSecret struct {
	Name      string
	Namespace string
}
