package database

import "time"

type Revoke struct {
	Token     string    `json:"token"`
	RevokedAt time.Time `json:"revoked_at"`
}

func (db *DB) RevokeToken(token string) error {
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	revoked := Revoke{
		Token:     token,
		RevokedAt: time.Now().UTC(),
	}
	dbStruct.Revoked[token] = revoked

	err = db.writeDB(dbStruct)
	if err != nil {
		return err
	}
	return nil
}
