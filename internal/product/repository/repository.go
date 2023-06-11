package repository

import (
	"context"
	"errors"
	"fmt"
	"social_media/internal/product/models"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(ctx context.Context, productID int) (*models.Product, error)
	GetReviewsByProductID(ctx context.Context, productID int) ([]*models.Review, error)
	CheckReviewExistence(ctx context.Context, reviewID int) (bool, error)
	LikeReview(ctx context.Context, reviewID int, userID int) error
	CancelLikeReview(ctx context.Context, reviewID int, userID int) error
}

type MySQLProductRepository struct {
	db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

func (r *MySQLProductRepository) GetProductByID(ctx context.Context, productID int) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.GetProductByID")
	defer span.Finish()

	var product models.Product
	err := r.db.WithContext(ctx).First(&product, productID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *MySQLProductRepository) GetReviewsByProductID(ctx context.Context, productID int) ([]*models.Review, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.GetReviewsByProductID")
	defer span.Finish()

	// Check if the product exists
	var product models.Product
	err := r.db.WithContext(ctx).First(&product, productID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with ID %d not found", productID)
		}
		return nil, err
	}
	var reviewData []*models.ReviewData
	err = r.db.WithContext(ctx).
		Model(&models.Review{}).
		Select("review_products.*, COALESCE(l.like_count, 0) AS like_count, members.username, members.gender, members.skintype, members.skincolor").
		Joins("INNER JOIN members ON review_products.id_member = members.ID_MEMBER").
		Joins("LEFT JOIN (SELECT id_review, COUNT(*) AS like_count FROM like_reviews GROUP BY id_review) l ON review_products.id_review = l.id_review").
		Where("review_products.id_product = ?", productID).
		Table("review_products").
		Scan(&reviewData).
		Error

	reviews := make([]*models.Review, len(reviewData))
	for i, data := range reviewData {
		reviews[i] = &models.Review{
			ReviewData: *data,
		}
	}

	if err != nil {
		return nil, err
	}

	return reviews, nil

}

func (r *MySQLProductRepository) CheckReviewExistence(ctx context.Context, reviewID int) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.CheckReviewExistence")
	defer span.Finish()

	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Review{}).
		Where("id_review = ?", reviewID).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MySQLProductRepository) LikeReview(ctx context.Context, reviewID int, userID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.LikeReview")
	defer span.Finish()

	// Check if the user has already liked the review
	exists, err := r.checkLikeExistence(ctx, reviewID, userID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user has already liked the review")
	}

	like := models.LikeReview{
		ReviewID: reviewID,
		MemberID: userID,
	}

	err = r.db.WithContext(ctx).Create(&like).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLProductRepository) CancelLikeReview(ctx context.Context, reviewID int, userID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.CancelLikeReview")
	defer span.Finish()

	// Check if the user has liked the review
	exists, err := r.checkLikeExistence(ctx, reviewID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user has not liked the review")
	}

	err = r.db.WithContext(ctx).
		Where("id_review = ? AND id_member = ?", reviewID, userID).
		Delete(&models.LikeReview{}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLProductRepository) checkLikeExistence(ctx context.Context, reviewID int, userID int) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.LikeReview{}).
		Where("id_review = ? AND id_member = ?", reviewID, userID).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
