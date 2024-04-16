package service

import (
	"coffee_shop_backend/types"
	"encoding/base64"
	"errors"
	"os"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SALT_SIZE = 16

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) (*ProductService, error) {
	err := db.AutoMigrate(&types.Product{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&types.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&types.Session{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&types.Order{})
	if err != nil {
		return nil, err
	}

	productService := ProductService{db}

	// root user setup
	password := os.Getenv("COFFEESHOP_ROOT_PASSWORD")
	if password == "" {
		return nil, errors.New("COFFEESHOP_ROOT_PASSWORD variable not set")
	}

	var user types.User
	result := productService.db.First(&user, "username = ?", "root")

	if result.RowsAffected == 1 {
		err = productService.ChangePassword(&user, password)
		if err != nil {
			return nil, err
		}
	} else {
		err = productService.PostNewUser(&types.UserCreateDto{
			Username:    "root",
			Password:    password,
			AdminRights: true,
		})

		if err != nil {
			return nil, err
		}
	}

	return &productService, nil
}

func (p *ProductService) GetProducts() ([]types.Product, error) {
	var products []types.Product
	result := p.db.Find(&products)
	return products, result.Error
}

func (p *ProductService) GetProductById(id uint64) (*types.Product, error) {
	var product types.Product
	result := p.db.First(&product, id)
	return &product, result.Error
}

func (p *ProductService) PostProduct(product *types.Product) error {
	result := p.db.Create(product)
	return result.Error
}

func (p *ProductService) makeSalt() (string, error) {
	salt := make([]byte, SALT_SIZE)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}

func (p *ProductService) PostNewUser(userDto *types.UserCreateDto) error {
	// check for username uniqueness

	salt, err := p.makeSalt()
	if err != nil {
		return err
	}

	password := append([]byte(userDto.Password), salt...)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := userDto.ToUser(string(salt), string(hash))

	result := p.db.Create(user)
	return result.Error
}

func (p *ProductService) PostLoginUser(userDto *types.UserLoginDto) (string, error) {
	var user types.User
	result := p.db.First(&user, "username = ?", userDto.Username)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected != 1 {
		return "", errors.New("username not found")
	}

	password := append([]byte(userDto.Password), []byte(user.Salt)...)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), password); err != nil {
		return "", err
	}

	token, session := types.NewSession(user.ID)

	result = p.db.Create(session)
	if result.Error != nil {
		return "", result.Error
	}

	return token, nil
}

func (p *ProductService) HasAdminRights(session_id string) bool {
	var session types.Session
	result := p.db.First(&session, "token = ?", session_id)
	if result.RowsAffected != 1 {
		return false
	}

	var user types.User
	result = p.db.First(&user, session.UserID)
	if result.RowsAffected != 1 {
		return false
	}

	return user.AdminRights
}

func (p *ProductService) ChangePassword(user *types.User, password string) error {
	salt, err := p.makeSalt()
	if err != nil {
		return err
	}

	password_bytes := append([]byte(password), salt...)

	hash, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Hash = string(hash)
	user.Salt = string(salt)

	result := p.db.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductService) GetUser(session_id string) (*types.User, error) {
	var session types.Session
	result := p.db.First(&session, "token = ?", session_id)
	if result.RowsAffected != 1 {
		return nil, errors.New("token not found")
	}

	var user types.User
	result = p.db.First(&user, session.UserID)
	if result.RowsAffected != 1 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (p *ProductService) GetUserDetails(session_id string) (*types.UserDetailsDto, error) {
	user, err := p.GetUser(session_id)
	if err != nil {
		return nil, err
	}

	return user.ToDetailsDto(), nil
}
