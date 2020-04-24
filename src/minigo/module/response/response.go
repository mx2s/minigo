package response

import (
	"encoding/json"
	"fmt"
	error2 "minigo/module/error"
	"net/http"
)

func Item(w http.ResponseWriter, code int, key string, data string) string {
	res := fmt.Sprintf(`{"data":{"%v": %v}}`, key, data)
	w.WriteHeader(code)
	fmt.Fprintln(w, res)
	return res
}

func Items(w http.ResponseWriter, code int, key string, data string) string {
	return Item(w, code, key, fmt.Sprintf("[%s]", data))
}

func Data(w http.ResponseWriter, code int, jsonStr string) string {
	res := fmt.Sprintf(`{"data": %v }`, jsonStr)
	w.WriteHeader(code)
	fmt.Fprintln(w, res)
	return res
}

func AppErr(w http.ResponseWriter, err *error2.AppError) string {
	res := fmt.Sprintf(`{"errors":[{"code": %v,"message":"%s"}]}`, err.Code, err.Message)
	w.WriteHeader(err.Code)
	fmt.Fprintln(w, res)
	return res
}

func Error(w http.ResponseWriter, code int, message string) string {
	res := fmt.Sprintf(`{"errors":[{"code": %v,"message":"%s"}]}`, code, message)
	w.WriteHeader(code)
	fmt.Fprintln(w, res)
	return res
}

func ManyErrors(w http.ResponseWriter, code int, errors []string) string {
	var errorsArray []error2.AppError
	for _, errorMessage := range errors {
		errorsArray = append(errorsArray, error2.AppError{
			Code:    code,
			Message: errorMessage,
		})
	}
	resB, _ := json.Marshal(errorsArray)
	res := string(resB)
	res = fmt.Sprintf(`{"errors":%s}`, res)
	w.WriteHeader(code)
	fmt.Fprintln(w, res)
	return res
}
