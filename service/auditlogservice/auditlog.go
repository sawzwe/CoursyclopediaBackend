package auditlogsvc

import (
	"BackendCoursyclopedia/model"
	auditlogrepo "BackendCoursyclopedia/repository/auditlogrepository"
	"context"
)

type IAuditLogService interface {
	GetAllAuditLogs(ctx context.Context) ([]model.AuditLog, error)
}

type AuditLogService struct {
	AuditLogRepository auditlogrepo.IAuditLogRepository
}

func NewAuditLogService(auditlogrepo auditlogrepo.IAuditLogRepository) IAuditLogService {
	return &AuditLogService{
		AuditLogRepository: auditlogrepo,
	}
}

func (s *AuditLogService) GetAllAuditLogs(ctx context.Context) ([]model.AuditLog, error) {
	return s.AuditLogRepository.FindAllAuditLogs(ctx) // Return the result from the repository
}
