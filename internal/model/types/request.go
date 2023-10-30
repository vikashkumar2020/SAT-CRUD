package model

type CreateSatResultRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Pincode  string `json:"pincode"`
	SatScore int    `json:"sat_score"`
}

type UpdateSatResultRequest struct {
	Name     string `json:"name"`
	SatScore int    `json:"sat_score"`
}

type GetRankRequest struct {
	Name string `json:"name"`
}