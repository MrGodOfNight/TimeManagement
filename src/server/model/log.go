/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

package model

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Color constants for log messages
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

// Logger is a struct for logger.
type Logger struct {
	fileLogger    *log.Logger
	consoleLogger *log.Logger
}

// NewLogger creates a new logger.
func NewLogger() (*Logger, error) {
	// Get current time
	currentTime := time.Now()

	// Format current time
	formattedTime := currentTime.Format("02-01-2006")

	// Create file name
	fileName := "logs/" + formattedTime + ".log"

	// Create directory if needed
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return nil, fmt.Errorf("error creating directory for logs: %v", err)
		}
	}

	// Open file
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	// Create loggers
	fileLogger := log.New(file, "", log.Ldate|log.Ltime)
	consoleLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Logger{fileLogger: fileLogger, consoleLogger: consoleLogger}, nil
}

// Info writes info messages.
func (cl *Logger) Info(file bool, args ...any) {
	cl.consoleLogger.Println(ColorGreen+"INFO: ", args, ColorReset)
	if file {
		cl.fileLogger.Println("INFO: ", args)
		return
	}
}

// Warning writes warning messages.
func (cl *Logger) Warning(file bool, args ...any) {
	cl.consoleLogger.Println(ColorYellow+"WARNING: ", args, ColorReset)
	if file {
		cl.fileLogger.Println("WARNING: ", args)
		return
	}
}

// Error writes error messages.
func (cl *Logger) Error(file bool, args ...any) {
	cl.consoleLogger.Println(ColorRed+"ERROR: ", args, ColorReset)
	if file {
		cl.fileLogger.Println("ERROR: ", args)
		return
	}
}

// Debug writes debug messages.
func (cl *Logger) Debug(file bool, args ...any) {
	cl.consoleLogger.Println(ColorBlue+"DEBUG: ", args, ColorReset)
	if file {
		cl.fileLogger.Println("DEBUG: ", args)
		return
	}
}
