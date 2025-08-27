//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	/*
		Parameter isError akan dimasukkan sebagai argument dari NewSimpleRepository karena
		NewSimpleRepository memeliki parameter boolean dan isError memiliki tipe boolean (sama)
	*/
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}
