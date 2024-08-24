package user

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/app/appmodel"
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/modules/user/userdto"
	"gofiber-boilerplate/modules/user/usermodel"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	jwtIssuer = "appUser"
)

type UserService interface {
	jwt.JwtMiddleware
	Init(db db.DbService)
	Insert(user *usermodel.UserModel) (*userdto.UserDTO, error)
	Update(idString string, updateDTO *userdto.UpdateUserDTO) (*userdto.UserDTO, error)
	List(req *appmodel.GetListRequest) (*appmodel.PaginationResponseList, error)
	Detail(idString string) (*userdto.UserDTO, error)
	Delete(idString string) error
	Login(req *userdto.LoginDTO) (*userdto.LoginResponseDTO, error)
}

type userServiceImpl struct {
	jwtService jwt.JwtService
	db         *gorm.DB
}

func NewUserService(jwtService jwt.JwtService) UserService {
	return &userServiceImpl{
		jwtService: jwtService,
	}
}

func (service *userServiceImpl) validateUniquePhone(phoneNumber string) error {
	var count int64
	service.db.Model(&usermodel.UserModel{}).Where("phone_number = ?", phoneNumber).Count(&count)
	if count > 0 {
		return fiber.NewError(400, "Phone Number already registered")
	}
	return nil
}

// impl `UserService` start

func (service *userServiceImpl) Init(db db.DbService) {
	service.db = db.Default()
}

func (service *userServiceImpl) Insert(user *usermodel.UserModel) (*userdto.UserDTO, error) {
	err := service.validateUniquePhone(user.PhoneNumber)
	if err != nil {
		return nil, err
	}

	pin, err := bcrypt.GenerateFromPassword([]byte(*user.Pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	pinString := string(pin)
	user.Pin = &pinString
	result := service.db.Create(user)
	dto := userdto.MapUserModelToDTO(user)
	return dto, result.Error
}

func (service *userServiceImpl) Update(idString string, updateDTO *userdto.UpdateUserDTO) (*userdto.UserDTO, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	if updateDTO.Password != nil {
		pwd, err := bcrypt.GenerateFromPassword([]byte(*updateDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		pwdString := string(pwd)
		updateDTO.Password = &pwdString
	}
	user := userdto.UserDTO{BaseModel: base.BaseModel{ID: id}}
	result := service.db.Model(&user).Updates(updateDTO)
	if result.Error != nil {
		return nil, result.Error
	}
	return service.Detail(idString)
}

func (service *userServiceImpl) List(req *appmodel.GetListRequest) (*appmodel.PaginationResponseList, error) {
	var count int64
	users := []usermodel.ReadonlyUserModel{}
	query := service.db.Model(users)
	if req.Search != "" {
		query.Where("name ILIKE ?", "%"+req.Search+"%")
	}

	query = query.Session(&gorm.Session{})

	var wg sync.WaitGroup
	wg.Add(2)

	// Perform count and find concurrently using goroutines
	errChan := make(chan error, 2)
	go func() {
		defer wg.Done()
		errChan <- query.Count(&count).Error
	}()

	go func() {
		defer wg.Done()
		query = query.Session(&gorm.Session{})
		errChan <- query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&users).Error
	}()

	wg.Wait()

	var err error
	for i := 0; i < 2; i++ {
		select {
		case err = <-errChan:
			if err != nil {
				return nil, err
			}
		default:
		}
	}

	count32 := int(count)

	return &appmodel.PaginationResponseList{
		Pagination: &appmodel.PaginationResponsePagination{
			Page:  &req.Page,
			Size:  &req.Limit,
			Total: &count32,
		},
		Content: users,
	}, nil
}

func (service *userServiceImpl) Detail(idString string) (*userdto.UserDTO, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	var user userdto.UserDTO
	result := service.db.First(&user, id)
	return &user, result.Error
}

func (service *userServiceImpl) Delete(idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	var user userdto.UserDTO
	result := service.db.Delete(&user, id)
	return result.Error
}

func (service *userServiceImpl) Login(req *userdto.LoginDTO) (response *userdto.LoginResponseDTO, err error) {
	var user usermodel.UserModel
	result := service.db.Where("phone_number = ?", req.PhoneNumber).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(*user.Pin), []byte(req.Pin)) != nil {
		return nil, fiber.NewError(400, "Phone Number and PIN doesn't match.")
	}

	response, err = service.jwtService.GenerateToken(user.ID, jwtIssuer)
	return
}

// impl `UserService` end

// impl `jwt.JwtService` start

func (service *userServiceImpl) CanAccess(c *fiber.Ctx) error {
	return service.jwtService.CanAccess(c, jwtIssuer)
}
func (service *userServiceImpl) CanRefresh(c *fiber.Ctx) error {
	return service.jwtService.CanRefresh(c, jwtIssuer)
}

// impl `jwt.JwtService` end
