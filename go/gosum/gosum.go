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

func (s *GoSum) CalcNTimes(n int64, x float64) float64 {
	for i := 0; int64(i) < n; i++ {
		s.sum += x
	}
	return s.sum
}
