package taskgraph_op

// Rosenbrock is used as standard test function for optimization.
//  f(x, y) = (1-x)^2 + 100(y-x^2)^2, it has a global minimum of 0
// We can have multiple copies of this function, and different scaling too to
// make it harder to solve.
type Rosenbrock struct {
	numOfCopies int
	shouldScale bool
	count       uint64
}

// This implementation should only work with parameter that have indexes
// range from [0, 2*numOfCopies).
func (r *Rosenbrock) Evaluate(in, out Parameter) float32 {
	sum := float64(0)
	for i := 0; i < r.numOfCopies; i += 1 {
		scale := 1
		if r.shouldScale {
			scale = i + 1
		}
		t0 := in.Get(2*i+1) - in.Get(2*i)*in.Get(2*i)
		t1 := 1.0 - in.Get(2*i)
		sum += scale * (100*t0*t0 + t1*t1)
		out.Set(2*i+0, scale*(-400*t0*in.Get(2*i)-2*t1))
		out.Set(2*i+1, scale*200*t0)
	}
	r.count += 1
	return float32(sum)
}
