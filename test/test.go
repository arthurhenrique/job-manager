package test

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func AssertEqual(t *testing.T, obj1, obj2 interface{}) {
	if cmp.Diff(obj1, obj2) != "" {
		t.Errorf("obj1 (%v) and obj2 (%v) should be equal", obj1, obj2)
	}
}

func AssertGreaterThan(t *testing.T, size, expected int) {
	if size < expected {
		t.Errorf("size (%d) should be greater than %d", size, expected)
	}
}

func AssertNil(t *testing.T, obj interface{}) {
	val := reflect.ValueOf(obj)
	if !val.IsNil() {
		t.Errorf("obj (%v) should be nil", val)
	}
}

func AssertNotNil(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Errorf("obj (%v) should not be nil", reflect.TypeOf(obj))
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Error("expecting error")
	}
}

func RequireNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func MockHTTP(t *testing.T, handler http.HandlerFunc) {
	mock := httptest.NewUnstartedServer(
		http.HandlerFunc(handler),
	)
	mock.Listener.Close()
	l, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		t.Fatalf("error binding 127.0.0.1:9001 %+v", err)
		return
	}
	mock.Listener = l
	mock.Start()
	t.Cleanup(func() {
		mock.Close()
	})
	time.Sleep(500 * time.Millisecond)
}

type APITestCase struct {
	Name         string
	Route        string
	Method       string
	Status       int
	Payload      string
	BodyContains string
	Headers      http.Header
}

// Run execute test cases
func (tc APITestCase) Run(t *testing.T) {
	var resp *http.Response
	var err error
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(tc.Method, tc.Route, bytes.NewBuffer([]byte(tc.Payload)))
	if err != nil {
		t.Errorf("Error in request %s", err)
	}
	if tc.Headers != nil {
		req.Header = tc.Headers
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("io error calling %s %s", tc.Name, err)
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("body read error calling %s %s", tc.Name, err)
		return
	}

	body := string(bytes)
	if resp.StatusCode != tc.Status {
		t.Errorf("Status error calling %s %d %s", tc.Name, resp.StatusCode, body)
		return
	}

	if resp.StatusCode/100 != 3 {
		if !strings.Contains(body, tc.BodyContains) {
			t.Errorf("unexpected response body in '%s'\n Received '%s'\n Expected that contains '%s'", tc.Name, body, tc.BodyContains)
			return
		}
	}
}
