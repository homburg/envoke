package main_test

import (
	"flag"
	gc "github.com/homburg/goconvey/convey"
	"testing"
)

func TestHyphenArgNotFlag(t *testing.T) {
	gc.Convey("Given a flagset and some args, including a hyphen", t, func() {

		f := &flag.FlagSet{}
		isF := f.Bool("f", false, "")
		isX := f.Bool("x", false, "")
		args := []string{"-f", "-"}

		gc.Convey("When parsing the args", func() {

			f.Parse(args)

			gc.Convey("There should only by one flag parsed", func() {

				gc.So(f.NFlag(), gc.ShouldEqual, 1)

				gc.Convey("And it should be \"-f\"", func() {

					gc.So(*isF, gc.ShouldBeTrue)

					gc.Convey("And not \"-x\"", func() {
						gc.So(*isX, gc.ShouldBeFalse)
					})
				})
			})

			gc.Convey("There should only be one arg parsed", func() {
				gc.So(f.NArg(), gc.ShouldEqual, 1)
			})
		})
	})
}
