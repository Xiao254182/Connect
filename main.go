package main

func main() {
	server := NewServer("192.168.100.3", 8888)
	server.Start()
}
