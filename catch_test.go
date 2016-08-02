package try

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrHello = errors.New("hello")

func TestMust(t *testing.T) {
	assert.Panics(t, func() { Must(ErrHello) })
}

func TestThisNoError(t *testing.T) {
	This(func() {
		t.Log("Everything is OK")
	}).Catch(func(err error) {
		assert.Fail(t, "This should not be called, no error happened")
	})
}

func TestThisError(t *testing.T) {
	called := false
	This(func() {
		Must(ErrHello)
	}).Catch(func(err error) {
		called = true
		assert.Equal(t, ErrHello, err)
	})
	assert.True(t, called, "Catch was not called")
}

func TestThisPanic(t *testing.T) {
	assert.Panics(t, func() {
		This(func() {
			t.Log("Panic!")
			panic("random panic")
		}).Catch(func(err error) {
			assert.Fail(t, "This should not be called, no error happened")
		})
	})
}
