package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"nsure/vote/config"
	"nsure/vote/log"
	"nsure/vote/models"
	"reflect"
	"sync"
)

var gdb *gorm.DB
var store models.Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
}

func SharedStore() models.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			log.GetLog().Error("initDb failed",
				zap.Error(err))
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := config.GetConfig()

	url := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.DataSource.User, cfg.DataSource.Password, cfg.DataSource.Addr, cfg.DataSource.Database)
	var err error
	gdb, err = gorm.Open(cfg.DataSource.DriverName, url)
	if err != nil {
		log.GetLog().Error("gorm open failed",
			zap.String("url", url),
			zap.Error(err))
		return err
	} else {
		log.GetLog().Info("Success gorm open")
	}

	gdb.SingularTable(true)
	gdb.DB().SetMaxIdleConns(10)
	gdb.DB().SetMaxOpenConns(50)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "v_" + defaultTableName
	}

	if cfg.DataSource.EnableAutoMigrate {
		var tables = []interface{}{
			&models.Account{},
			&models.Bill{},
			&models.Qualification{},
			&models.Pending{},
			&models.Claim{},
			&models.Apply{},
			&models.Vote{},
			&models.VoteFill{},
			&models.Reward{},
			&models.RewardCha{},
			&models.PunishCha{},
			&models.Payment{},
			&models.Challenge{},
			&models.ChallengeFill{},
			&models.Withdraw{},
			&models.ChainClaim{},
		}
		for _, table := range tables {
			log.GetLog().Info("migrating database",
				zap.Reflect("table", reflect.TypeOf(table)))
			if err = gdb.AutoMigrate(table).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Store) BeginTx() (models.Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
