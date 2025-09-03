package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestTypeParameterInheritance(t *testing.T) {
	myManager := &MyManager{
		Name: "Nathan",
	}

	myVicePresident := &MyVicePresident{
		Name: "Garzya",
	}

	getManagerName := GetName[Manager](myManager)
	getVicePresidentName := GetName[VicePresident](myVicePresident)

	assert.Equal(t, "Nathan", getManagerName)
	assert.Equal(t, "Garzya", getVicePresidentName)
	/*
		=== RUN   TestTypeParameterInheritance
		--- PASS: TestTypeParameterInheritance (0.00s)
		PASS

	*/
}
