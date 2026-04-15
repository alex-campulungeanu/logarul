# Logarul

A Go logging library wrapper that combines console and file logging with automatic log rotation.

## Features

- **Console logging** with colored output (JSON format)
- **File logging** with automatic rotation
- **Tee handler** for logging to multiple outputs simultaneously

## Installation

```bash
go get alex-campulungeanu.github.com/logarul
```

## Usage

```go
package main

import (
	"log/slog"
	"alex-campulungeanu.github.com/logarul"
)

func main() {
	cfg := logarul.NewMinimalConfig()
	cfg.FilePath = "app.log"
	cfg.Level = slog.LevelDebug

	logger, _ := logarul.New(cfg)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warning message")
	logger.Error("error message")
}
```

## Configuration

| Field | Type | Description |
|------|------|-------------|
| `FilePath` | `string` | Path to log file |
| `Level` | `slog.Level` | Minimum log level |
| `MaxSize` | `int` | Max size in MB before rotation |
| `MaxBackups` | `int` | Number of backup files to keep |
| `MaxAge` | `int` | Days to retain old files |
| `Compress` | `bool` | Compress rotated files |
| `AddSource` | `bool` | Include source code location |
| `EnableFile` | `bool` | Enable file logging |
| `EnableConsole` | `bool` | Enable console logging |

## Example with Full Config

```go
cfg := logarul.Config{
    FilePath:      "app.log",
    Level:         slog.LevelDebug,
    MaxSize:       10,
    MaxBackups:    3,
    MaxAge:        7,
    Compress:      true,
    AddSource:    true,
    EnableFile:   true,
    EnableConsole: true,
}
```