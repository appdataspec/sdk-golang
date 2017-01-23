package path

import (
	"github.com/appdataspec/sdk-golang/util/vos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appdata", func() {
	Context("Global", func() {
		It("should return expected path", func() {
			/* arrange */
			expected := "/var/lib"

			objectUnderTest := New()

			/* act */
			result := objectUnderTest.Global()

			/* assert */
			Expect(result).To(Equal(expected))
		})
	})
	Context("PerUser", func() {
		Context("HOME env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedPerUser := "dummyHomeDirPath"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `HOME`:
						return expectedPerUser
					default:
						return ""
					}
				}

				objectUnderTest := NewWithVos(fakeVos)

				/* act */
				result := objectUnderTest.PerUser()

				/* assert */
				Expect(result).To(Equal(expectedPerUser))
			})
		})
		Context("HOME env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: HOME env var required"

				objectUnderTest := NewWithVos(new(vos.FakeVos))

				/* act */
				var actualPanic interface{}
				defer func() {
					actualPanic = recover()
				}()
				_ = objectUnderTest.PerUser()

				/* assert */
				Expect(actualPanic).To(Equal(expectedPanic))
			})
		})
	})
})
