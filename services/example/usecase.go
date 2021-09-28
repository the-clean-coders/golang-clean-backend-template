package example

type UseCase interface {
	AddExample(params AddExampleParams) (*AddExampleResponse, error)
	GetExample(params GetExampleParams) (*GetExampleResponse, error)
}
