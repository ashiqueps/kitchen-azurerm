package main

import (
	"C"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"log"
)

var (
	subscriptionID = "80b824de-ec53-4116-9868-3deeab10b0cd" //os.Getenv("AZURE_SUBSCRIPTION_ID")
)

type Param struct {
	Location          string `json:"location"`
	ResourceGroupName string `json:"resourceGroupName"`
}

//export create_resource_group
func create_resource_group(params *C.char) C.int {
	new_params := C.GoString(params)
	args := Param{}
	json.Unmarshal([]byte(new_params), &args)
	log.Println(fmt.Sprintf("Creating a new resource group %s on location: %s", args.ResourceGroupName, args.Location))

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	log.Println("After auth")
	ctx := context.Background()

	resourceGroupClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		log.Println(err)
		return 0
	}

	resourceGroupResp, err := resourceGroupClient.CreateOrUpdate(
		ctx,
		args.ResourceGroupName,
		armresources.ResourceGroup{
			Location: to.Ptr(args.Location),
		},
		nil)
	if err != nil {
		log.Println(err)
		return 0
	}
	log.Println(&resourceGroupResp.ResourceGroup)
	return 1
}

func main() {
	//create_resource_group(`{"location":"fasdfsdf", "resourceGroupName": "ashique-fasdfasfsafs"}`)
}
