package utils

type ExecData struct {
	values []any
	err    error
}

func ExecWithParams(function func(params ...any) ([]any, error), params ...any) ExecData {
	res, err := function(params...)
	return ExecData{res, err}
}

func ExecWithReturn(function func() ([]any, error)) ExecData {
	res, err := function()
	return ExecData{res, err}
}

func Exec(function func() error) ExecData {
	err := function()
	return ExecData{nil, err}
}

func (e ExecData) BindWithParams(function func(params ...any) ([]any, error)) ExecData {
	if e.err != nil {
		return e
	}
	return ExecWithParams(function, e.values...)
}

func (e ExecData) BindWithReturn(function func() ([]any, error)) ExecData {
	if e.err != nil {
		return e
	}
	return ExecWithReturn(function)
}

func (e ExecData) Bind(function func() error) ExecData {
	if e.err != nil {
		return e
	}
	return Exec(function)
}

func (e ExecData) Error() error {
	return e.err
}

func (e ExecData) Result() any {
	return e.values
}
