package flagtagconfig

import (
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// FlagGetter will populate a struct tagged with "cli" by Implementing TagValueGetter when run through tags.Process
type FlagGetter struct{}

func (g *FlagGetter) TagName() string {
	return "cli"
}

func (g *FlagGetter) Get(key string, f reflect.StructField) (val string) {
	fs := flag.NewFlagSet(key, flag.ErrorHandling(-1))
	fs.SetOutput(ioutil.Discard)
	fs.StringVar(&val, key, "", "")

	fakeOtherOptions(key, fs)

	_ = fs.Parse(os.Args[1:])

	return
}

func fakeOtherOptions(key string, fs *flag.FlagSet) {
	// Setup a dummy to parse all the flags we don't care about on this pass
	var dummy string
	for _, arg := range os.Args {
		arg := strings.Replace(strings.Split(arg, "=")[0], "-", "", -1)
		if arg != key {
			fs.StringVar(&dummy, arg, "", "")
		}
	}
}
