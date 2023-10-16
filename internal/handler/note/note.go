package noteHandler

import (
	"github.com/dimasyotama/go-notes/database"
	"github.com/dimasyotama/go-notes/internal/model"
	// "github.com/dimasyotama/go-notes/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
    db := database.DB
    var notes []model.Note

    // find all notes in the database
    db.Find(&notes)

    // If no note is present return an error
    if len(notes) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    // Else return notes
    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
    err := c.BodyParser(note)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
    // Add a uuid to the note
    note.ID = uuid.New()
    // Create the Note and return error if encountered
    err = db.Create(&note).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
    }
	
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func GetNotebyId(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"status":"error", "message":"No Note Available", "data":nil})
	}

	return c.JSON(fiber.Map{"status":"success", "message":"Notes Found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title		string `json:"title"`
		SubTitle	string	`json:"sub_title"`
		Text		string	`json:"text"`
	}
	
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message": "Data not found", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

	note.Title = updateNoteData.SubTitle
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}


func DeletNote(c *fiber.Ctx) error{
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
    }

    // Return success message
    return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}