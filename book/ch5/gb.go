/*
/*
@Time : 2020/11/23 4:49 下午
@Author : chengqunzhong
@File : gb
@Software: GoLand
*/
package main

import (
	"time"
)

const IssuesURL  = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items  []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt  time.Time `json:"created_at"`
	Body       string
}

type User struct {
	Login  string
	HTMLURL string `json:"html_url"`
}