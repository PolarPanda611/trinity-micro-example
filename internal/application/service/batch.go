package service

import (
	"fmt"
	"sync"

	"github.com/PolarPanda611/trinity-micro"
)

func init() {
	trinity.RegisterInstance("BatchService", &sync.Pool{
		New: func() interface{} {
			return new(batchServiceImpl)
		},
	})
}

type BatchService interface {
	Start() error
}

type batchServiceImpl struct {
}

func (s *batchServiceImpl) Start() error {
	fmt.Println("batch service running! ")
	return nil
}
