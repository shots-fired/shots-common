package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shots-fired/shots-common/models"
)

var _ = Describe("Streamers", func() {
	It("should marshal Streamer properly", func() {
		streamer := models.Streamer{
			Name:    "something",
			Status:  "online",
			Viewers: 5,
		}
		res, err := streamer.MarshalBinary()

		Expect(err).To(BeNil())
		Expect(res).To(Equal([]byte("{\"name\":\"something\",\"status\":\"online\",\"viewers\":5}")))
	})
})
