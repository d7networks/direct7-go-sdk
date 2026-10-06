package direct7

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const testOTPID = "0012c7f5-2ba5-49db-8901-4ee9be6dc8d1"

type recordedRequest struct {
	method string
	path   string
	body   map[string]interface{}
}

func newVerifyTestClient(t *testing.T) (*Client, *[]recordedRequest) {
	var requests []recordedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := recordedRequest{method: r.Method, path: r.URL.Path}
		raw, _ := ioutil.ReadAll(r.Body)
		if len(raw) > 0 {
			if err := json.Unmarshal(raw, &req.body); err != nil {
				t.Errorf("invalid JSON body: %v", err)
			}
		}
		requests = append(requests, req)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)
	client := NewClient("test-token")
	client.SetHost(server.URL)
	return client, &requests
}

func TestVerifyV1PathsUnchanged(t *testing.T) {
	client, requests := newVerifyTestClient(t)
	noErr(t)(client.verify.SendOTP("SignOTP", "+97150900XXXX", "Your code is: {}", "text", 600, 0))
	noErr(t)(client.verify.ResendOTP(testOTPID))
	noErr(t)(client.verify.VerifyOTP(testOTPID, "1425"))
	noErr(t)(client.verify.GetStatus(testOTPID))

	want := []string{
		"POST /verify/v1/otp/send-otp",
		"POST /verify/v1/otp/resend-otp",
		"POST /verify/v1/otp/verify-otp",
		"GET /verify/v1/report/" + testOTPID,
	}
	assertRoutes(t, *requests, want)
	assertBody(t, (*requests)[0].body, map[string]interface{}{
		"originator": "SignOTP", "recipient": "+97150900XXXX", "content": "Your code is: {}",
		"data_coding": "text", "expiry": float64(600), "template_id": float64(0),
	})
}

func TestVerifyV2Paths(t *testing.T) {
	client, requests := newVerifyTestClient(t)
	noErr(t)(client.verify.V2.SendOTP("+97150900XXXX", "login_flow"))
	noErr(t)(client.verify.V2.ResendOTP(testOTPID))
	noErr(t)(client.verify.V2.VerifyOTP(testOTPID, "1425"))
	noErr(t)(client.verify.V2.GetStatus(testOTPID))

	want := []string{
		"POST /verify/v2/otp/send-otp",
		"POST /verify/v2/otp/resend-otp",
		"POST /verify/v2/otp/verify-otp",
		"GET /verify/v2/report/" + testOTPID,
	}
	assertRoutes(t, *requests, want)
	assertBody(t, (*requests)[0].body, map[string]interface{}{"recipient": "+97150900XXXX", "flow_id": "login_flow"})
	assertBody(t, (*requests)[1].body, map[string]interface{}{"otp_id": testOTPID})
	assertBody(t, (*requests)[2].body, map[string]interface{}{"otp_id": testOTPID, "otp_code": "1425"})
}

func TestVerifyV2InvalidOTPID(t *testing.T) {
	client, requests := newVerifyTestClient(t)
	if _, err := client.verify.V2.ResendOTP("not-a-uuid"); err == nil {
		t.Error("expected error for invalid otp id on ResendOTP")
	}
	if _, err := client.verify.V2.VerifyOTP("not-a-uuid", "1425"); err == nil {
		t.Error("expected error for invalid otp id on VerifyOTP")
	}
	if len(*requests) != 0 {
		t.Errorf("expected no requests, got %d", len(*requests))
	}
}

func noErr(t *testing.T) func(string, error) {
	return func(_ string, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

func assertRoutes(t *testing.T, requests []recordedRequest, want []string) {
	t.Helper()
	got := make([]string, len(requests))
	for i, r := range requests {
		got[i] = r.method + " " + r.path
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("routes mismatch\n got: %v\nwant: %v", got, want)
	}
}

func assertBody(t *testing.T, got, want map[string]interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("body mismatch\n got: %v\nwant: %v", got, want)
	}
}
