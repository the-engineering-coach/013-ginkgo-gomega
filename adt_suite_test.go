package adt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestADTSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ADT Suite")
}

