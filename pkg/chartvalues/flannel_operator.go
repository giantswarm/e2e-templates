package chartvalues

import (
	"github.com/giantswarm/e2etemplates/internal/render"
	"github.com/giantswarm/microerror"
)

type FlannelOperatorConfig struct {
	ClusterName        string
	ClusterRole        FlannelOperatorClusterRole
	ClusterRolePSP     FlannelOperatorClusterRole
	PSP                FlannelOperatorPSP
	RegistryPullSecret string
}

type FlannelOperatorClusterRole struct {
	BindingName string
	Name        string
}

type FlannelOperatorPSP struct {
	Name string
}

func NewFlannelOperator(config FlannelOperatorConfig) (string, error) {
	if config.ClusterName == "" {
		return "", microerror.Maskf(invalidConfigError, "%T.ClusterName must not be empty", config)
	}
	if config.ClusterRole.BindingName == "" {
		config.ClusterRole.BindingName = "flannel-operator"
	}
	if config.ClusterRole.Name == "" {
		config.ClusterRole.Name = "flannel-operator"
	}
	if config.ClusterRolePSP.BindingName == "" {
		config.ClusterRolePSP.BindingName = "flannel-operator-psp"
	}
	if config.ClusterRolePSP.Name == "" {
		config.ClusterRolePSP.Name = "flannel-operator-psp"
	}
	if config.PSP.Name == "" {
		config.PSP.Name = "flannel-operator-psp"
	}
	if config.RegistryPullSecret == "" {
		return "", microerror.Maskf(invalidConfigError, "%T.RegistryPullSecret must not be empty", config)
	}

	values, err := render.Render(flannelOperatorTemplate, config)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return values, nil
}
