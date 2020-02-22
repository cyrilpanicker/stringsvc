package main

type StringService interface {
	Count(string) int
}

type stringServiceImplementation struct{}

func (stringServiceImplementation) Count(s string) int {
	return len(s)
}
