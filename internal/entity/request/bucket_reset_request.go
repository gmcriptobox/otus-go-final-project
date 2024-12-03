package request

type BucketResetRequest struct {
	Login string `json:"login"`
	IP    string `json:"ip"`
}
