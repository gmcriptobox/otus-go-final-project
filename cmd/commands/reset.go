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

var resetCmd = &cobra.Command{
	Use:   "reset [login] [ip]",
	Short: "Reset buckets",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverAddress, err := cmd.Flags().GetString("ip")
		if err != nil {
			fmt.Println(err)
			return
		}
		resetBuckets(args[0], args[1], serverAddress)
	},
}

func init() {
	bucketCmd.AddCommand(resetCmd)
}

func resetBuckets(login, ip, serverAddress string) {
	resetRequest := &request.BucketResetRequest{Login: login, IP: ip}
	requestBody, _ := easyjson.Marshal(resetRequest)

	httpRequest, err := http.NewRequestWithContext(context.Background(), http.MethodDelete,
		serverAddress+"/api/bucket", bytes.NewBuffer(requestBody))
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
		fmt.Println("Buckets successfully reset")
	case 400:
		fmt.Println("Invalid ip or login")
	case 500:
		fmt.Println("Internal server error")
	default:
		fmt.Println("Unknown error")
	}
}
