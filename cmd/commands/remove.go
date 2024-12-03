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

var removeCmd = &cobra.Command{
	Use:   "remove [ip] [mask]",
	Short: "Remove network from list",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverAddress, err := cmd.Flags().GetString("ip")
		if err != nil {
			fmt.Println(err)
			return
		}
		if cmd.Parent().Name() == whiteList {
			removeFromList(args[0], args[1], whiteList, serverAddress)
			return
		} else if cmd.Parent().Name() == blackList {
			removeFromList(args[0], args[1], blackList, serverAddress)
			return
		}
		fmt.Println("Unknown command")
	},
}

func init() {
	removeCmdForWhiteList := *removeCmd
	removeCmdForBlackList := *removeCmd
	whiteListCmd.AddCommand(&removeCmdForWhiteList)
	blackListCmd.AddCommand(&removeCmdForBlackList)
}

func removeFromList(ip, mask, list, serverAddress string) {
	networkRequest := &request.NetworkRequest{Network: ip + "/" + mask}
	requestBody, err := easyjson.Marshal(networkRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequest, err := http.NewRequestWithContext(context.Background(), http.MethodDelete,
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
	case 200:
		fmt.Println("Successfully removed from " + list)
	case 400:
		fmt.Println("Invalid ip or mask")
	case 500:
		fmt.Println("Internal server error")
	default:
		fmt.Println("Unknown error")
	}
}
