package utils

import (
	"errors"
	"reflect"
	"testing"
)

func multBy2(params ...any) ([]any, error) {
	var res []any = make([]any, len(params))
	for i := 0; i < len(params); i++ {
		res[i] = params[i].(int) * 2
	}
	return res, nil
}

func sum1(params ...any) ([]any, error) {
	var res []any = make([]any, len(params))
	for i := 0; i < len(params); i++ {
		res[i] = params[i].(int) + 1
	}
	return res, nil
}

func returnError1(param ...any) ([]any, error) {
	return []any{}, errors.New("Error 1")
}

func returnError2(param ...any) ([]any, error) {
	return []any{}, errors.New("Error 2")
}

func Test_ExecWithoutErrors(t *testing.T) {
	result := ExecWithParams(multBy2, 1, 2, 3).
		BindWithParams(sum1)

	e := result.Error()
	if e != nil {
		t.Fatalf(`Expected no error and got %s`, e)
	}

	value := result.Result()

	expected := []any{3, 5, 7}
	if !reflect.DeepEqual(value, expected) {
		t.Fatalf(`Expected %v and got %v`, expected, value)
	}
}

func Test_ExecWithErrors(t *testing.T) {
	result := ExecWithParams(returnError1, 1).
		BindWithParams(returnError2)

	e := result.Error()
	if e == nil {
		t.Fatal(`Expected error and got none`)
	}

	if e.Error() != "Error 1" {
		t.Fatalf(`Expected Error 1 and got %s`, e)
	}
}

func Test_ExecWithoutParams(t *testing.T) {
	result := ExecWithParams(sum1, 1).
		BindWithReturn(func() ([]any, error) { return []any{"potato"}, nil })

	e := result.Error()
	if e != nil {
		t.Fatalf(`Error was not expected, got %s`, e)
	}

	value := result.Result()

	expected := []any{"potato"}
	if !reflect.DeepEqual(value, expected) {
		t.Fatalf(`Expected %v and got %v`, expected, value)
	}
}
