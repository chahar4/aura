package tools

import "errors"

var NotFoundErrDb = errors.New("Record not found")
var ProblemErrDb = errors.New("Database feels sick")
var TokenExpiredErr = errors.New("I don't remeber either")
var TokenNotEqualErr = errors.New("I don't remeber either")
var NewPasswordSameAsOldErr = errors.New("This is a joke for you. That is your old password")
var PasswordHashErr = errors.New("Password is not valid")
var LoginErr = errors.New("Email or Password is wronge")
var EmailErr = errors.New("Email service is not reachable")
