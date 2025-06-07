package server

type Server struct {
	url string
}

func New(url string) *Server{
	return &Server{
		url: url,
	}
}
