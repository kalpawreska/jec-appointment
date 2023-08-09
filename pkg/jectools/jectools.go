// Declares the package name
package jectools

//	Import library
import (
	"database/sql"
	"os"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Get Environtment Value based on Key
func GetEnv(key string) (value string) {
	value = os.Getenv(key)

	return
}

func NewSQLNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
