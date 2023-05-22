package service

import (
	"beetle/app/im/internal/service/ws"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewImService, ws.NewService)
