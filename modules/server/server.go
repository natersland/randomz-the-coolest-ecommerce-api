package server

type ServerInterface interface {
}

func NewServer() ServerInterface {
	return &server{}
}

type server struct {
}
