# Gateway with Javascript
This recipe is a gateway which runs some javascript.

## Installation
* Install [Go](https://golang.org/)

## Setup
```
git clone https://github.com/project-flogo/jsexec
cd jsexec/examples/microgateway/api
```

## Testing

Start the gateway:
```
go run main.go
```

Run the following command:
```
curl http://localhost:9096/calculate"
```

You should see the following like response:
```json
{"sum":3}
```
