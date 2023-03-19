package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/ServiceWeaver/weaver"

	"github.com/hyorimitsu/hello-serviceweaver/api/notificationservice"
	"github.com/hyorimitsu/hello-serviceweaver/api/paymentservice"
	"github.com/hyorimitsu/hello-serviceweaver/api/productservice"
)

type Server struct {
	httpServer   http.Server
	notification notificationservice.T
	payment      paymentservice.T
	product      productservice.T
}

func NewServer(
	notification notificationservice.T,
	payment paymentservice.T,
	product productservice.T,
) *Server {
	s := &Server{
		notification: notification,
		payment:      payment,
		product:      product,
	}
	s.httpServer.Handler = instrument(s)
	return s
}

func instrument(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		weaver.InstrumentHandler(r.URL.Path, handler).ServeHTTP(w, r)
	})
}

func (s *Server) Run(root weaver.Instance) error {
	lis, err := root.Listener("hello-serviceweaver", weaver.ListenerOptions{
		LocalAddress: fmt.Sprintf(":%s", os.Getenv("PORT")),
	})
	if err != nil {
		return err
	}
	root.Logger().Debug("hello-serviceweaver service available", "address", lis)
	return s.httpServer.Serve(lis)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/products":
		s.listProducts(w, r)
	default:
		pattern := regexp.MustCompile(`^/products/([^/]+)(/(purchase)?)?$`)
		matches := pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		id := matches[1]
		action := matches[3]
		if action == "" {
			s.getProduct(w, r, id)
		} else {
			s.purchaseProduct(w, r, id)
		}
	}
}

func (s *Server) listProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := s.product.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

func (s *Server) getProduct(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	product, err := s.product.Get(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

func (s *Server) purchaseProduct(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	product, err := s.product.Get(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.payment.Purchase(ctx, product.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.notification.Send(ctx, "to@example.com", fmt.Sprintf("Purchased %s ($%d)", product.Name, product.Amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
