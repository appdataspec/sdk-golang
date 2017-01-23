package path

import (
	"github.com/appdataspec/sdk-golang/util/vos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appdata", func() {
	Context("Global", func() {
		Context("PROGRAMDATA env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedGlobal := "dummyGlobal"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `PROGRAMDATA`:
						return expectedGlobal
					default:
						return ""
					}
				}

				objectUnderTest := NewWithVos(fakeVos)

				/* act */
				result := objectUnderTest.Global()

				/* assert */
				Expect(result).To(Equal(expectedGlobal))
			})
		})
		Context("PROGRAMDATA env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: PROGRAMDATA env var required"

				objectUnderTest := NewWithVos(new(vos.FakeVos))

				/* act */
				var actualPanic interface{}
				defer func() {
					actualPanic = recover()
				}()
				_ = objectUnderTest.Global()

				/* assert */
				Expect(actualPanic).To(Equal(expectedPanic))
			})
		})
	})
	Context("PerUser", func() {
		Context("LOCALAPPDATA env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedPerUser := "dummyHomeDirPath"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `LOCALAPPDATA`:
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
		Context("LOCALAPPDATA env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: LOCALAPPDATA env var required"

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
