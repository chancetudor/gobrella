package client

import (
	"fmt"
	"github.com/chancetudor/gobrella/destination"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestUmbrellaClient_GetDestinations(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.GetDestinations("12345")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if r == nil {
		t.Error()
		t.Fail()
	}

	fmt.Println(*r)
}

func TestUmbrellaClient_PostDestinations(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	dests := []*destination.PostDestination{
		destination.NewPostDestination("test"+strconv.Itoa(0)+".com",
			"this is test domain "+strconv.Itoa(0)),
	}
	r, err := testClient.PostDestinations("12345", dests)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if r != 200 {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}

func TestUmbrellaClient_DeleteDestinations(t *testing.T) {
	err := godotenv.Load("/home/chance/dev/gobrella/dev.env")
	if err != nil {
		t.Error(err)
	}
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	destIDs := []int{4059230}
	r, err := testClient.DeleteDestinations("12345", destIDs)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if r != 200 {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}
