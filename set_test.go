package adt_test

import (
	"adt"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

var s *adt.Set
var _ = Describe("Set", func() {

	BeforeEach(func() {
		s = adt.NewSet()
	})


	Describe("Emptiness", func() {
		Context("When the set does not contain items", func() {
			It("Should be empty", func() {
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})

		Context("When the set contains items", func() {
			It("Should not be empty", func() {
				s.Add("red")

				Expect(s.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return an increasing size", func() {
				By("Empty set size being 0", func() {
					Expect(s.Size()).To(BeZero())
				})

				By("Adding a first item", func() {
					s.Add("red")

					Expect(s.Size()).To(Equal(1))
				})

				By("Adding a second item", func() {
					s.Add("blue")

					Expect(s.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When red has not been added", func() {
			It("Should not contain red", func() {
				Expect(s.Contains("red")).To(BeFalse())
			})
		})

		Context("When red has been adding", func() {
			It("Should contain red", func() {
				s.Add("red")

				Expect(s.Contains("red")).To(BeTrue())
			})
		})
	})
})
