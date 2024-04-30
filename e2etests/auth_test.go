package e2etests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"apiServer/internal/dto"
)

var endPointUrl string

func TestMain(m *testing.M) {
	// Context to control the lifecycle of the containers
	ctx := context.Background()

	headers := make(map[string]string)

	spaRequest := testcontainers.ContainerRequest{
		Image:        "apiservergolangtest:latest",
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForHTTP("/healthcheck").WithPort("3000/tcp").WithHeaders(headers),
	}

	// Start SPA container
	spaContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: spaRequest,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start SPA container: %v", err)
	}

	endPoint, err := spaContainer.PortEndpoint(ctx, "3000", "")
	endPointUrl = endPoint
	if err != nil {
		log.Fatalf("Failed to get endpoint from SPA container: %v", err)
	}
	defer spaContainer.Terminate(ctx)

	// Login and get the token to be used in successive api calls

	// Run tests
	code := m.Run()

	// Ensure cleanup
	os.Exit(code)
}

func TestSignup(t *testing.T) {

	// We will call the signup api with request body and response body

	payload := []byte(`{
        "email": "spp.user1@pos.com.my",
        "password": "uASmwQx5dkyxB2ra",
         "name":"shubham"
    }`)
	requestUrl := fmt.Sprintf("http://%s/user", endPointUrl)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(payload))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	assert.Nil(t, err)
	var responseStruct dto.Response
	err = json.Unmarshal(body, &responseStruct)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "user created successfully", responseStruct.Message)

}
