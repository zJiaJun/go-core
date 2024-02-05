package factory

func Run() {
	store, err := New("mem")
	if err != nil {
		panic(err)
	}
	store.Create("k", "v")
}
