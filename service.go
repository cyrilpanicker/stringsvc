package main

type StringService interface {
	Count(string) int
}

type stringService struct{}

func (stringService) Count(s string) int {
	return len(s)
}
