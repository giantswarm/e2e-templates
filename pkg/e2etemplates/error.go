package e2etemplates

import "github.com/giantswarm/microerror"

var invalidDataError = &microerror.Error{
	Kind: "invalidData",
}

// IsInvalidData asserts invalidDataError.
func IsInvalidData(err error) bool {
	return microerror.Cause(err) == invalidDataError
}
