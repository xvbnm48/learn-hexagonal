package ports

type DBPort interface {
	CloseDBConnection() error
	AddHistory(answer int32, operation string) error
}
