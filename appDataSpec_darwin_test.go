package appdataspec

import (
	"github.com/appdataspec/sdk-golang/util/vos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appdata", func() {
	Context("GlobalPath", func() {
		It("should return expected path", func() {
			/* arrange */
			expected := "/Library/Application Support"

			objectUnderTest := New(new(vos.FakeVos))

			/* act */
			result := objectUnderTest.GlobalPath()

			/* assert */
			Expect(result).To(Equal(expected))
		})
	})
	Context("PerUserPath", func() {
		Context("HOME env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedPerUserPath := "dummyHomeDirPath"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `HOME`:
						return expectedPerUserPath
					default:
						return ""
					}
				}

				objectUnderTest := New(fakeVos)

				/* act */
				result := objectUnderTest.PerUserPath()

				/* assert */
				Expect(result).To(Equal(expectedPerUserPath))
			})
		})
		Context("HOME env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: HOME env var required"

				objectUnderTest := New(new(vos.FakeVos))

				/* act */
				var actualPanic interface{}
				defer func() {
					actualPanic = recover()
				}()
				_ = objectUnderTest.PerUserPath()

				/* assert */
				Expect(actualPanic).To(Equal(expectedPanic))
			})
		})
	})
})
