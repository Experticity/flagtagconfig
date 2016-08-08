package flagtagconfig_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/Experticity/flagtagconfig"
	"github.com/Experticity/tagconfig"
	"github.com/stretchr/testify/assert"
)

type Options struct {
	DevMode bool   `cli:"dev.mode"`
	Addr    string `cli:"addr"`
	Num     int    `cli:"num"`
}

var (
	addr      = ":2345"
	num       = 25
	testFlags = []string{"appName", "-addr=" + addr, "-dev.mode=true", fmt.Sprintf("-num=%d", num), "-blah", "hey"}
)

func TestFlagGet(t *testing.T) {
	os.Args = testFlags
	opts := &Options{}

	err := tagconfig.Process(&FlagGetter{}, opts)
	assert.NoError(t, err)

	assert.Equal(t, opts.Addr, addr)
	assert.True(t, opts.DevMode)
	assert.Equal(t, opts.Num, num)
}
