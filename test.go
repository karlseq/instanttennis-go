package main

type TestRequest struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func TestFunc(req TestRequest) (Response, error) {
	return Response{
		Message: "Hello" + req.Name,
	}, nil
}
