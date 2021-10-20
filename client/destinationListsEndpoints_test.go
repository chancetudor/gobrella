package client

import (
	"fmt"
	destList "github.com/chancetudor/gobrella/destinationLists"
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

func TestUmbrellaClient_GetDestinationList(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/umbrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.GetDestinationList("4092066")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if r == nil {
		t.Error()
		t.Fail()
	}

	fmt.Println(r)
}

func TestUmbrellaClient_PostDestinationList(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/umbrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	newList := &destList.DestinationListCreate{
		Access:       "allow",
		IsGlobal:     false,
		Name:         "SOAR DEV TEST LIST 2",
		Destinations: nil,
	}
	r, err := testClient.PostDestinationList(newList)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if r.StatusCode != 200 {
		t.Error()
		t.Fail()
	}

	fmt.Println(r)
}
