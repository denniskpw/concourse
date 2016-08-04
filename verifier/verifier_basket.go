package verifier

import (
	"net/http"

	"github.com/hashicorp/go-multierror"
	"code.cloudfoundry.org/lager"
)

type VerifierBasket struct {
	verifiers []Verifier
}

func NewVerifierBasket(verifiers ...Verifier) VerifierBasket {
	return VerifierBasket{verifiers: verifiers}
}

func (vb VerifierBasket) Verify(logger lager.Logger, client *http.Client) (bool, error) {
	var errors error

	for _, verifier := range vb.verifiers {
		verified, err := verifier.Verify(logger, client)
		if err != nil {
			errors = multierror.Append(errors, err)
			continue
		}
		if verified {
			return true, nil
		}
	}

	return false, errors
}
