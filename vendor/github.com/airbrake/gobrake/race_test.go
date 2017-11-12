package gobrake_test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	. "github.com/onsi/ginkgo"

	"github.com/airbrake/gobrake"
)

var _ = Describe("Notifier", func() {
	var notifier *gobrake.Notifier

	BeforeEach(func() {
		handler := func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"id":"123"}`))
		}
		server := httptest.NewServer(http.HandlerFunc(handler))

		notifier = gobrake.NewNotifier(1, "key")
		notifier.SetHost(server.URL)
	})

	It("is race free", func() {
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				notifier.Notify("hello", nil)
			}()
		}
		wg.Wait()

		notifier.Flush()
	})
})
