package Base

type IRepository[T ResponseItem] interface {
	Create(model *T) error
	Read(id int) (*T, error)
	Update(model *T) error
	Delete(id int) error
	Aggregate(pipeline []interface{}) (*[]T, error)
}
