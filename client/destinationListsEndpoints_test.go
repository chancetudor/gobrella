package client

import (
	"fmt"
	destList "github.com/chancetudor/gobrella/destinationLists"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestUmbrellaClient_GetDestinationLists(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
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
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.GetDestinationList("12345")
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
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	newList := &destList.DestinationListCreate{
		Access:       "allow",
		IsGlobal:     false,
		Name:         "TEST LIST 4",
		Destinations: nil,
	}
	r, err := testClient.PostDestinationList(newList)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}

func TestUmbrellaClient_PatchDestinationList(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.PatchDestinationList("12345", "TEST LIST 4.1")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}

func TestUmbrellaClient_DeleteDestinationList(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.DeleteDestinationList("12345")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}
