package http

import (
	"net/http"
	_ "social_media/internal/product/models"
	"social_media/internal/product/usecase"
	"social_media/pkg/utils"
	"social_media/pkg/zap"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecaseInterface
	logger         zap.Logger
}

func MapProductRoutes(productGroup *echo.Group, logger zap.Logger, productUsecase usecase.ProductUsecaseInterface) {
	h := &ProductHandler{
		ProductUsecase: productUsecase,
		logger:         logger,
	}

	productGroup.GET("/:id", h.GetProductWithReview)
	productGroup.POST("/reviews/:userId/:id/like", h.LikeReview)
	productGroup.DELETE("/reviews/:userId/:id/like", h.CancelLikeReview)
}

// GetProductWithReview godoc
// @Tags Product
// @Summary Get product with reviews
// @Description Get a product along with its reviews by ID
// @ID getProductWithReview
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} utils.Response{data=models.ProductWithReview{models.Product, []models.Review}}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductWithReview(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	ctx := c.Request().Context()
	productWithReview, err := h.ProductUsecase.GetProductWithReview(ctx, productID)
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "Get Product with reviews successfully", productWithReview))
}

// LikeReview godoc
// @Tags Product
// @Summary Like a review
// @Description Like a review by review ID and user ID
// @ID likeReview
// @Param id path int true "Review ID"
// @Param userId path int true "User ID"
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/reviews/{userId}/{id}/like [post]
func (h *ProductHandler) LikeReview(c echo.Context) error {
	reviewID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}
	// Assuming you have user authentication in place and can retrieve the user ID
	// userID := GetUserIDFromRequest(c)

	ctx := c.Request().Context()
	err = h.ProductUsecase.LikeReview(ctx, reviewID, userID)
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "Review liked successfully", nil))
}

// CancelLikeReview godoc
// @Tags Product
// @Summary Cancel like on a review
// @Description Cancel the like on a review by review ID and user ID
// @ID cancelLikeReview
// @Param id path int true "Review ID"
// @Param userId path int true "User ID"
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products/reviews/{userId}/{id}/like [delete]
func (h *ProductHandler) CancelLikeReview(c echo.Context) error {
	reviewID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}
	// Assuming you have user authentication in place and can retrieve the user ID
	// userID := GetUserIDFromRequest(c)

	ctx := c.Request().Context()
	err = h.ProductUsecase.CancelLikeReview(ctx, reviewID, userID)
	if err != nil {
		h.logger.Infof(err.Error())
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "Review like canceled successfully", nil))
}
