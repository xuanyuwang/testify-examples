package mockandsubtest

type ExternalService interface {
	Work()
}

type Server struct {
	ExternalService
}

func NewServer(externalService ExternalService) *Server {
	return &Server{
		ExternalService: externalService,
	}
}
