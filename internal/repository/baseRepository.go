package repository

type baseRepository struct {
	UserRepository *UserRepository
}

var baseRepositoryInstance *baseRepository

// InitBaseRepository
//
// Parameters:
//
// Returns:
func InitBaseRepository() {
	baseRepositoryInstance = &baseRepository{
		UserRepository: NewUserRepository(),
	}
}

// GetRepository
//
// Parameters:
//
// Returns:
// - *baseRepository
func GetRepository() *baseRepository {
	return baseRepositoryInstance
}
