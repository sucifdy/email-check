package email_test

import (
	. "email"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Email Validation", func() {
	Context("Valid emails", func() {
		It("should return true for a valid email format", func() {
			Expect(IsValidEmail("example@domain.com")).To(BeTrue())
			Expect(IsValidEmail("user.name+tag@sub.domain.co")).To(BeTrue())
		})
	})

	Context("Invalid emails", func() {
		It("should return false for missing @ symbol", func() {
			Expect(IsValidEmail("exampledomain.com")).To(BeFalse())
		})

		It("should return false for missing domain", func() {
			Expect(IsValidEmail("example@.com")).To(BeFalse())
		})

		It("should return false for missing top-level domain", func() {
			Expect(IsValidEmail("example@domain")).To(BeFalse())
		})

		It("should return false for special characters", func() {
			Expect(IsValidEmail("example@domain!com")).To(BeFalse())
		})
	})
})
