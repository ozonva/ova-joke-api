package ova_joke_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOvaJokeApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OvaJokeApi Suite")
}
