package controller

import (
	"sync"
	"trinity-micro-api/internal/application/service"

	"github.com/PolarPanda611/trinity-micro"
)

func init() {
	BatchControllerPool := &sync.Pool{
		New: func() interface{} {
			return new(batchControllerImpl)
		},
	}
	trinity.RegisterInstance("BatchController", BatchControllerPool)
}

type BatchController interface {
	Start() error
}

type batchControllerImpl struct {
	BatchSrv service.BatchService `container:"autowire:true;resource:BatchService"`
}

func (c *batchControllerImpl) Start() error {
	return c.BatchSrv.Start()
}
