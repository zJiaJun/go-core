package chat

func Chat() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
