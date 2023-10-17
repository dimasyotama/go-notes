package noteRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/dimasyotama/go-notes/internal/handler/note"
)
func SetupNoteRoutes(router fiber.Router){
	note := router.Group("/note")

	//Create a note
	note.Post("/note-post", noteHandler.CreateNotes)

	//Read all notes
	note.Get("/get-all-note", noteHandler.GetNotes)

	//Read one note
	note.Get("/:noteId", noteHandler.GetNotebyId)

	//Update one Note
	note.Put("/:noteId", noteHandler.UpdateNote)
    
	// Delete one Note
    note.Delete("/:noteId", noteHandler.DeleteNote)
}