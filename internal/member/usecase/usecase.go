package usecase

import (
	"context"
	"social_media/internal/member/models"
	"social_media/internal/member/repository"

	"github.com/opentracing/opentracing-go"
)

type MemberUsecase struct {
	MemberRepository *repository.MySQLRepository
}

type MemberUsecaseInterface interface {
	GetMemberByID(ctx context.Context, id int) (*models.Member, error)
	GetAllMember(ctx context.Context) ([]*models.Member, error)
	UpdateMember(ctx context.Context, id int, member *models.Member) error
	DeleteMember(ctx context.Context, id int) error
	AddNewMember(ctx context.Context, member *models.Member) error
}

func NewMemberUsecase(MemberRepository *repository.MySQLRepository) *MemberUsecase {
	return &MemberUsecase{MemberRepository: MemberRepository}
}

func (h *MemberUsecase) GetMemberByID(ctx context.Context, id int) (*models.Member, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.GetMemberByID")
	defer span.Finish()

	result, err := h.MemberRepository.GetMemberByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h *MemberUsecase) GetAllMember(ctx context.Context) ([]*models.Member, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.GetMemberByID")
	defer span.Finish()

	result, err := h.MemberRepository.GetAllMembers(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h *MemberUsecase) UpdateMember(ctx context.Context, id int, member *models.Member) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.UpdateMember")
	defer span.Finish()

	err := h.MemberRepository.UpdateMemberByID(ctx, member, id)
	if err != nil {
		return err
	}
	return nil
}

func (h *MemberUsecase) DeleteMember(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.DeleteMember")
	defer span.Finish()

	err := h.MemberRepository.DeleteMemberByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (h *MemberUsecase) AddNewMember(ctx context.Context, member *models.Member) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.AddNewMember")
	defer span.Finish()

	err := h.MemberRepository.AddNewMember(ctx, member)
	if err != nil {
		return err
	}
	return nil
}
