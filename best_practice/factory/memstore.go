package factory

type memStore struct {
	box map[string]string
}

func init() {
	Register("mem", &memStore{
		box: make(map[string]string, 10),
	})
}

func (mem *memStore) Create(key string, val string) error {
	mem.box[key] = val
	return nil
}
