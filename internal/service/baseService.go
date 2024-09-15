package service

type baseService struct {
	UserService     *UserService
	CategoryService *CategoryService
}

var ServiceInstance *baseService

func InitService() *baseService {
	ServiceInstance = &baseService{
		UserService:     NewUserService(),
		CategoryService: NewCategoryService(),
	}

	return ServiceInstance
}
