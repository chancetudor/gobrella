package client

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
	"testing"
)

func TestNewManagementURL(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/umbrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	expectedURL := "https://client.api.umbrella.com/v1/organizations/" + os.Getenv("ID")

	if !strings.EqualFold(testClient.BaseURL.String(), expectedURL) {
		t.Errorf("URL was incorrect, got: %s and wanted: %s", testClient.BaseURL.String(), expectedURL)
	}
}

func TestGet(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/umbrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	response, err := testClient.get("destinationlists")
	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Errorf("Error: Response is nil")
	}
}

// func TestPost(t *testing.T) {
// 	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("PWD"), os.Getenv("ID"))
// 	response, err := testClient.post("destinationlists", []byte("body: this is a test"))
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	if response == nil {
// 		t.Errorf("Error: Response is nil")
// 	}
// }

func TestFormURL(t *testing.T) {
	const DefaultManagementURL = "https://client.api.umbrella.com/v1"
	expectedURL := "https://client.api.umbrella.com/v1/test/path12345"
	actualURL, _ := formURL(DefaultManagementURL, "test", "path12345")

	if !strings.EqualFold(expectedURL, actualURL.String()) {
		t.Errorf("URL was incorrect, got: %s and wanted: %s", actualURL, expectedURL)
	}
}

func TestNewManagementWithClient(t *testing.T) {

}
