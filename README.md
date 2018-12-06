# golang-jlogger

# Install 

```
go get github.com/jerryjobs/golang-jlogger

```

## if use dep 

```
dep ensure -add github.com/jerryjobs/golang-jlogger
```

# Example

```
// New instance 
var logger = jlogger.New("system")

// then
logger.Info("test log")
logger.InfoF("ok test printf [%s]", "Hello!")

```

# TODO

* Auto create log file dir
* Support log level
* Support config file