package registry

import "io"

type InputData interface {
	String() string
}

type RegisteredInputData struct {
	Function string
	MethodID string
	Params   []*ConstractParams
}

type ConstractParams struct {
	Name string
	Type string
	Data string
}

func (r *RegisteredInputData) String() string {
	panic("implement me")
}

type CommonInputData struct {
	//Function string
	MethodID string
	Params   []string
}

func (c *CommonInputData) String() string {
	panic("implement me")
}

type HexString string

func (h HexString) MarshalCBOR(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (h HexString) UnmarshalCBOR(r io.Reader) error {
	//TODO implement me
	panic("implement me")
}
