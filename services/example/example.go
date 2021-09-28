package example

import "my_module_path/my_module_name/entities"

type GetExampleParams struct {
	ExampleID string
}

type GetExampleResponse struct {
	Example *entities.Example
}

func (s *Service) GetExample(params GetExampleParams) (*GetExampleResponse, error) {
	ex, err := s.repository.GetExample(params.ExampleID)
	if err != nil {
		return nil, err
	}

	response := &GetExampleResponse{
		Example: ex,
	}

	return response, nil
}

type AddExampleParams struct {
	NewExample entities.Example
}

type AddExampleResponse struct {
	ExampleID string
}

func (s *Service) AddExample(params AddExampleParams) (*AddExampleResponse, error) {
	id, err := s.repository.AddExample(params.NewExample)
	if err != nil {
		return nil, err
	}

	response := &AddExampleResponse{
		ExampleID: id,
	}

	return response, nil
}
