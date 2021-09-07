package flusher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestFlusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flusher Suite")
}
