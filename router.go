package main

import (
	"net/http"
	_"fmt"
)

type Router struct{
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHanler(path string) (http.HandlerFunc, bool){
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	handler, exist := r.FindHanler(request.URL.Path)

	if !exist{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w,request)
}