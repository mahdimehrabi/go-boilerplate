package service

import (
	entity "bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	mock_logger "bbdk/mock/infrastructure"
	mock_user "bbdk/mock/repository"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		user         *entity.User
		ctx          context.Context
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: nil,
		},
		{
			name: "UserRepoError",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(err)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			err := service.CreateUser(test.user)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}
			loggerMock.EXPECT()
			userRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Create(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil)
	loggerMock := mock_logger.NewMockLogger(ctrl)
	b.ResetTimer()
	service := NewUserService(userRepoMock, loggerMock)
	user := &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"}
	service.CreateUser(user)
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("address_user service-createBatchAddresses takes too long to run")
	}
	loggerMock.EXPECT()
	userRepoMock.EXPECT()
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	user := entity.User{Name: "John Doe", Email: "john@example.com"}
	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		id           uint
		result       *entity.User
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(&user, nil)
				return userRepoMock
			},
			id:     1,
			result: &user,
			error:  nil,
		},
		{
			name: "not found",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(nil, userRepo.ErrNotFound)
				return userRepoMock
			},
			id:     2,
			result: nil,
			error:  userRepo.ErrNotFound,
		},
		{
			name: "repo error",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(nil, err)
				return userRepoMock
			},
			id:     3,
			result: nil,
			error:  err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			result, err := service.GetUserByID(test.id)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}

			if !gomock.Eq(result).Matches(test.result) {
				t.Errorf("result:%v is not equal to %v", result, test.result)
			}
		})
	}
}