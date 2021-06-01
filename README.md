# Mila_go
Mila_go is a web development framework written in Golang which implements the server-side Model View Controller (MVC) pattern. Many of its components and concepts will seem familiar to those of us with experience in other web frameworks like Ruby on Rails or Phoenix (Elixir).

## Features
### Router with Fiber 
`https://gofiber.io/`
* Controller for every model

### View custom with template Raymond 
`https://github.com/aymerick/raymond`
* View with master layout
* Template with helper

### ORM with GORM 
`https://gorm.io/`
* Model with cast and validate
### Live reload with AIR
`https://github.com/cosmtrek/air`
* See the complete [air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml)

## Install
* clone `https://github.com/LangPham/mila_go`
* `cd mila_go`
* `go get -u -v -f all`

## Usage
* Run air with command `air` in dir `mila_go`