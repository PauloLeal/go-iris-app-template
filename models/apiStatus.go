package models

type ApiStatusData struct {
	Code    int
	Message string
}

type ApiStatus interface {
	Status() ApiStatusData
}

// ----

type ApiOk struct {
}

func (ok ApiOk) Status() ApiStatusData {
	return ApiStatusData{Code: 200, Message: "ok"}
}

// ----

type ApiNotFoundError struct {
}

func (nfe ApiNotFoundError) Status() ApiStatusData {
	return ApiStatusData{Code: 404, Message: "Not found"}
}

// ----

type ApiForbiddenError struct {
}

func (fe ApiForbiddenError) Status() ApiStatusData {
	return ApiStatusData{Code: 403, Message: "Forbidden"}
}

// ----

type ApiBadRequestError struct {
}

func (bre ApiBadRequestError) Status() ApiStatusData {
	return ApiStatusData{Code: 400, Message: "Bad Request"}
}

// ----

type ApiInternalServerError struct {
	Message string
}

func (ise ApiInternalServerError) Status() ApiStatusData {
	if len(ise.Message) == 0 {
		ise.Message = "Internal Server Status"
	}
	return ApiStatusData{Code: 500, Message: ise.Message}
}
