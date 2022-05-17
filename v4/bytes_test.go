package aidego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThisBytesContains(t *testing.T) {
	assert.True(t, ThisBytesContains(
		[]byte("Integer vitae felis in augue gravida condimentum sed eu nunc"),
		"felis"))
	assert.False(t, ThisBytesContains(
		[]byte("Integer vitae felis in augue gravida condimentum sed eu nunc"),
		"neque"))

	assert.True(t, ThisBytesContains(
		[]byte("Nunc non tortor non eros bibendum ullamcorper ac vel nunc"),
		"bend"))
	assert.False(t, ThisBytesContains(
		[]byte("Nunc non tortor non eros bibendum ullamcorper ac vel nunc"),
		"erat"))
}
