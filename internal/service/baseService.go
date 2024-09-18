package service

type baseService struct {
	UserService     *UserService
	CategoryService *CategoryService
}

var serviceInstance *baseService

// InitService
//
// Parameters:
//
// Returns:
func InitService() {
	serviceInstance = &baseService{
		UserService:     NewUserService(),
		CategoryService: NewCategoryService(),
	}
}

// GetServiceInstance
//
// Parameters:
//
// Returns:
// - *baseService
func GetServiceInstance() *baseService {
	return serviceInstance
}
