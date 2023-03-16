package vstlocation

type LocationStorage interface {
	FindLocationById(id int64) error
}
