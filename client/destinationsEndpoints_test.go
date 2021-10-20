package client

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestUmbrellaClient_GetDestinations(t *testing.T) {
	_ = godotenv.Load("/home/chance/dev/umbrella/dev.env")
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("PWD"), os.Getenv("ID"))
	r, err := testClient.GetDestinations("4092066")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(r.Destination)
}
