# Logging

## Introduction

Logging is a log wrapper on top of go log package which could be used to have
leveled logs. Currently it has 5 levels:

1. Trace
2. Debug
3. Info
4. Warning
5. Error


## Usage

To use this package :

1. Set enviornment variable `LOG_LEVEL` to either of above mentioned levels.
2. Import this package in you go code :

```
import "github.com/ebay/fabio/logging"
```

3. Use logger in your code

```
logging.Trace("Trace logs")
logging.Debug("Debug logs")
logging.Info("Info logs")
logging.Warn("Warning logs %s","ZZZZZZ")
logging.Error("Error logs %s","KABOOM")
```
