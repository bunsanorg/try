package scope

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrHello = errors.New("hello")

func TestMust(t *testing.T) {
	assert.Panics(t, func() { Must(ErrHello) })
}

func TestTryNoError(t *testing.T) {
	Try(func() {
		t.Log("Everything is OK")
	}).Catch(func(err error) {
		assert.Fail(t, "This should not be called, no error happened")
	})
}

func TestTryError(t *testing.T) {
	called := false
	Try(func() {
		Must(ErrHello)
	}).Catch(func(err error) {
		called = true
		assert.Equal(t, ErrHello, err)
	})
	assert.True(t, called, "Catch was not called")
}

func TestReturn(t *testing.T) {
	assert.Equal(t, ErrHello, Try(func() {
		Must(ErrHello)
	}).Return())
}

func TestCatchReturn(t *testing.T) {
	called := false
	assert.Equal(t, ErrHello, Try(func() {
		Must(ErrHello)
	}).Catch(func(err error) {
		called = true
		assert.Equal(t, ErrHello, err)
	}).Return())
	assert.True(t, called, "Catch was not called")
}

func TestTryPanic(t *testing.T) {
	assert.Panics(t, func() {
		Try(func() {
			t.Log("Panic!")
			panic("random panic")
		}).Catch(func(err error) {
			assert.Fail(t, "This should not be called, no error happened")
		})
	})
}
