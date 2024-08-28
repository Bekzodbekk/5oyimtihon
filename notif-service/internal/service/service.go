package service

import (
	"context"
	"notif-service/internal/repository"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
)

type Service struct {
	*notif.UnimplementedNotifServiceServer
	repo repository.NotifRepository
}

func NewSerivce(repo *repository.NotifRepository) *Service {
	return &Service{
		repo: *repo,
	}
}

func (s *Service) Notification(ctx context.Context, req *notif.NotificationReq) (*notif.VoidNotif, error) {
	return s.repo.Notification(ctx, req)
}
