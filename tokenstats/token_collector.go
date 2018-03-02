package tokenstats

import (
	"context"
)

type TokenCollector interface {
	Collect(ctx context.Context) TokenStats
}
