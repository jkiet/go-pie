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
{"0":"1"}
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
- you can invoke init script to setup environment variables (see coments inside)
	
	source ./init

- download dependencies

	go get

- build

	go build

- run as root

	sudo ./go-pie

