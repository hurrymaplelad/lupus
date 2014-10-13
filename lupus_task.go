package lupus

import (
	"github.com/jingweno/gotask/tasking"
)

// NAME
//    hello - Say hello world
//
// DESCRIPTION
//    Say hello world
//
// OPTIONS
//    --verbose, -v
//        run in verbose mode
func TaskHello(t *tasking.T) {
	t.Log("Hello world")
}