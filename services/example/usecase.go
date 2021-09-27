package example

import "my_module_path/my_module_name/entities"

type UseCase interface {
	GetExample(id string) (entities.Example, error)
}
