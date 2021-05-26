package main

func main() {
	server := NewServer()
	server.ListenAndServe(":8080")
}
