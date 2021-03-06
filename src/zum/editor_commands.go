package main

import (
	"log"
	"math"
	"strings"
)

func NavigateUp() {
	doc := CurrentDoc()
	if doc.Cursor.Y > 0 {
		doc.Cursor.Y -= 1

		if doc.Cursor.Y < doc.Scroll.Y {
			doc.Scroll.Y = doc.Cursor.Y
		}
	}
}

func NavigateDown() {
	doc := CurrentDoc()
	if doc.Cursor.Y < (math.MaxInt32 - 1) {
		doc.Cursor.Y += 1
	}
}

func NavigateLeft() {
	doc := CurrentDoc()
	if doc.Cursor.X > 0 {
		doc.Cursor.X -= 1

		if doc.Cursor.X < doc.Scroll.X {
			doc.Scroll.X = doc.Cursor.X
		}
	}
}

func NavigateRight() {
	doc := CurrentDoc()
	if doc.Cursor.X < (math.MaxInt32 - 1) {
		doc.Cursor.X += 1

		if doc.Cursor.X >= ViewportSize.X {
			doc.Scroll.X += 1
		}
	}
}

func exitApplication() {
	quit, exists := GetArg(1)
	quit = strings.ToLower(quit)

	if exists && (quit == "y" || quit == "yes") {
		applicationRunning = false
	} else {
		EnableInputMode("Quit (Y/n): ", "", func(line string) {
			result := strings.ToLower(line)
			if result == "" || result == "y" || result == "yes" {
				applicationRunning = false
			}
		})
	}
}

func saveDocument() {
	doc := CurrentDoc()
	filename, exists := GetArg(1)

	if exists {
		doc.Filename = filename
	}

	if doc.Filename == "" {
		EnableInputMode("Save file: ", "", func(line string) {
			doc.Filename = line
			if doc.Filename != "" {
				err := doc.Save()
				if err != nil {
					log.Println("Could not save document: ", err)
				}
			}
		})
	} else {
		err := doc.Save()
		if err != nil {
			log.Println("Could not save document: ", err)
		}
	}
}

func openDocumentCallback(filename string) {
	if filename == "" {
		return
	}

	doc, err := LoadDocument(filename)
	if err != nil {
		log.Println("Could not open document: ", err)
	} else {
		currentDocument = doc
	}
}

func openDocument() {
	filename, exists := GetArg(1)

	if exists {
		openDocumentCallback(filename)
	} else {
		EnableInputMode("Open file: ", "", openDocumentCallback)
	}
}

func modifyColumnWidth() {
	change, exists := GetIntArg(1)
	if !exists {
		return
	}

	CurrentDoc().ModifyColumnWidth(CurrentDoc().Cursor.X, change)
}
