package http_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"

	"github.com/uesleicarvalhoo/marsrover/internal/config"
	transport "github.com/uesleicarvalhoo/marsrover/internal/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFullMissionIntegration(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}

	// Setup
	config.Set("HTTP_PORT", "8080")
	server := transport.NewServer()

	go func() {
		_ = server.ListenAndServe()
	}()

	defer server.Close()

	// Criar arquivo de missão válido
	fileContent := "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM"
	expectedResponse := "1 3 N\n5 1 E"

	// Arrange
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "mission.txt")
	require.NoError(t, err)
	_, err = io.Copy(part, strings.NewReader(fileContent))
	require.NoError(t, err)
	writer.Close()

	// Act
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/missions", body)
	require.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, expectedResponse, string(responseBody))
}
