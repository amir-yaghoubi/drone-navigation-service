package http_test

import (
	"bytes"
	"drone_navigation_service/pkg/dns"
	"drone_navigation_service/pkg/dns/mocks"
	handler "drone_navigation_service/pkg/handler/http"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func strP(v string) *string {
	return &v
}

func TestHandleDNSRequest(t *testing.T) {

	// Setup
	gin.SetMode(gin.TestMode)

	mockDns := new(mocks.MockDroneNavigationService)

	router := gin.New()
	handler := handler.NewHandler(router.Group(""), mockDns)
	handler.Setup()

	cases := []struct {
		description            string
		requestBody            string
		mockDnsArguments       []interface{}
		mockDnsReturnArguments []interface{}
		expectedStatusCode     int
		expectedResponse       *string
	}{
		{
			"NonJsonRequest",
			"some random text",
			nil,
			nil,
			http.StatusBadRequest,
			nil,
		},
		{
			"EmptyJsonRequest",
			`{}`,
			nil,
			nil,
			http.StatusBadRequest,
			nil,
		},
		{
			"InvalidFieldValues",
			`{"x": 1, y:"NaN", "z": true, "vel":"23Sx3"}`,
			nil,
			nil,
			http.StatusBadRequest,
			nil,
		},
		{
			"SimpleValues",
			`{"x": "1", "y": "1", "z": "1", "vel": "1"}`,
			[]interface{}{mock.AnythingOfType("*context.emptyCtx"), mock.Anything},
			[]interface{}{&dns.DnsResponse{Location: 4}, nil},
			http.StatusOK,
			strP(`{"loc":4}`),
		},
		{
			"FloatingValues",
			`{"x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0"}`,
			[]interface{}{mock.AnythingOfType("*context.emptyCtx"), dns.DnsRequest{X: 123.12, Y: 456.56, Z: 789.89, Velocity: 20.0}},
			[]interface{}{&dns.DnsResponse{Location: 1389.57}, nil},
			http.StatusOK,
			strP(`{"loc":1389.57}`),
		},
		{
			"IfDnsFailed",
			`{"x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0"}`,
			[]interface{}{mock.AnythingOfType("*context.emptyCtx"), dns.DnsRequest{X: 123.12, Y: 456.56, Z: 789.89, Velocity: 20.0}},
			[]interface{}{nil, errors.New("some random error")},
			http.StatusBadRequest,
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mockDnsArguments != nil && tc.mockDnsReturnArguments != nil {
				mockDns.On("CalculateLocation", tc.mockDnsArguments...).Once().Return(tc.mockDnsReturnArguments...)
			}

			request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tc.requestBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, request)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)

			if tc.expectedResponse != nil {
				assert.Equal(t, *tc.expectedResponse, rr.Body.String())
			}

			mockDns.AssertExpectations(t)
		})
	}

}
