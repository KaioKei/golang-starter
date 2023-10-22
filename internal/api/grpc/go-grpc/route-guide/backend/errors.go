package backend

type UnknownGrpcMethodError struct {
	Name string
	Msg  string
}

func (e *UnknownGrpcMethodError) Error() (string, string) {
	// you can set default values for MyError struct and return them
	// to let default empty values (typed defaults), do not edit them here
	// you can still modify them in implementations for context
	e.Msg = "Unknown GRPC Method"
	return e.Msg, e.Name
}
