package appInit

import (
	"context"
	productServire "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
	"golang.org/x/sync/errgroup"
)

func (a *App) ProductControllerGetterInit() (err error) {
	// if already initialized
	if a.ProductController != nil {
		return nil
	}

	//// product controller
	//a.ProductController = apiController.NewServer(
	//	a.ProductUseCase,
	//	productServire.UnimplementedProductServiceServer{},
	//)

	return nil
}

func (a *App) ProductControllerGetterRun(ctx context.Context) (err error) {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startHTTP(ctx)
	})

	//grp.Go(func() error {
	//	return a.startGRPC(ctx, a.ProductController)
	//})

	return grp.Wait()

	return err
}

func (a *App) startGRPC(ctx context.Context, server productServire.ProductServiceServer) error {
	return nil
}

func (a *App) startHTTP(ctx context.Context) error {
	return nil
}
