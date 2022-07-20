package beer

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int) error
}

type Service struct{}

func NewService() *Service {
	return &Service{} //retorna a posição em memoria
}

func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(ID int) error {
	return nil
}
