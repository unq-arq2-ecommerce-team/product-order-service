package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type InterfaceMocks struct {
	Logger      *MockLogger
	SellerRepo  *MockSellerRepository
	ProductRepo *MockProductRepository
	OrderRepo   *MockOrderRepository
}

// NewInterfaceMocks create an *InterfaceMocks with their mocked interfaces initialized
func NewInterfaceMocks(t *testing.T) *InterfaceMocks {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	baseLogger := NewMockLogger(ctrl)
	setUpLoggerMock(baseLogger)
	return &InterfaceMocks{
		Logger:      baseLogger,
		SellerRepo:  NewMockSellerRepository(ctrl),
		ProductRepo: NewMockProductRepository(ctrl),
		OrderRepo:   NewMockOrderRepository(ctrl),
	}
}

func setUpLoggerMock(logger *MockLogger) {
	logger.EXPECT().WithFields(gomock.Any()).Return(logger).AnyTimes()
	logger.EXPECT().WithRequestId(gomock.Any()).Return(logger).AnyTimes()

	logger.EXPECT().Print(gomock.Any()).AnyTimes()
	logger.EXPECT().Debug(gomock.Any()).AnyTimes()
	logger.EXPECT().Info(gomock.Any()).AnyTimes()
	logger.EXPECT().Warn(gomock.Any()).AnyTimes()
	logger.EXPECT().Error(gomock.Any()).AnyTimes()
	logger.EXPECT().Fatal(gomock.Any()).AnyTimes()
	logger.EXPECT().Panic(gomock.Any()).AnyTimes()

	logger.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Infof(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Errorf(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Fatalf(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Panicf(gomock.Any(), gomock.Any()).AnyTimes()
}
