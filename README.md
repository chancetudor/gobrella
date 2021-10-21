# gobrella

gobrella is a client library for the [APIs](https://docs.umbrella.com/developer) provided by Cisco Umbrella, formerly OpenDNS. It currently supports endpoints dealing with [Destination Lists](https://developer.cisco.com/docs/cloud-security/#!destination-lists-overview), but further functionality will be added in the future. Contributions are welcome.

### Library

Import the library.

```go
import "github.com/chancetudor/gobrella"
```

Create a client for the Management and Provisioning API. Requires an API key and password, as well as an organization ID. Please follow the Cisco docs for finding this information.

```go
client := umbrella.NewUmbrellaClient(key, pwd, organizationID)
```

The NewUmbrellaClient's HTTP client is Go's default client, but a custom HTTP client can be passed in like so:

```go
client := umbrella.NewUmbrellaClient(key, pwd, organizationID, 
	WithClient(PointerToCustomHTTPClient))
```

Retrieve all destination lists of organization:

```go
lists, err := client.GetDestinationLists()
if err != nil {
  return err
}
for index, value := range lists.Data {
  // do something
}
```

Retrieve a single destination list of an organization:

```go
list, err := client.GetDestinationList("12345")
// do something with list.Data
```

Create a new destination list for an organization:

```go
list, err := client.PostDestinationList(newList)
// do something with list.Data
```

Rename a destination list:

```go
statusCode, err := client.PatchDestinationList(listID, newName)
// check statusCode
```

Delete a destination list:

```go
statusCode, err := client.DeleteDestinationList(listID)
// check statusCode
```