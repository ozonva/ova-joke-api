package flusher_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flusher Suite")
}
