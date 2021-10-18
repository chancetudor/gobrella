package client

import (
	"strings"
	"testing"
)

func TestNewManagementURL(t *testing.T) {
	testClient := NewUmbrellaClient("testKey", "testPwd", "12345")
	expectedURL := "https://client.api.umbrella.com/v1/organizations/12345"

	if !strings.EqualFold(testClient.BaseURL.String(), expectedURL) {
		t.Errorf("URL was incorrect, got: %s and wanted: %s", testClient.BaseURL.String(), expectedURL)
	}
}

func TestGet(t *testing.T) {
	testClient := NewUmbrellaClient("testKey", "testPwd", "12345")
	_, response, err := testClient.get("destinationlists")
	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Errorf("Error: Response is nil")
	}
}

func TestPost(t *testing.T) {
	testClient := NewUmbrellaClient("testKey", "testPwd", "12345")
	_, response, err := testClient.post("destinationlists", []byte("body: this is a test"))
	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Errorf("Error: Response is nil")
	}
}

func TestNewManagementWithClient(t *testing.T) {

}
