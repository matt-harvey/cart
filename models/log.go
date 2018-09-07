package models

import (
	"log"
	"os"
)

var Log *log.Logger = log.New(os.Stdout, "cart ", log.Lshortfile|log.LstdFlags)
