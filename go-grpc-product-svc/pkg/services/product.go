package services

import (
	"context"
	"math"
	"net/http"

	"github.com/eatrisno/go-grpc-product-svc/pkg/db"
	"github.com/eatrisno/go-grpc-product-svc/pkg/models"
	pb "github.com/eatrisno/go-grpc-product-svc/pkg/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	var products []models.Product
	listProductData := []*pb.ListProductData{}

	offset := (req.Page - 1) * req.Limit
	result := s.H.DB.Limit(int(req.Limit)).Offset(int(offset)).Find(&products)

	if result.Error != nil {
		return &pb.ListProductResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	var totalRows int64
	result1 := s.H.DB.Model(products).Count(&totalRows)
	totalPages := int32(math.Ceil(float64(totalRows) / float64(req.Limit)))

	if result1.Error != nil {
		return &pb.ListProductResponse{
			Status: http.StatusNotFound,
			Error:  result1.Error.Error(),
		}, nil
	}

	for i := range products {
		listProductData = append(listProductData, &pb.ListProductData{
			Id:    products[i].Id,
			Name:  products[i].Name,
			Stock: products[i].Stock,
			Price: products[i].Price,
		})
	}

	return &pb.ListProductResponse{
		Status:     http.StatusOK,
		Error:      "",
		Data:       listProductData,
		TotalPages: totalPages,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
