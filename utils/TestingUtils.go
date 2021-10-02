package utils

import (
	"reflect"
	"testing"
)

/*
	Compara dos valores
*/
func AssertEquals(t *testing.T, esperado interface{}, actual interface{}) {
	if esperado != actual {
		t.Errorf("Esperado %v (tipo: %v), actual %v (tipo: %v)", esperado, reflect.TypeOf(esperado), actual, reflect.TypeOf(actual))
	}
}

func AssertNotEquals(t *testing.T, esperado interface{}, actual interface{}) {
	if esperado == actual {
		t.Errorf("No se esperaba que fueran iguales los 2 valores %v (tipo: %v), %v (tipo: %v)", esperado, reflect.TypeOf(esperado), actual, reflect.TypeOf(actual))
	}
}

func AssertNotError(t *testing.T, err error) {
	if err != nil {
		t.Error("Se ha producido un error", err)
	}
}
