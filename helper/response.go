package helper

import "api-service-go/dto"


type ResponseWithData struct {
	Status bool `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
	Paginate *dto.PaginateDto`json:"paginate"`
	Data any `json:"data"`
}

type ResponseWithoutData struct {
	Status bool `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
}


func Response(params dto.ResponseParams) any{
	var response any
	var status bool


	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = true
	} else {
		status = false
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Status : status,
			Code: params.StatusCode,
			Message: params.Message,
			Paginate: params.Paginate,
			Data: params.Data,
		}
	}else{
		response = &ResponseWithoutData{
			Status: status,
			Code: params.StatusCode,
			Message: params.Message,
		}
	}
	return response
}