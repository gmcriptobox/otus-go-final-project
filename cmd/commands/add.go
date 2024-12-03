package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/KarimovKamil/otus-go-final-project/internal/entity/request"
	"github.com/mailru/easyjson"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ip] [mask]",
	Short: "Add network to list",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverAddress, err := cmd.Flags().GetString("ip")
		if err != nil {
			fmt.Println(err)
			return
		}
		if cmd.Parent().Name() == whiteList {
			addToList(args[0], args[1], whiteList, serverAddress)
			return
		} else if cmd.Parent().Name() == blackList {
			addToList(args[0], args[1], blackList, serverAddress)
			return
		}
		fmt.Println("Unknown command")
	},
}

func init() {
	addCmdForWhiteList := *addCmd
	addCmdForBlackList := *addCmd
	whiteListCmd.AddCommand(&addCmdForWhiteList)
	blackListCmd.AddCommand(&addCmdForBlackList)
}

func addToList(ip, mask, list, serverAddress string) {
	networkRequest := &request.NetworkRequest{Network: ip + "/" + mask}
	requestBody, err := easyjson.Marshal(networkRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequest, err := http.NewRequestWithContext(context.Background(), http.MethodPost,
		serverAddress+"/api/"+list, bytes.NewBuffer(requestBody))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(httpRequest.Body)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.DefaultClient.Do(httpRequest)
	if response == nil {
		fmt.Println("No response, check server address")
		return
	}
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	switch response.StatusCode {
	case 201:
		fmt.Println("Successfully added to " + list)
	case 400:
		fmt.Println("Invalid ip or mask")
	case 409:
		fmt.Println("Network already exists in " + list)
	case 500:
		fmt.Println("Internal server error")
	default:
		fmt.Println("Unknown error")
	}
}
