package service

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Process(input string) string {
	return "processed: " + input
}
