package mathy

type Operator struct {
	symbol string
	intfunc func(int, int) int
	float64func func(float64, float64) float64
	stringfunc func(string, string) string
}

func NewOperator(symbol string) *Operator {
	return &Operator{symbol, nil, nil, nil}
}

func (op *Operator) IntFunc(f func (int, int) int) {
	op.intfunc = f
}

func (op *Operator) Float64Func(f func (float64, float64) float64) {
	op.float64func = f
}

func (op *Operator) StringFunc(f func (string, string) string) {
	op.stringfunc = f
}

func EvalInt(expression string, operators... *Operator) int {
	return 422
}

func EvalFloat64(expression string, operators... *Operator) float64 {
	return 1.2345
}

func EvalString(expression string, operators... *Operator) string {
	return "string"
}
