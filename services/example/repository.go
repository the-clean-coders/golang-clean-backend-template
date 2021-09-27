package example

import "my_module_path/my_module_name/entities"

// Repository represents the Example repository.
type Repository interface {
	// GetExample Get an Example entity from its ID
	GetExample(id string) (entities.Example, error)
}
