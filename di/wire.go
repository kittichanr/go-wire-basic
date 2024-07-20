// di/wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/kittichanr/go-wire/greeter"
)

func InitializeApplication(ctx context.Context) (greeter.Event, func(), error) {
	wire.Build(greeter.MainBindingSet)
	return greeter.Event{}, func() {}, nil
}
