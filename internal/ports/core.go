package ports

type ArithmaticPort interface {
	Addition(a, b int) (int32, error)
	Subtraction(a, b int) (int32, error)
	Multiplication(a, b int) (int32, error)
	Division(a, b int) (int32, error)
}
