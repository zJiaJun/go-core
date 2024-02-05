package factory

type Store interface {
	Create(key string, val string) error
}
