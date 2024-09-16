package repository

type baseRepository struct {
	UserRepository *UserRepository
}

var baseRepositoryInstance *baseRepository

func InitBaseRepository() {
	baseRepositoryInstance = &baseRepository{
		UserRepository: NewUserRepository(),
	}
}

func GetRepository() *baseRepository {
	return baseRepositoryInstance
}
