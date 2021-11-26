package adapters

import (
	"context"
	"rosenchat/src/logger"
)

// implCleanup implements IAdapter for the several cleanup operations of the whole Application.
type implCleanup struct{}

func (i *implCleanup) Name() string {
	return "Cleanup"
}

func (i *implCleanup) Start(ctx context.Context) error {
	return nil
}

func (i *implCleanup) Stop(ctx context.Context) error {
	// Logger cleanup.
	logger.Get().Close(ctx)

	/* Add more cleanups below (database connection closures etc). */

	return nil
}

func (i *implCleanup) init() {}
