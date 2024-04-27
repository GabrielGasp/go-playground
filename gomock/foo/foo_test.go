package foo_test

import (
	"gomock-playground/foo"
	"gomock-playground/foo/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)

	// By default this expects a single call, but we can change it with Times().
	m.
		EXPECT().
		Bar(2).
		Return(100)

	result := foo.SUT(m)

	assert.Equal(t, 100, result)
}
