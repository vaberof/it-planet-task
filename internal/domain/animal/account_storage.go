package animal

type AccountStorage interface {
	FindAccountById(id int32) error
}
