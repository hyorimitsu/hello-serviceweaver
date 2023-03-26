package notificationservice

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

// T is notification service interface.
// NOTE: "Every method in a component interface must receive a context.Context as its first argument and return an error as its final result. All other arguments must be serializable." (https://serviceweaver.dev/docs.html#components-interfaces)
type T interface {
	Send(ctx context.Context, to, message string) error
}

type impl struct {
	weaver.Implements[T]
	weaver.WithConfig[config]
}

// config is component configures.
// NOTE: If you run on single, `weaver.toml` is not referenced (run via `go run`, not `weaver deploy`), so no value is set to config.
type config struct {
	From string `toml:"notification_from"`
}

// Validate is validate configs.
// NOTE: This function called when execute `deploy` command.
func (c *config) Validate() error {
	if c.From == "" {
		return fmt.Errorf("required notification_from setting in weaver.toml")
	}
	return nil
}

// Init initialize a component.
// NOTE: This function called when a component instance is created.
func (s *impl) Init(context.Context) error {
	s.Logger().Info("initialize notification service")
	return nil
}

func (s *impl) Send(_ context.Context, to, message string) error {
	s.Logger().Info("called send notification function")

	cfg := s.Config()
	s.Logger().Info(fmt.Sprintf("send message [from: %s, to: %s, message: %s]", cfg.From, to, message))

	return nil
}
