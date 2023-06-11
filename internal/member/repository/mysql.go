package repository

import (
	"context"
	"errors"
	"social_media/internal/member/models"
	"social_media/pkg/utils"

	"gorm.io/gorm"
)

type MemberRepository interface {
	AddNewMember(ctx context.Context, member *models.Member) error
	UpdateMemberByID(ctx context.Context, member *models.Member, id int) error
	GetMemberByID(ctx context.Context, id int) (*models.Member, error)
	GetAllMembers(ctx context.Context) ([]*models.Member, error)
	DeleteMemberByID(ctx context.Context, id int) error
}

type MySQLRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetMemberByID(ctx context.Context, id int) (*models.Member, error) {
	var member models.Member
	err := r.db.First(&member, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(utils.MemberNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MySQLRepository) GetAllMembers(ctx context.Context) ([]*models.Member, error) {
	var members []*models.Member
	err := r.db.Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *MySQLRepository) AddNewMember(ctx context.Context, member *models.Member) error {
	// Check if username already exists
	existingMember, err := r.GetMemberByUsername(ctx, member.Username)
	if err != nil {
		return err
	}
	if existingMember != nil {
		return errors.New(utils.UsernameAlreadyExists)
	}

	// Create the new member
	return r.db.Create(member).Error
}

func (r *MySQLRepository) GetMemberByUsername(ctx context.Context, username string) (*models.Member, error) {
	var member models.Member
	err := r.db.Where("username = ?", username).First(&member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Username does not exist
	}
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MySQLRepository) UpdateMemberByID(ctx context.Context, member *models.Member, id int) error {
	existingMember, err := r.GetMemberByID(ctx, id)
	if err != nil {
		return err
	}
	if existingMember == nil {
		return errors.New(utils.MemberNotFound)
	}

	// Perform the update
	return r.db.Model(&models.Member{}).Where("id = ?", id).Updates(member).Error
}

func (r *MySQLRepository) DeleteMemberByID(ctx context.Context, id int) error {
	existingMember, err := r.GetMemberByID(ctx, id)
	if err != nil {
		return err
	}
	if existingMember == nil {
		return errors.New(utils.MemberNotFound)
	}

	// Perform the delete operation
	return r.db.Delete(&models.Member{}, id).Error
}
