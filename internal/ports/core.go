package ports

type ArithmaticPort interface {
	Addition(a, b int32) (int32, error)
	Subtraction(a, b int32) (int32, error)
	Multiplication(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}
