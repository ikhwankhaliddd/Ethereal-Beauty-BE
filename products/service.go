package products

type Service interface {
	GetProducts(userID int) ([]Product, error)
	GetProductDetail(input GetProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(userID int) ([]Product, error) {
	if userID != 0 {
		product, err := s.repository.FindByUserID(userID)
		if err != nil {
			return product, err
		}
		return product, nil
	}
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) GetProductDetail(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}
