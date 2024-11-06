package mapper

type IMapper[T any, U any] interface {
	Map(object T) U
}
