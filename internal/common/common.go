package common

type ApiResult struct {
	Message      string     `json:"Message"`
	Status       string     `json:"Status"`
	ReturnedRows int64      `json:"ReturnedRows"`
	DeletedRows  int64      `json:"DeletedRows"`
	UpdatedRows  int64      `json:"UpdatedRows"`
	InsertedRows int64      `json:"InsertedRows"`
	Data         []struct{} `json:"Data"`
}
