// Declares the package name
package jecconfiguration

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

//	Declare Database Configuration Entity
type DbConfiguration struct {
	Host       string
	Port       string
	Dbname     string
	Dbuser     string
	Dbpassword string
	Sslmode    string
}
