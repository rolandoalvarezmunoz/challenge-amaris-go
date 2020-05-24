package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_goDotEnvVariable(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goDotEnvVariable(tt.args.key); got != tt.want {
				t.Errorf("goDotEnvVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_getData(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getData(tt.args.w, tt.args.r)
		})
	}
}

func Test_getData2(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error de estado: devolvio %v se esperaba %v",
			status, http.StatusOK)
	}
}
