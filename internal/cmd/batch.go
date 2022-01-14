// Author: Daniel TAN
// Date: 2021-08-18 00:34:11
// LastEditors: Daniel TAN
// LastEditTime: 2021-10-02 01:24:18
// FilePath: /trinity-micro/example/crud/cmd/batch.go
// Description:
/*
 * @Author: Daniel TAN
 * @Date: 2021-08-18 00:34:11
 * @LastEditors: Daniel TAN
 * @LastEditTime: 2021-09-09 00:08:47
 * @FilePath: /trinity-micro/example/cmd/batch.go
 * @Description:
 */
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"trinity-micro-api/internal/adapter/controller"
	"trinity-micro-api/internal/application/model"
	"trinity-micro-api/internal/config"
	"trinity-micro-api/internal/consts"

	"github.com/BurntSushi/toml"
	"github.com/PolarPanda611/trinity-micro"
	"github.com/PolarPanda611/trinity-micro/core/dbx"
	"github.com/PolarPanda611/trinity-micro/core/logx"
	"github.com/PolarPanda611/trinity-micro/core/tracerx"
	"github.com/PolarPanda611/trinity-micro/core/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	batchCmd = &cobra.Command{
		Use:   consts.BatchCmdString,
		Short: fmt.Sprintf("starting the %v service for %v", consts.BatchCmdString, consts.ProjectName),
		Long:  fmt.Sprintf("This is the %v service for %v", consts.BatchCmdString, consts.ProjectName),
		Run:   RunBatch,
	}
)

func init() {
	rootCmd.AddCommand(batchCmd)
}

func RunBatch(cmd *cobra.Command, args []string) {
	serviceName := fmt.Sprintf("%v-%v", consts.ProjectName, consts.BatchCmdString)

	// infra set up
	logger := logx.Init(logx.Config{
		ServiceName: serviceName,
		LogfilePath: fmt.Sprintf("%v.log", serviceName),
	})
	log.Printf("%v-%v service starting!\n", consts.ProjectName, consts.BatchCmdString)
	logger.Infof("%v-%v service starting!", consts.ProjectName, consts.BatchCmdString)
	defer func() {
		log.Printf("%v-%v service shutdown!\n", consts.ProjectName, consts.BatchCmdString)
		logger.Infof("%v-%v service shutdown!", consts.ProjectName, consts.BatchCmdString)
	}()

	currentPath, _ := os.Getwd()
	configPath := filepath.Join(currentPath + "/conf/config.toml")
	if _, err := toml.DecodeFile(configPath, &config.Conf); err != nil {
		logger.Fatalf("load config :%v failed, err: %v", configPath, err)
	}
	logger.Infof("load config: %v successfully", config.Conf)

	dbx.Init(&dbx.Config{
		Type:        config.Conf.Database.Type,
		DSN:         config.Conf.Database.DSN,
		TablePrefix: config.Conf.Database.TablePrefix,
		MaxIdleConn: config.Conf.Database.MaxIdleConn,
		MaxOpenConn: config.Conf.Database.MaxOpenConn,
		Logger:      logger.WithField("app", "database"),
	})
	// handle multi tenant initialize
	{
		tenants := make([]string, 0)
		sessionDB := dbx.DB.Session(&gorm.Session{
			NewDB: true,
		})
		sessionDB.AutoMigrate(&model.Tenant{})
		sessionDB.FirstOrCreate(&model.Tenant{
			Model: dbx.Model{
				ID: utils.GetSnowflakeID(),
			},
			Name: "tn_01",
		})
		var res []model.Tenant
		if err := sessionDB.Find(&res).Error; err != nil {
			logger.Fatalf("list tenant failed, err: ", err)
		}
		for _, tenant := range res {
			tenants = append(tenants, fmt.Sprintf("tn_%d", tenant.ID))
		}
		dbx.Migrate(context.Background(), tenants...)
	}

	tracerx.Init(tracerx.Config{
		Type:        config.Conf.Tracer.Type,
		ServiceName: config.Conf.Tracer.ServiceName,
		Host:        config.Conf.Tracer.Host,
	})
	t := trinity.New(trinity.Config{
		Logger: logger,
	})
	if err := t.Run("BatchController", func(ctx context.Context, c controller.BatchController) error {
		return c.Start()
	}); err != nil {
		log.Fatalf("err: %v", err)
	}

}
