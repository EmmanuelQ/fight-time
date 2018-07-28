package fightersdao

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/emmanuelq/lotus/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	COLLECTION = "fighters"
)

type FightersDAO struct {
	Server   string
	Port     string
	Database string
}

var db *mgo.Database
var fightersDAO FightersDAO

func (dao *FightersDAO) Connect() {
	if _, err := toml.DecodeFile("./config/config.toml", &dao); err != nil {
		log.Fatal(err)
	}

	session, err := mgo.Dial(dao.Server)
	if err != nil {
		panic(err)
	}

	db = session.DB(dao.Database)
}

func (dao *FightersDAO) FindByName(name string) ([]fighter.Fighter, error) {
	var fighters []fighter.Fighter
	err := db.C(COLLECTION).Find(bson.M{"FirstName": name}).All(fighters)

	return fighters, err
}

func (dao *FightersDAO) FindByID(id string) (fighter.Fighter, error) {
	var fighter fighter.Fighter
	err := db.C(COLLECTION).FindId(id).One(&fighter)

	return fighter, err
}

func (dao *FightersDAO) Insert(fighter fighter.Fighter) error {
	err := db.C(COLLECTION).Insert(&fighter)

	return err
}

func (dao *FightersDAO) FindAll() ([]fighter.Fighter, error) {
	var fighters []fighter.Fighter
	err := db.C(COLLECTION).Find(bson.M{}).All(&fighters)

	return fighters, err
}
