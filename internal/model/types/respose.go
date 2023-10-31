package model

type ResponseRequest struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type Result struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Pincode  string `json:"pincode"`
	SatScore int    `json:"sat_score"`
	Passed   string `json:"passed"`
}

type CreateSatResultResponse struct {
	*ResponseRequest
	Name    string `json:"name"`
}

type UpdateSatResultResponse struct {
	ResponseRequest
	Name    string `json:"name"`
	SatScore int    `json:"sat_score"`
}

type GetRankResponse struct {
	ResponseRequest
	Name    string `json:"name"`
	SatScore int    `json:"sat_score"`
	Rank    int    `json:"rank"`
}

type DeleteSatResultResponse struct {
	ResponseRequest
	Name    string `json:"name"`
}

type GetAllResultResponse struct {
	ResponseRequest
	Results []Result `json:"results"`
}