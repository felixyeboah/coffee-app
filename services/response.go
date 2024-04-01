package services

type JsonResponse struct {
	Error  bool        `json:"error"`
	Messge string      `json:"message"`
	Data   interface{} `json:"data,omitresponse"`
}
