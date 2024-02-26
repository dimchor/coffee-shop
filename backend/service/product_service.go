package service

import (
	"coffee_shop_backend/types"
	"encoding/base64"

	"bytes"
	"crypto/rand"
	"errors"

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

	return &ProductService{db}, nil
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
	if result != nil {
		return "", result.Error
	}

	password_hash := append([]byte(userDto.Password), []byte(user.Salt)...)

	if !bytes.Equal(password_hash, []byte(user.Hash)) {
		return "", errors.New("incorrect password")
	}

	token, session := types.NewSession(user.ID)

	result = p.db.Create(session)
	if result.Error != nil {
		return "", result.Error
	}

	return token.String(), nil
}
