package services

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/eatrisno/go-grpc-auth-svc/pkg/db"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/models"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/pb"
	"github.com/eatrisno/go-grpc-auth-svc/pkg/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user = models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  utils.HashPassword(req.Password),
		CreatedAt: time.Now(),
		Status:    int8(1), // 1 = active
	}

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if result := s.H.DB.Where(&models.User{Email: req.Email, Status: 1}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User name not found or incorrect password.",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User name not found or incorrect password.",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	if result := s.H.DB.Model(&user).Updates(models.User{LastLoginAt: time.Now()}); result.Error != nil {
		log.Println(result.Error)
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

type BodylinkEmail struct {
	URL string
}

func (s *Server) Forgot(ctx context.Context, req *pb.ForgotRequest) (*pb.ForgotResponse, error) {
	var user models.User
	if result := s.H.DB.Where(&models.User{Email: req.Email, Status: 1}).First(&user); result.Error != nil {
		return &pb.ForgotResponse{
			Status: http.StatusNotFound,
			Msg:    "User name not found ",
		}, nil
	}

	log.Println("Email sent")

	templateData := BodylinkEmail{
		URL: "https://web.id/",
	}

	utils.SendEmailTemplate(req.Email, "Forgot Password", templateData, "templates/email_forgot_password.html")

	log.Println("Email sent")
	return &pb.ForgotResponse{
		Status: http.StatusOK,
		Msg:    "Password reset link sent to your email",
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
