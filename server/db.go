package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/tukangkod/go-gobok/utils"
	"github.com/tukangkod/go-gobok/tm"
	"github.com/spf13/viper"
)

var (
	db *pg.DB
)

type TagMsg struct {
	ID       int
	ClientIp string
	ServerIp string
	Tags     map[string]string
	Message  string
}

// DB connection
func DBConnect() {
	utils.Log.Infof(utils.LogTemplate(), utils.GetFunctionName(DBConnect), "Start DB Connection")

	username := viper.GetString("username")
	password := viper.GetString("password")
	database := viper.GetString("database")

	utils.Log.Infof("[%s] : Parameter : user: %s, pass %s, db: %s", utils.GetFunctionName(DBConnect), username, password, database)

	db = pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: database,
	})

	// todo - fatal if connection failed
}

// Create table based on TagMsg struct
func CreateTable() {
	utils.Log.Infof(utils.LogTemplate(), utils.GetFunctionName(CreateTable), "Creating Table")

	err := db.CreateTable(&TagMsg{}, &orm.CreateTableOptions{});
	if err != nil {
		utils.Log.Warnf(utils.ErrTemplate(), utils.GetFunctionName(CreateTable), err)
		// todo - if table already created, show as info.
	}
}

// save tag msg into database
func SaveTagMsg(msg *tm.PutRequest) error {
	fnName := utils.GetFunctionName(SaveTagMsg)
	utils.Log.Infof(utils.LogTemplate(), fnName, "Saving TagMsg: ", msg)

	err := db.Insert(&TagMsg{ClientIp: msg.ClientIp, ServerIp: msg.ServerIp, Tags: msg.Tags, Message: msg.Message})
	if err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), fnName, err)
		return err
	}

	utils.Log.Infof(utils.LogTemplate(), fnName, "Saved")
	return nil
}

func SearchTagMsg(msg *tm.SearchRequest) ([]TagMsg, error) {
	fnName := utils.GetFunctionName(SearchTagMsg)

	var tagmsg []TagMsg

	err := db.Model(&tagmsg).
	WhereGroup(func(q *orm.Query) (*orm.Query, error) {
		q = q.WhereOr("client_ip = ?", msg.ClientIp).
			WhereOr("server_ip = ?", msg.ServerIp).
			WhereOr("tags @> ?", msg.Tags)
		return q, nil
	}).Select()

	if err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), fnName, err)
		return nil, err
	}

	utils.Log.Infof(utils.LogTemplate(), fnName, "Query Success for data :", msg)
	return tagmsg, nil
}
