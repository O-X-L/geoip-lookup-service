package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func TestHandleGeoIPLookup_MissingParams(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGeoIPLookup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetClientIP(t *testing.T) {
	cnf.CLIENT_IP_FWD_HDR = false
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "10.0.0.1:44444"

	ip, err := getClientIP(req)
	if err != nil || ip != "10.0.0.1" {
		t.Errorf("Expected 10.0.0.1, got %s", ip)
	}

	// Test X-Forwarded-For
	req.Header.Set("X-Forwarded-For", "1.1.1.1, 2.2.2.2")
	ip, err = getClientIP(req)
	if err != nil || ip != "10.0.0.1" {
		t.Errorf("Expected 10.0.0.1, got %s", ip)
	}

	cnf.CLIENT_IP_FWD_HDR = true
	defer func() { cnf.CLIENT_IP_FWD_HDR = false }()
	req.Header.Set("X-Forwarded-For", "1.1.1.1, 2.2.2.2")
	ip, err = getClientIP(req)
	if err != nil || ip != "2.2.2.2" {
		t.Errorf("Expected 2.2.2.2, got %s", ip)
	}

	// Test X-Real-IP
	req.Header.Del("X-Forwarded-For")
	req.Header.Set("X-Real-IP", "3.3.3.3")
	ip, _ = getClientIP(req)
	if ip != "3.3.3.3" {
		t.Errorf("Expected 3.3.3.3, got %s", ip)
	}
}
