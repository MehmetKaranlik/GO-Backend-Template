package Database

type IDatabase interface {
	Connect() error
}
