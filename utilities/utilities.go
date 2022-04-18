package utilities

// globally accessible utilities

import (
	"log"
)


// Check Error and log
func CE(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
