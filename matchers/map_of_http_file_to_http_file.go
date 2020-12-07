// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	http "net/http"
)

func AnyMapOfHttpFileToHttpFile() map[http.File]http.File {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(map[http.File]http.File))(nil)).Elem()))
	var nullValue map[http.File]http.File
	return nullValue
}

func EqMapOfHttpFileToHttpFile(value map[http.File]http.File) map[http.File]http.File {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue map[http.File]http.File
	return nullValue
}

func NotEqMapOfHttpFileToHttpFile(value map[http.File]http.File) map[http.File]http.File {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue map[http.File]http.File
	return nullValue
}

func MapOfHttpFileToHttpFileThat(matcher pegomock.ArgumentMatcher) map[http.File]http.File {
	pegomock.RegisterMatcher(matcher)
	var nullValue map[http.File]http.File
	return nullValue
}
