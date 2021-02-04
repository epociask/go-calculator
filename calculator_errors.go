package calculator

type InputError struct{}

func (ie *InputError) Error() string {
	return "Invalid data type passed into calculator.. expecting only type of int or float"
}

type InitError struct{}

func (ie *InitError) Error() string {
	return "variadic float64 values in constructor can only be either length one or zero"
}
