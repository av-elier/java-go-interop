package gosum

type GoSum struct {
	sum float64
}

func CreateGoSum() *GoSum {
	return &GoSum{}
}

func (s *GoSum) Calc(x float64) float64 {
	s.sum += x
	return s.sum
}
