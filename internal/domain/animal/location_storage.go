package animal

type LocationStorage interface {
	FindLocationById(id int64) error
}
