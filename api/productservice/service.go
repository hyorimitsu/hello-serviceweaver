package productservice

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

var (
	//go:embed products.json
	productsData []byte
)

type Product struct {
	weaver.AutoMarshal
	ID          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	HasStock    bool   `json:"has_stock"`
}

// T is product service interface.
// NOTE: "Every method in a component interface must receive a context.Context as its first argument and return an error as its final result. All other arguments must be serializable." (https://serviceweaver.dev/docs.html#components-interfaces)
type T interface {
	List(ctx context.Context) ([]*Product, error)
	Get(ctx context.Context, id string) (*Product, error)
}

type impl struct {
	weaver.Implements[T]

	products []*Product
}

// Init initialize a component.
// NOTE: This function called when a component instance is created.
func (s *impl) Init(context.Context) error {
	s.Logger().Info("initialize product service")

	var products []*Product
	if err := json.Unmarshal(productsData, &products); err != nil {
		return err
	}

	s.products = products

	return nil
}

func (s *impl) List(_ context.Context) ([]*Product, error) {
	s.Logger().Info("called list products function")
	return s.products, nil
}

func (s *impl) Get(_ context.Context, id string) (*Product, error) {
	s.Logger().Info("called get product function")

	for _, p := range s.products {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("not found product [id: %s]", id))
}
