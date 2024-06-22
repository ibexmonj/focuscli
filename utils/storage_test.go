package utils_test

import (
	"os"
	"time"

	"github.com/ibexmonj/focuscli/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage", func() {
	const dataFile = "sessions.json"

	AfterEach(func() {
		err := os.Remove(dataFile)
		if err != nil {
			return
		}
	})

	Describe("SaveSession", func() {
		It("should save a session correctly", func() {
			session := utils.Session{
				Type:      "focus",
				Duration:  25,
				Timestamp: time.Now(),
			}

			err := utils.SaveSession(session)
			Expect(err).ToNot(HaveOccurred())

			sessions, err := utils.LoadSessions()
			Expect(err).ToNot(HaveOccurred())
			Expect(sessions).To(HaveLen(1))
			Expect(sessions[0].Type).To(Equal(session.Type))
			Expect(sessions[0].Duration).To(Equal(session.Duration))
		})
	})

	Describe("LoadSessions", func() {
		It("should load sessions correctly", func() {
			session1 := utils.Session{
				Type:      "focus",
				Duration:  25,
				Timestamp: time.Now(),
			}
			session2 := utils.Session{
				Type:      "break",
				Duration:  5,
				Timestamp: time.Now(),
			}

			err := utils.SaveSession(session1)
			if err != nil {
				return
			}
			err = utils.SaveSession(session2)
			if err != nil {
				return
			}

			sessions, err := utils.LoadSessions()
			Expect(err).ToNot(HaveOccurred())
			Expect(sessions).To(HaveLen(2))
			Expect(sessions[0].Type).To(Equal(session1.Type))
			Expect(sessions[0].Duration).To(Equal(session1.Duration))
			Expect(sessions[1].Type).To(Equal(session2.Type))
			Expect(sessions[1].Duration).To(Equal(session2.Duration))
		})
	})
})
