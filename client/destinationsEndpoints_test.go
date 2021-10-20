package client

import (
	"fmt"
	"github.com/chancetudor/umbrella/destination"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestUmbrellaClient_GetDestinations(t *testing.T) {
	_ = godotenv.Load("/home/chance/dev/umbrella/dev.env")
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	r, err := testClient.GetDestinations("4092066")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(*r)
}

func TestUmbrellaClient_PostDestinations(t *testing.T) {
	_ = godotenv.Load("/home/chance/dev/umbrella/dev.env")
	testClient, _ := NewUmbrellaClient(os.Getenv("KEY"), os.Getenv("P"), os.Getenv("ID"))
	dests := []*destination.PostDestination{
		destination.NewPostDestination("test"+strconv.Itoa(0)+".com",
			"chtudor - this is test domain "+strconv.Itoa(0)),
	}
	r, err := testClient.PostDestinations("4092066", dests)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(r)
}
