package repository

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
)

type NotifRepository interface {
	Notification(ctx context.Context, req *notif.NotificationReq) (*notif.VoidNotif, error)
}
