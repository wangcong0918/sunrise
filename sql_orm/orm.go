package sql_orm

import (
	"fmt"
	"time"

	"github.com/wangcong0918/sunrise/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

func (e Engine) GetOrmEngine() (engine *xorm.Engine, err error) {
	if EngineCon.Engine != nil {
		if err := EngineCon.Engine.Ping(); err != nil {
			// 关闭原来的链接
			EngineCon.Engine.Close()

			engine, err := e.createEngine()

			if err != nil {
				log.Logger.Error("create engine err --> ", err.Error())
				return nil, err
			}
			EngineCon.Engine = engine
		}
	} else {
		engine, err := e.createEngine()

		if err != nil {
			log.Logger.Error("create init engine err --> ", err.Error())
			return nil, err
		}

		EngineCon.Engine = engine
	}

	return EngineCon.Engine, nil
}

func (e *Engine) createEngine() (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine(DriverName, DataSourceName)
	if err != nil {
		return nil, err
	}

	pingState := make(chan bool)

	defer close(pingState)
	go func() {
		if err := engine.Ping(); err != nil {
			log.Logger.Error("connection db error --> ", err.Error())
		}
		pingState <- true
	}()

	t := time.AfterFunc(5*time.Second, func() {
		pingState <- false
	})

	select {
	case state := <-pingState:
		if state == false {
			return nil, errors.New("connection db error")
		} else {
			t.Stop()
			goto END
		}
	}

END:

	engine.ShowSQL(true)
	//engine.SetMaxOpenConns(1000)
	engine.SetMaxOpenConns(e.MaxOpenConns)
	engine.SetMaxIdleConns(e.MaxIdleConns)

	// 设置时区
	engine.DatabaseTZ = cstZone // 必须
	engine.TZLocation = cstZone // 必须
	//engine.SetTZLocation(cstZone)
	//engine.SetTZDatabase(cstZone)
	//engine.TZLocation,_ = time.LoadLocation("Asia/Shanghai") // cstZone //
	//engine.SetTZLocation(engine.TZLocation)
	//engine.SetTZDatabase(engine.TZLocation)
	e.State = true

	if err != nil {
		log.Logger.Warning("set orm engine location err --> ", err.Error())
	}

	return engine, nil
}

func (s *ShortEngine) GetShortEngine() (engine *xorm.Engine, err error) {
	ShortDataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		s.User, s.Pwd, s.Host,
		s.Port, s.DbName, s.Charset)
	engine, err = xorm.NewEngine(s.DriverName, ShortDataSourceName)
	if err != nil {
		return nil, err
	}

	pingState := make(chan bool)

	defer close(pingState)
	go func() {
		if err := engine.Ping(); err != nil {
			log.Logger.Error("connection db error --> ", err.Error())
		}
		pingState <- true
	}()

	t := time.AfterFunc(5*time.Second, func() {
		pingState <- false
	})

	select {
	case state := <-pingState:
		if state == false {
			return nil, errors.New("connection db error")
		} else {
			t.Stop()
			goto END
		}
	}

END:
	engine.ShowSQL(true)
	// 设置时区
	engine.DatabaseTZ = cstZone // 必须
	engine.TZLocation = cstZone // 必须

	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.SetTZLocation(engine.TZLocation)
	if err != nil {
		log.Logger.Warning("set orm engine location err --> ", err.Error())
	}

	return engine, nil
}
