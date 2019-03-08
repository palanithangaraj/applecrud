# Apple CRUD API
## Description
The Apple CRUD API: Web service which combines two existing web services.
1. Fetch a random name from http://uinames.com/api/
2. Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?
firstName=John&lastName=Doe&limitTo=[nerdy]
3. Combine the results and return them to the use

# Example
## FETCHING A NAME
$ curl http://uinames.com/api/
{"name":"Δαμέας","surname":"Γιάνναρης","gender":"male","region":"Greece"}
FETCHING A JOKE

$ curl 'http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=\[nerdy\]'
{ "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.",
"categories": [“nerdy”] } }
USING THE NEW WEB SERVICE

$ curl ‘http://localhost:5000’
Δαμέας Γιάνναρης’s OSI network model has only one layer - Physical..

## Externalized Configuration
The service loads externalized configuration from the config.json file at startup. This file must be present in the service program's path when the program is launched.
The following JSON sample illustrates the expected structure of the config file.
```json
{
  "port": "5000",
  "preLog": "TNG - applecrud - ",
  "log-level": "trace"
}
```

## API Endpoints

>GET /health

Endpoint for monitoring systems to determine avaialability of the service.

>GET /

Combines the random name and Chuck Norris Joke with the latest data, returns cobined string output
```
Svc1:
{
    "name": "Mioara",
    "surname": "Naghi",
    "gender": "female",
    "region": "Romania"
}

svc2:
{
    "type": "success",
    "value": {
        "id": 495,
        "joke": "John Doe doesn't needs try-catch, exceptions are too afraid to raise.",
        "categories": [
            "nerdy"
        ]
    }
}

output:
	string

```
## Logging levels
There are 7 modes of logging available ranging from "trace" (most complete) to "panic" (most restricted).  Please set the level in the config file. If none is set or the value is invalid, the config will default to "panic".

```go
  case TraceLevel:
                return "trace"
        case DebugLevel:
                return "debug"
        case InfoLevel:
                return "info"
        case WarnLevel:
                return "warning"
        case ErrorLevel:
                return "error"
        case FatalLevel:
                return "fatal"
        case PanicLevel:
                return "panic"
```
