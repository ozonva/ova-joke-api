package ova_joke_api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestOvaJokeApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OvaJokeApi Suite")
}
