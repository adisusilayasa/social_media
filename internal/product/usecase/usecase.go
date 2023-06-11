package usecase

import (
	"context"
	"errors"
	"social_media/internal/product/models"
	"social_media/internal/product/repository"

	"github.com/opentracing/opentracing-go"
)

type ProductUsecase struct {
	ProductRepository repository.ProductRepository
}

type ProductUsecaseInterface interface {
	GetProductWithReview(ctx context.Context, productID int) (*models.ProductWithReview, error)
	LikeReview(ctx context.Context, reviewID int, userID int) error
	CancelLikeReview(ctx context.Context, reviewID int, userID int) error
}

func NewProductUsecase(productRepository repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u *ProductUsecase) GetProductWithReview(ctx context.Context, productID int) (*models.ProductWithReview, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.GetProductWithReview")
	defer span.Finish()

	product, err := u.ProductRepository.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	reviews, err := u.ProductRepository.GetReviewsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	productWithReview := &models.ProductWithReview{
		Product: product,
		Reviews: reviews,
	}

	return productWithReview, nil
}

func (u *ProductUsecase) LikeReview(ctx context.Context, reviewID int, userID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.LikeReview")
	defer span.Finish()

	// Check if the review exists
	exists, err := u.ProductRepository.CheckReviewExistence(ctx, reviewID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("review not found")
	}

	// Like the review
	err = u.ProductRepository.LikeReview(ctx, reviewID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUsecase) CancelLikeReview(ctx context.Context, reviewID int, userID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.CancelLikeReview")
	defer span.Finish()

	// Check if the review exists
	exists, err := u.ProductRepository.CheckReviewExistence(ctx, reviewID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("review not found")
	}

	// Cancel the like on the review
	err = u.ProductRepository.CancelLikeReview(ctx, reviewID, userID)
	if err != nil {
		return err
	}

	return nil
}
