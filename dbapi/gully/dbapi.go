package gully

import (
	"github.com/parthdesai/sc-notifications/db"
)

func GetFromDatabase(dbConn *db.MConn, user string, applicationID string, organization string, ident string) *Gully {
	gully := new(Gully)
	if dbConn.GetOne(GullyCollection, db.M{"app_id": applicationID,
		"org": organization, "user": user, "ident": ident}, gully) != nil {
		return nil
	}
	return gully
}

func DeleteFromDatabase(dbConn *db.MConn, gully *Gully) error {
	return dbConn.Delete(GullyCollection, db.M{"app_id": gully.ApplicationID,
		"org": gully.Organization, "user": gully.User, "ident": gully.Ident})
}

func InsertIntoDatabase(dbConn *db.MConn, gully *Gully) string {
	return dbConn.Insert(GullyCollection, gully)
}

func FindAppropriateGullyForUser(dbConn *db.MConn, user string, applicationID string, organization string, ident string) *Gully {
	var err error
	gully := new(Gully)
	err = dbConn.GetOne(GullyCollection, db.M{
		"user":   user,
		"app_id": applicationID,
		"org":    organization,
		"ident":  ident,
	}, gully)

	if err == nil {
		return gully
	}

	err = dbConn.GetOne(GullyCollection, db.M{
		"user":   user,
		"app_id": applicationID,
		"ident":  ident,
	}, gully)

	if err == nil {
		return gully
	}

	err = dbConn.GetOne(GullyCollection, db.M{
		"user":  user,
		"org":   organization,
		"ident": ident,
	}, gully)

	if err == nil {
		return gully
	}
	return nil
}

func FindAppropriateOrganizationGully(dbConn *db.MConn, applicationID string, organization string, ident string) *Gully {
	var err error
	gully := new(Gully)
	err = dbConn.GetOne(GullyCollection, db.M{
		"app_id": applicationID,
		"org":    organization,
		"ident":  ident,
	}, gully)

	if err == nil {
		return gully
	}

	err = dbConn.GetOne(GullyCollection, db.M{
		"org":   organization,
		"ident": ident,
	}, gully)

	if err == nil {
		return gully
	}

	return nil

}

func FindGlobalGully(dbConn *db.MConn, ident string) *Gully {
	var err error
	gully := new(Gully)
	err = dbConn.GetOne(GullyCollection, db.M{
		"ident": ident,
	}, gully)

	if err == nil {
		return gully
	}

	return nil

}

func FindAppropriateGully(dbConn *db.MConn, user string, applicationID string, organization string, ident string) *Gully {

	var gully *Gully

	gully = FindAppropriateGullyForUser(dbConn, user, applicationID, organization, ident)

	if gully != nil {
		return gully
	}

	gully = FindAppropriateOrganizationGully(dbConn, applicationID, organization, ident)

	if gully != nil {
		return gully
	}

	gully = FindGlobalGully(dbConn, ident)

	if gully != nil {
		return gully
	}

	return nil

}
