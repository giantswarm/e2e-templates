package chartvalues

const apiExtensionsAppE2ETemplate = `
apps:
  - name: "{{ .App.Name }}"
    namespace: "{{ .App.Namespace }}"
    catalog: "{{ .App.Catalog }}"
    config:
      configMap:
        name: "{{ .App.Config.ConfigMap.Name }}"
        namespace: "{{ .App.Config.ConfigMap.Namespace }}"
      secret:
        name: "{{ .App.Config.Secret.Name }}"
        namespace: "{{ .App.Config.Secret.Namespace }}"
    kubeConfig:
      inCluster: {{ .App.KubeConfig.InCluster }}
      secret:
        name: "{{ .App.KubeConfig.Secret.Name }}"
        namespace: "{{ .App.KubeConfig.Secret.Namespace }}"
    version: "{{ .App.Version }}"
  # Added chart-operator app CR for e2e testing purpose.
  - name: "chart-operator"
    namespace: "giantswarm"
    catalog: "giantswarm-catalog"
    kubeconfig:
      inCluster: "true"
    version: "0.9.0"

appCatalogs:
  - name: "{{ .AppCatalog.Name }}"
    title: "{{ .AppCatalog.Title }}"
    description: "{{ .AppCatalog.Description }}"
    logoURL: "{{ .AppCatalog.LogoURL }}"
    storage:
      type: "{{ .AppCatalog.Storage.Type }}"
      url: "{{ .AppCatalog.Storage.URL }}"
  - name: "giantswarm-catalog"
    title: "giantswarm-catalog"
    description: "giantswarm catalog"
    logoUrl: "http://giantswarm.com/catalog-logo.png"
    storage:
      type: "helm"
      url: "https://giantswarm.github.com/giantswarm-catalog/"

appOperator:
  version: "{{ .AppOperator.Version }}"

configMaps:
  {{ .App.Config.ConfigMap.Name }}:
    {{ .ConfigMap.ValuesYAML }}

namespace: "{{ .Namespace }}"

secrets:
  {{ .App.Config.Secret.Name }}:
    {{ .Secret.ValuesYAML }}`
