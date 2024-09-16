package service

type baseService struct {
	UserService     *UserService
	CategoryService *CategoryService
}

var serviceInstance *baseService

func InitService() {
	serviceInstance = &baseService{
		UserService:     NewUserService(),
		CategoryService: NewCategoryService(),
	}
}

func GetServiceInstance() *baseService {
	return serviceInstance
}
