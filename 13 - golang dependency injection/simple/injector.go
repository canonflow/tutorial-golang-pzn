//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

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

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

// Provider Set
// Hanya Groping antar Provider
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Binding Interface
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // Kalau ada provider yang butuh SayHello, maka kirim *SayHelloImpl
)

func InitialieHelloService() *HelloService {
	/* ===== Harapannya =====
	sayHello := NewSayHelloImpl()
	helloService := NewHelloService(sayHello)
	*/

	/*
		Salah -> Karena parameter HelloService berupa interface SayHello, walaupun SayHelloImpl merupakan
		struct kontrak dari interface SayHello
	*/
	//wire.Build(NewHelloService, NewSayHelloImpl)

	// Perlu melakukan Binding Interface (kode di atas function ini)
	wire.Build(helloSet, NewHelloService)

	return nil
}

// Struct Provider
// ===== STRUCT PROVIDER =====
func InitializedFooBarStruct() *FooBarStruct {
	wire.Build(
		NewFoo,
		NewBar,
		wire.Struct(new(FooBarStruct), "Foo", "Bar"), // Use * for all fields
	)

	return nil
}

// Binding Values
var fooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}), // Bisa juga dengan variable lainnya
)

func InitializeFooBarUsingVariable() *FooBarStruct {
	wire.Build(fooBarValueSet, wire.Struct(new(FooBarStruct), "*"))

	return nil
}

// Interface Value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))

	return nil
}

// Struct Field Provider
func InitializedConfiguration() *Configuration {
	// Mencari Configuration di Application dan dijadikan sbg provider
	/*
		application := NewApplication()
		configuration := application.Configuration
		return configuration
	*/
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)

	return nil
}
