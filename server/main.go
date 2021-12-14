package main

func main() {
	serverListener := NewServerListener()
	defer serverListener.Listener.Close()
	serverListener.Listen()

}
