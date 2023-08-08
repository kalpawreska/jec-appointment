// Declares the package name
package jectools

//	Import library
import "os"

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

//  Get Environtment Value based on Key
func GetEnv(key string) (value string) {
	value = os.Getenv(key)

	return
}
