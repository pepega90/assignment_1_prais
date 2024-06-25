package postgres_gorm

import (
	"assignment_1/entity"
	"assignment_1/service"
	"context"
)

type submissionRepository struct {
	db GormDBIface
}

// NewSubmissionRepository membuat instance baru dari submissionRepository
func NewSubmissionRepository(db GormDBIface) service.ISubmissionRepository {
	return &submissionRepository{db: db}
}

func (s *submissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

func (s *submissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

func (s *submissionRepository) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

func (s *submissionRepository) DeleteSubmission(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *submissionRepository) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}
