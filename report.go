package usecase

import (
	"context"
	"mini-soccer-cli/entity"
	"mini-soccer-cli/repository"
)

type ReportUsecase struct{ repo repository.ReportRepository }

func NewReportUsecase(repo repository.ReportRepository) *ReportUsecase {
	return &ReportUsecase{repo: repo}
}

func (u *ReportUsecase) Revenue(ctx context.Context) ([]entity.Revenue, error) {
	return u.repo.GetRevenues(ctx)
}

func (u *ReportUsecase) PendingPayments(ctx context.Context) ([]entity.PendingPayment, error) {
	return u.repo.GetPendingPayments(ctx)
}
