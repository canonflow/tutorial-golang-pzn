//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService() *SimpleService {
	wire.Build(
		SimpleRepository,
		SimpleService,
	)
	return nil
}
