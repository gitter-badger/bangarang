// Code generated by ffjson <https://github.com/pquerna/ffjson>
//
// This should be automatically deleted by running 'ffjson',
// if leftover, please delete it.

package alarm

import (
	ffjsonshared "github.com/pquerna/ffjson/shared"
)

func FFJSONExpose() []ffjsonshared.InceptionType {
	rv := make([]ffjsonshared.InceptionType, 0)

	rv = append(rv, ffjsonshared.InceptionType{Obj: Condition{}, Options: ffjsonshared.StructOptions{SkipDecoder: false, SkipEncoder: false}})

	rv = append(rv, ffjsonshared.InceptionType{Obj: StdDev{}, Options: ffjsonshared.StructOptions{SkipDecoder: false, SkipEncoder: false}})

	rv = append(rv, ffjsonshared.InceptionType{Obj: EventTracker{}, Options: ffjsonshared.StructOptions{SkipDecoder: false, SkipEncoder: false}})

	return rv
}