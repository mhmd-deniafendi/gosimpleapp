package main

import (
	"net/http"
	"testing"
)

func Test_renderTemplate(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		tmpl string
		page *Data
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderTemplate(tt.args.w, tt.args.tmpl, tt.args.page)
		})
	}
}

func TestIndexHandler(t *testing.T) {
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
			IndexHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestAboutHandler(t *testing.T) {
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
			AboutHandler(tt.args.w, tt.args.r)
		})
	}
}
