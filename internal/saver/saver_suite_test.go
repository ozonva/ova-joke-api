package saver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver Suite")
}
