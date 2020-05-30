package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func TestChili(t *testing.T) {

	rq, err := http.NewRequest("GET", "/chili", nil)
	if err != nil {
		t.Fatal(err)
	}

	rs := httptest.NewRecorder()
	handler := http.HandlerFunc(Chili)

	handler.ServeHTTP(rs, rq)
	if status := rs.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	beans := `{"answer": "beans"}`
	nobeans := `{"answer": "no beans"}`
	
	eq1, err := AreEqualJSON(rs.Body.String(), beans)
	eq2, err := AreEqualJSON(rs.Body.String(), nobeans)
	
	if !eq1 && !eq2 {
		t.Errorf("handler returned unexpected body: got %v want %v", rs.Body.String(), beans)
	}
	
}

func TestFu(t *testing.T) {

	rq, err := http.NewRequest("GET", "/fu", nil)
	if err != nil {
		t.Fatal(err)
	}

	rs := httptest.NewRecorder()
	handler := http.HandlerFunc(Fu)

	handler.ServeHTTP(rs, rq)
	if status := rs.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	s := `{"message": "FU AROCK!"}`
	
	eq, _ := AreEqualJSON(rs.Body.String(),s)
	if !eq {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rs.Body.String(), s)
	}

}

func TestCaliber(t *testing.T) {

	rq, err := http.NewRequest("GET", "/caliber/9mmOr45", nil)
	if err != nil {
		t.Fatal(err)
	}

	rs := httptest.NewRecorder()
	handler := http.HandlerFunc(Caliber)

	handler.ServeHTTP(rs, rq)
	if status := rs.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	s := `{"error": "divide by zero"}`
	eq, _ := AreEqualJSON(rs.Body.String(),s)
	if !eq {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rs.Body.String(), s)
	}

}

func TestInfo(t *testing.T) {
	rq, err := http.NewRequest("GET", "info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rs := httptest.NewRecorder()
	handler := http.HandlerFunc(Info)

	handler.ServeHTTP(rs, rq)
	if status := rs.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHealth(t *testing.T) {
	rq, err := http.NewRequest("GET", "health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rs := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rs, rq)
	if status := rs.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}