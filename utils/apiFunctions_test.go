package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRespond(t *testing.T) {
	expected_request_status := 200
	w := httptest.NewRecorder()
	var result_map interface{}
	response_map := map[string]interface{}{
		"testing": true,
	}

	jData, _ := json.Marshal(response_map)
	Respond(w, jData)
	result_body, _ := ioutil.ReadAll(w.Result().Body)
	_ = json.Unmarshal(result_body, &result_map)

	if w.Result().StatusCode != expected_request_status {
		t.Error(fmt.Sprintf("Error, received status `%v` is different than %v", w.Result().StatusCode, expected_request_status))
	}

	if !reflect.DeepEqual(result_map, response_map) {
		t.Error(fmt.Sprintf("Error, received response `%v` is different from expected response `%v`", result_map, response_map))
	}
}
