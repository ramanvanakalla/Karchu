package exceptions

import "net/http"

type GeneralException struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func BadRequestError(message string, status string) *GeneralException {
	Ex := GeneralException{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Status:     status,
	}
	return &Ex
}

func InternalServerError(message string, status string) *GeneralException {
	Ex := GeneralException{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Status:     status,
	}
	return &Ex
}

func UnAuthorizedError(message string, status string) *GeneralException {
	Ex := GeneralException{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Status:     status,
	}
	return &Ex
}
