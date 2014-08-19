package jobs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	bmreljob "github.com/cloudfoundry/bosh-micro-cli/release/jobs"
)

var _ = Describe("Job", func() {
	var job bmreljob.Job

	Describe("FindTemplateByValue", func() {
		Context("when a template with the value exists", func() {
			var expectedTemplate map[string]string

			BeforeEach(func() {
				expectedTemplate = map[string]string{
					"fake-template-name": "fake-template-value",
				}
				job = bmreljob.Job{
					Templates: expectedTemplate,
				}
			})

			It("returns the template and true", func() {
				actualTemplate, ok := job.FindTemplateByValue("fake-template-value")
				Expect(actualTemplate).To(Equal("fake-template-name"))
				Expect(ok).To(BeTrue())
			})
		})

		Context("when the template does not exist", func() {
			It("returns nil and false", func() {
				_, ok := job.FindTemplateByValue("nonsense")
				Expect(ok).To(BeFalse())
			})
		})
	})
})