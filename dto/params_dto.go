package dto

type ResponseParams struct {
	StatusCode int 
	Message string
	Paginate *PaginateDto
	Data any
}