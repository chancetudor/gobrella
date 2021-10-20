package client

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestUmbrellaClient_GetDestinationLists(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/umbrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.GetDestinationLists()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(r)
}
