package services

import (
	"context"
	"math"
	"net/http"

	"github.com/eatrisno/go-grpc-product-svc/pkg/db"
	"github.com/eatrisno/go-grpc-product-svc/pkg/models"
	"github.com/eatrisno/go-grpc-product-svc/pkg/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product := models.Product{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

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

	limit := req.Limit
	page := req.Page

	switch {
	case limit < 1:
		limit = 5
	case limit > 50:
		limit = 50
	}

	var totalRows int64
	resultTotalRow := s.H.DB.Model(products).Count(&totalRows)
	totalPages := int32(math.Ceil(float64(totalRows) / float64(limit)))

	if resultTotalRow.Error != nil {
		return &pb.ListProductResponse{
			Status: http.StatusNotFound,
			Error:  resultTotalRow.Error.Error(),
		}, nil
	}

	switch {
	case page < 1:
		page = 1
	case page > totalPages:
		page = totalPages
	}

	offset := (page - 1) * limit
	result := s.H.DB.Limit(int(limit)).Offset(int(offset)).Order("id").Find(&products)

	if result.Error != nil {
		return &pb.ListProductResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
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
