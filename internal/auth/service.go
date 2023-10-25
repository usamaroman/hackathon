package auth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/usamaroman/hackathon/internal/user"
	"github.com/usamaroman/hackathon/pkg/jwt"
)

type service struct {
	repository user.Storage
}

func NewService(rep user.Storage) *service {
	return &service{
		repository: rep,
	}
}

func (s *service) Registration(ctx *gin.Context, dto RegistrationDto) error {
	_, userErr := s.repository.GetOneUserByEmail(ctx, dto.Email)
	if userErr == nil {
		return fmt.Errorf("почта занята")
	}

	password := dto.Password
	hashedPassword, _ := hashPassword(password)

	var cu = user.CreateUserDto{
		Email:    dto.Email,
		Password: hashedPassword,
	}

	createErr := s.repository.CreateUser(ctx, &cu)
	if createErr != nil {
		log.Printf("Error: %v", createErr)
		return fmt.Errorf(createErr.Error())
	}

	return nil
}

func (s *service) Login(ctx *gin.Context, dto LoginDto) (u user.GetUsersDto, token string, err error) {
	u, userErr := s.repository.GetOneUserByEmail(ctx, dto.Email)
	if userErr != nil {
		return u, "", fmt.Errorf("пользователь не найден")
	}

	hashedPassword := checkPasswordHash(dto.Password, u.Password)
	if !hashedPassword {
		return u, "", fmt.Errorf("неверынй пароль")
	}

	token, err = jwt.GenerateAccessToken(u.Id, u.Email, u.Role)
	if err != nil {
		return u, "", err
	}

	log.Printf("user %s logined", u.Email)
	return u, token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
