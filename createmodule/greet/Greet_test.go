package greet

import (
	"regexp"
	"testing"
)

// тестируем функцию Hello с передачей имени и проверяем
// допустимое возвращаемое значение
func TestHelloName(t *testing.T) {
	name := "Naruto"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Naruto")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Naruto") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// тестируем функцию Hello с передачей пустой строки в аргумент
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
