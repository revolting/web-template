package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Directory",
        "GET",
        "/directory",
        Directory,
    },
    Route{
        "Authenticate",
        "GET",
        "/authenticate",
        Authenticate,
    },
    Route{
        "Authenticate",
        "POST",
        "/authenticate",
        Authenticate,
    },
    Route{
        "Validate",
        "GET",
        "/validate",
        Validate,
    },
    Route{
        "Validate",
        "POST",
        "/validate",
        Validate,
    },
}
