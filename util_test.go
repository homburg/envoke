package main

import (
	. "github.com/homburg/goconvey/convey"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	Convey("Given os.Environ()", t, func() {
		env := []string{"USER=thomas", "UID=1000", "GID=1000", "SOMETHING=ELSE=ENTIRELY"}

		Convey("We can construct a map[string]map (environment) of the values", func() {

			So(
				getEnvironment(env),
				ShouldResemble,
				environment{
					"USER":      "thomas",
					"UID":       "1000",
					"GID":       "1000",
					"SOMETHING": "ELSE=ENTIRELY",
				},
			)

			So(
				getEnvironment(env),
				ShouldNotResemble,
				environment{
					"USER":      "thomas",
					"UID":       "1000",
					"GID":       "1000",
					"SOMETHING": "ELSE",
				},
			)
		})
	})
}
