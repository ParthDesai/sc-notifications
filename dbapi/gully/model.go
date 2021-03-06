package gully

import (
	"github.com/changer/khabar/db"
	"github.com/changer/khabar/dbapi"
)

const (
	GullyCollection = "gullys"
)

type Gully struct {
	db.BaseModel `bson:",inline"`
	User         string                 `json:"user" bson:"user"`
	Organization string                 `json:"org" bson:"org"`
	AppName      string                 `json:"app_name" bson:"app_name"`
	Data         map[string]interface{} `json:"data" bson:"data" required:"true"`
	Ident        string                 `json:"ident" bson:"ident" required:"true"`
}

func (self *Gully) IsValid(op_type int) bool {
	if (len(self.User) == 0) && (len(self.Organization) == 0) && (len(self.AppName) == 0) {
		return false
	}

	if len(self.Ident) == 0 {
		return false
	}

	if op_type == dbapi.INSERT_OPERATION {
		if len(self.Data) == 0 {
			return false
		}

	}

	return true
}
