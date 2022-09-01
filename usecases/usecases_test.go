package usecases

import "testing"

func TestGivenTwoNumbersWhenSumarThenReturnResult(t *testing.T) {
	a := 5
	b := 9
	expected := 14
	result := Sumar(a, b)
	if result != expected {
		t.Fatalf("Resultado invalido %d + %d = %d [Expected = %d]", a, b, result, expected)
	}
}
