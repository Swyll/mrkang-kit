package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

//MyEngine mysql引擎
type MyEngine struct {
	UserName     string `ini:"username"`
	PassWord     string `ini:"userpwd"`
	Host         string `ini:"ip"`
	Port         string `ini:"port"`
	DbName       string `ini:"dbname"`
	Timeout      string `ini:"timeout"`
	ReadTimeout  string `ini:"readrimeout"`
	MaxIdleConns int    `ini:"maxidleconns"`
	MaxOpenConns int    `ini:"maxopenconns"`

	db *sql.DB `comment:"自带连接池"`
}

type Opt func(e *MyEngine)

func WithConf(my *MyEngine) Opt {
	return func(e *MyEngine) {
		e.UserName = my.UserName
		e.PassWord = my.PassWord
		e.Host = my.Host
		e.Port = my.Port
		e.DbName = my.DbName
		e.Timeout = my.Timeout
		e.ReadTimeout = my.ReadTimeout
		e.MaxIdleConns = my.MaxIdleConns
		e.MaxOpenConns = my.MaxOpenConns
	}
}

//NewEngine 新建MyEngine
func NewEngine(opts ...Opt) *MyEngine {
	e := &MyEngine{
		UserName:     "root",
		PassWord:     "",
		Host:         "127.0.0.1",
		Port:         "3306",
		DbName:       "energy",
		Timeout:      "5",
		ReadTimeout:  "10",
		MaxIdleConns: 500,
		MaxOpenConns: 500,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

//Init 连接mysql
func (e *MyEngine) Init() error {
	db, err := sql.Open("mysql", e.formatDSN())
	if err != nil {
		return errors.WithStack(err)
	}
	e.db = db
	e.db.SetMaxIdleConns(e.MaxIdleConns)
	e.db.SetMaxOpenConns(e.MaxOpenConns)

	//err = e.InsertDefaultInfo()
	//if err != nil {
	//	return err
	//}

	return nil
}

func (e *MyEngine) formatDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%ss&readTimeout=%ss", e.UserName, e.PassWord, e.Host, e.Port, e.DbName, e.Timeout, e.ReadTimeout)
}

//GetDB 获取db
func (e *MyEngine) GetDB() *sql.DB {
	return e.db
}

func (e *MyEngine) Exec(sql string, params ...interface{}) error {
	db := e.db

	tx, err := db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	return nil
}
