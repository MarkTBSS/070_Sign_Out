package middlewaresUsecases

import _pkgMiddlewaresMiddlewaresRepositories "github.com/MarkTBSS/070_Sign_Out/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
