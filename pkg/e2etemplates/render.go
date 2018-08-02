package e2etemplates

import (
	"bytes"
	"text/template"

	"github.com/giantswarm/microerror"
)

type Data interface {
	// Validate returns error matched by IsInvalidData when template data
	// is invalid.
	Validate() error
}

func Render(content string, data Data) (string, error) {
	err := data.Validate()
	if err != nil {
		return "", microerror.Mask(err)
	}

	t, err := template.New("e2etemplate").Parse(content)
	if err != nil {
		return "", microerror.Mask(err)
	}

	b := new(bytes.Buffer)
	err = t.Execute(b, data)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return b.String(), nil
}
