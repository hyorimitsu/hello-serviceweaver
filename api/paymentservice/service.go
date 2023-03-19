package paymentservice

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

// T is payment service interface.
// NOTE: "Every method in a component interface must receive a context.Context as its first argument and return an error as its final result. All other arguments must be serializable." (https://serviceweaver.dev/docs.html#components-interfaces)
type T interface {
	Purchase(ctx context.Context, amount int) error
}

type impl struct {
	weaver.Implements[T]
}

func (s *impl) Purchase(_ context.Context, amount int) error {
	s.Logger().Info("called purchase function")

	s.Logger().Info(fmt.Sprintf("purchase product [amount: %d]", amount))

	return nil
}
