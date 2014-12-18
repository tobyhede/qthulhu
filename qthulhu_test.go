package qthulhu

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func dbPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}

type tFn func() (bool, error)

type eFn func(error)

func WaitForTrue(f tFn, e eFn) {
	retries := 100

	for retries > 0 {
		time.Sleep(1 * time.Second)
		retries--

		success, err := f()
		if success {
			return
		}

		if retries == 0 {
			e(err)
		}
	}
}

func WaitForLeader(r *Raft) {

	f := func() (bool, error) {
		if leader := r.Leader(); leader != nil {
			return true, nil
		}
		return false, nil
	}
	e := func(err error) {
		log.Fatal("Failed to find leader: %v", err)
	}
	WaitForTrue(f, e)

	return
}

// Test functions From https://github.com/benbjohnson/testing
// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func puts(i interface{}) {
	fmt.Printf("\033[32m%v\033[39m\n", i)
}

func inspect(i interface{}) {
	puts(i)
	fmt.Printf("\033[32m%#v\033[39m\n", i)
}
