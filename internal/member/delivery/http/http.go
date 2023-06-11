package http

import (
	"net/http"
	"social_media/internal/member/models"
	"social_media/internal/member/usecase"
	"social_media/pkg/utils"
	"strconv"

	"social_media/pkg/zap"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type MemberHandler struct {
	MemberUsecase *usecase.MemberUsecase
	logger        zap.Logger
}

func MapMemberRoute(MemberGroup *echo.Group, logger zap.Logger, memberUsecase *usecase.MemberUsecase) {
	h := MemberHandler{
		logger:        logger,
		MemberUsecase: memberUsecase,
	}

	MemberGroup.GET("/all", h.GetAllMembers)
	MemberGroup.GET("/:id", h.GetMemberByID)
	MemberGroup.POST("/", h.AddNewMember)
	MemberGroup.PUT("/:id", h.UpdateMember)
	MemberGroup.DELETE("/:id", h.DeleteMember)
}

// GetAllMembers godoc
// @Tags Member
// @Summary Get all members
// @Description Get a list of all members
// @ID getAllMembers
// @Produce json
// @Success 200 {object} utils.Response{data=[]models.Member}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /members/all [get]
func (h *MemberHandler) GetAllMembers(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "http.GetAllMembers")
	defer span.Finish()

	result, err := h.MemberUsecase.GetAllMember(ctx)

	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "ok", result))
}

// GetMemberByID godoc
// @Tags Member
// @Summary Get member by ID
// @Description Get a member by their ID
// @ID getMemberByID
// @Param id path int true "Member ID"
// @Produce json
// @Success 200 {object} utils.Response{data=models.Member}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /members/{id} [get]
func (h *MemberHandler) GetMemberByID(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "http.GetMemberByID")
	defer span.Finish()

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	result, err := h.MemberUsecase.GetMemberByID(ctx, newId)

	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "ok", result))
}

// AddNewMember godoc
// @Tags Member
// @Summary Add a new member
// @Description Add a new member
// @ID addNewMember
// @Accept json
// @Produce json
// @Param member body models.Member true "Member object"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /members/ [post]
func (h *MemberHandler) AddNewMember(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "http.AddNewMember")
	defer span.Finish()

	var member models.Member
	if err := c.Bind(&member); err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	err := h.MemberUsecase.AddNewMember(ctx, &member)

	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "ok", nil))
}

// UpdateMember godoc
// @Tags Member
// @Summary Update member
// @Description Update an existing member
// @ID updateMember
// @Param id path int true "Member ID"
// @Accept json
// @Produce json
// @Param member body models.Member true "Member object"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /members/{id} [put]
func (h *MemberHandler) UpdateMember(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "http.UpdateMember")
	defer span.Finish()

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)

	var member models.Member
	if err := c.Bind(&member); err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	err := h.MemberUsecase.UpdateMember(ctx, newId, &member)

	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "ok", nil))
}

// DeleteMember godoc
// @Tags Member
// @Summary Delete member
// @Description Delete a member by ID
// @ID deleteMember
// @Param id path int true "Member ID"
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /members/{id} [delete]
func (h *MemberHandler) DeleteMember(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "http.DeleteMember")
	defer span.Finish()

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)

	err := h.MemberUsecase.DeleteMember(ctx, newId)

	if err != nil {
		return c.JSON(utils.ErrorResponse(c, err))
	}

	return c.JSON(utils.SuccessResponse(c, http.StatusOK, "ok", nil))
}
