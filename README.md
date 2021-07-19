# Greg

Greg is a very simple tool who enables user to send data through CURL, and display these data as a Prometheus format. Array is reinialized each time Prometheus scrapes the endpoint.

This is not supposed to be useful to a lot of people, because we are bypassing the concept of Prometheus : we are doing passive monitoring. Use it with caution !

## Endpoints :

* `GET` `/metrics` : Display all current metrics in a Prometheus format, and flush the content of the array
* `POST` `/metrics` : Add data to the array

Payload :
```json
{
	"name":"donuts_available{location=\"kitchen\"}",
	"value": 8
}
```

Value is a float : you can pass integer, or a floating object.


## Build & run

I suppose yuou already configured your GOROOT and GOPATH.

From root folder :

```sh
go get -t github.com/gorilla/mux
go build -o greg ./src
```

Note : use GOOS and GOARCH for other compilations.

Server will bind on port 9666.


## Todolist :
* Basic authentication on POST endpoint
* Dockerfile


## Licence

Licenced under MIT licence. 

Original work (c) 2021 Artheriom.


## Enjoy !

