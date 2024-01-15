package auth

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strconv"
)

func TestAuth(t *testing.T) {
	type test struct {	
		header http.Header
		want string
	}

	tests := make([]test, 5)

	headers := make([]http.Header, 5)

	for i, header := range headers {
		istr := strconv.Itoa(i)
		r := httptest.NewRequest("GET", "http://localhost:8000", nil)
		r.Header.Add("Authorization", "ApiKey " + istr)
		header = r.Header
		tests[i] = test {
			header: header,
			want: istr,
		}
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.header)
		if err != nil {
			t.Errorf("GetAPIKey(%v) = %v", tc.header, err)
		}
		if got != tc.want {
			t.Errorf("GetAPIKey(%v) = %v, want %v", tc.header, got, tc.want)
		}

	}
}
