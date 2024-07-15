package system

import (
	"context"
	"log/slog"
	"time"
)

func StartTask(ctx context.Context, task func(context.Context) error, interval time.Duration, name string) {
	slog.Info("Starting task", "name", name, "interval", interval)

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			go func() {
				err := task(ctx)
				if err != nil {
					slog.Error("Error on task", "name", name, "error", err)
				}
			}()
		}
	}
}
