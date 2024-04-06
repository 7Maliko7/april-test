package structs

type Car struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type AddCarRequest struct {
	RegNum string `json:"regNum,omitempty"`
}

type AddCarResponse struct {
	Id int64 `json:"id,omitempty"`
}

type DeleteCarRequest struct {
	RegNum string `json:"regNum,omitempty"`
}

type DeleteCarResponse struct {
	Response
}

type Response struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

type UpdateCarRequest struct {
	Id      int64 `json:"id,omitempty"`
	Car     Car   `json:"car,omitempty"`
	OwnerID int   `json:"owner_id,omitempty"`
}

type UpdateCarResponse struct {
	Response
}

type GetCarRequest struct {
	RegNum string `json:"regNum,omitempty"`
}

type GetCarResponse struct {
	Car Car
}

type GetCarListRequest struct {
	Limit  int64  `json:"limit,omitempty"`
	Offset int64  `json:"offset,omitempty"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
}

type GetCarListResponse struct {
	CarList []Car
}
