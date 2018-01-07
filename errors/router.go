package main

import (
    "net/http"
    "log"
)

type HandlerFunc  func(*http.Response)

type ResponseRouter struct {
    Handlers map[int]HandlerFunc
    DefaultHandler  HandlerFunc
}

func NewRouter() *ResponseRouter {
    return &ResponseRouter {
        Handlers: make(map[int]HandlerFunc),
        DefaultHandler: func(r *http.Responce) {
                 log.Fatalln("Unhandled Response:", r.StatusCode)           
        },
    }
}

func (r *ResponseRouter) Register(status int, handler HandlerFunc) {
    r.Handler[status] = handler
}

func (r *ResponseRouter) 