# Raspberry Pi lights controller

Controls lights attached to GPIOs via http in two modes:
- as the server which listen to JSON commands
- as comet client which polls comet server for new commands

Structure of the JSON is the same in both cases - the map (string/string) where key contains lamp number and value contains pattern to perform.

## Patterns

Patterns are described by up to three parameters

	<turn_on_time>[/<period>] [<total_time>]

- period with leading / is optional - if not specified is equal to turn_on_time
- total_time is optional - if not specified is equal 0 which means forever

```
 ________________             _______
|                |___________|       ...____
<- turn_on_time ->
<---------- period ---------->
<--------------- total_time --------------->

```

Time unit for patterns is 100ms.

### Examples

- turn on lamp 0
```
{"0":"1"}
```

- turn off lamp 0
```
{"0":"0"}
```

- flash lamp 2 for 2s
```
{"2":"1 20"}
```

- blink lamp 1 (500ms on / 500 ms off) for 7s
```
{"1":"5/10 70"}
```

- blink lamp 2 - device 0 (200ms on / 400 ms off) for 6s. Flash lamp 1 - device 3 for 10s
```
{
  "0-2":"2/6 60",
  "3-1":"1 100"
}
```
 
## How to build
- download golang package for your architecture and setup GOROOT (and add to GOPATH $GOROOT/bin) with it
- clone this repo to your golang project directory structure (ie. to: ~/goworkspace/src/github.com/jkiet/go-pie)
- you can invoke init script to setup environment variables (see comments inside)
```
source ./init
```
- download dependencies
```
go get
```
- build
```
go build
```
- run as root
```
sudo ./go-pie
```

## Examples

- config file
```
section: 0
layout: [25, 24, 23, 18, 22, 27, 17, 4]
```
- run as restful service
```
sudo ./go-pie listen 0.0.0.0:8888 config.yml
```
- send command
```
curl -i -XPOST -H'Content-Type: application/json' -d'{"0":"1 10", "1":"1 20", "2":"1 30", "3":"1 40", "4": "1 50", "5": "1 60", "6":"1 70", "7":"1 80"}' http://127.0.0.1:8888/lamps/reload/
```
expected response looks like this:
```
HTTP/1.1 200 OK                                                                                                                                                                                                  
Content-Type: application/json
Date: Thu, 28 Jan 2016 23:33:06 GMT
Content-Length: 167

{
  "data": {
   "0": "OK",
   "1": "OK",
   "2": "OK",
   "3": "OK",
   "4": "OK",
   "5": "OK",
   "6": "OK",
   "7": "OK"
  },
  "_meta": {
   "status": "ok"
  }
```

- another command
```
curl -i -XPOST -H'Content-Type: application/json' -d'{"0":"1/9", "1":"2/9", "2":"3/9", "3":"4/9", "4": "4/9", "5": "3/9", "6":"2/9", "7":"1/9"}' http://127.0.0.1:8888/lamps/reload/
```
