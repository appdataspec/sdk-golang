package appdataspec

import (
	"github.com/appdataspec/sdk-golang/util/vos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appdata", func() {
	Context("GlobalPath", func() {
		Context("PROGRAMDATA env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedGlobalPath := "dummyGlobalPath"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `PROGRAMDATA`:
						return expectedGlobalPath
					default:
						return ""
					}
				}

				objectUnderTest := New(fakeVos)

				/* act */
				result := objectUnderTest.GlobalPath()

				/* assert */
				Expect(result).To(Equal(expectedGlobalPath))
			})
		})
		Context("PROGRAMDATA env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: PROGRAMDATA env var required"

				objectUnderTest := New(new(vos.FakeVos))

				/* act */
				var actualPanic interface{}
				defer func() {
					actualPanic = recover()
				}()
				_ = objectUnderTest.GlobalPath()

				/* assert */
				Expect(actualPanic).To(Equal(expectedPanic))
			})
		})
	})
	Context("PerUserPath", func() {
		Context("LOCALAPPDATA env var exists", func() {
			It("should return expected path", func() {
				/* arrange */
				expectedPerUserPath := "dummyHomeDirPath"

				fakeVos := new(vos.FakeVos)
				fakeVos.GetenvStub = func(key string) string {
					switch key {
					case `LOCALAPPDATA`:
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
		Context("LOCALAPPDATA env var doesn't exist", func() {
			It("should panic w/ expected message", func() {
				/* arrange */
				expectedPanic := "Unable to determine per user app data path. Error was: LOCALAPPDATA env var required"

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
