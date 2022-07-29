package ports

type DBPort interface {
	CloseDBConnection()
	AddHistory(answer int32, operation string) error
}
