package handler_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/uesleicarvalhoo/marsrover/internal/http/handler"
	"github.com/uesleicarvalhoo/marsrover/orchestrator/mocks"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

func TestMissionFromFile(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about          string
		method         string
		fileContent    string
		mockSetup      func(m *mocks.MissionUseCaseMock)
		expectedStatus int
		expectedBody   string
	}{
		{
			about:       "Successful POST with valid file",
			method:      http.MethodPost,
			fileContent: "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
			mockSetup: func(m *mocks.MissionUseCaseMock) {
				m.On("Execute", mock.Anything, mock.Anything).
					Return([]rover.Params{
						{Name: "Rover-1", Coordinates: rover.Coordinates{X: 1, Y: 3}, Direction: rover.North},
						{Name: "Rover-2", Coordinates: rover.Coordinates{X: 5, Y: 1}, Direction: rover.East},
					}, nil).
					Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "1 3 N\n5 1 E",
		},
		{
			about:          "Missing file returns 400",
			method:         http.MethodPost,
			fileContent:    "",
			mockSetup:      nil,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "mission data is empty",
		},
		{
			about:  "Execute returns domain error",
			method: http.MethodPost,
			fileContent: `5 5
1 2 N
INVALID
`,
			mockSetup:      func(m *mocks.MissionUseCaseMock) {},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "invalid command",
		},
		{
			about:          "GET method returns 405",
			method:         http.MethodGet,
			fileContent:    "",
			mockSetup:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "method not allowed",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			mockSvc := &mocks.MissionUseCaseMock{}
			if tc.mockSetup != nil {
				tc.mockSetup(mockSvc)
			}

			handlerFunc := handler.MissionFromFile(mockSvc)

			var req *http.Request
			if tc.method == http.MethodPost {
				body := bytes.NewBuffer([]byte{})

				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("file", "mission.txt")

				require.NoError(t, err)
				_, err = io.Copy(part, strings.NewReader(tc.fileContent))
				require.NoError(t, err)
				writer.Close()

				req = httptest.NewRequest(tc.method, "/missions", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
			} else {
				req = httptest.NewRequest(tc.method, "/missions", nil)
			}

			w := httptest.NewRecorder()

			// Act
			handlerFunc.ServeHTTP(w, req)

			// Assert
			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			assert.Contains(t, string(body), tc.expectedBody)

			mockSvc.AssertExpectations(t)
		})
	}
}
