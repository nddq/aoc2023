package common

type Puzzle struct {
	Data []byte
}

type Solver interface {
	Solve()
}
