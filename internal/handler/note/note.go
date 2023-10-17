package noteHandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dimasyotama/go-notes/database"
	"github.com/dimasyotama/go-notes/internal/model"
	"github.com/dimasyotama/go-notes/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
    db := database.DB
    var notes []model.Note
	redis_connect := redis.ConnectRedis()
	key := "go-notes:getall"

    // find all notes in the database
    db.Find(&notes)

	// If no note is present return an error
    if len(notes) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

	cached_notes, _ := redis_connect.Get(context.Background(), key).Result()

	if cached_notes != "" {
        if err := json.Unmarshal([]byte(cached_notes), &notes); err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error unmarshalling cached notes", "data": nil})
        }
    } else {
        // Notes were not found in the cache, fetch from the database
        db.Find(&notes)

        // Store notes in the Redis cache for future requests
        notes_json, err := json.Marshal(notes)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error marshalling notes", "data": nil})
        }
        err = redis_connect.Set(context.Background(), key, notes_json, 0).Err()
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error caching notes", "data": nil})
        }
    }

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

    // Return the created note
    return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func GetNotebyId(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	redis_connect := redis.ConnectRedis() //connect into redis

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	key := fmt.Sprintf("%s:%s", "go-notes:get-by-id", id) 


	if note.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"status":"error", "message":"No Note Available", "data":nil})
	}

	cached_notes, _ := redis_connect.Get(context.Background(), key).Result()

	//if note are found in cache, unmarshall the JSON Data
	if cached_notes != "" {
        if err := json.Unmarshal([]byte(cached_notes), &note); err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error unmarshalling cached notes", "data": nil})
        }
    } else {
        // Note were not found in the cache, fetch from the database
        db.Find(&note)

        // Store note in the Redis cache for future requests
        notes_json, err := json.Marshal(note)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error marshalling notes", "data": nil})
        }
        err = redis_connect.Set(context.Background(), key, notes_json, 0).Err()
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error caching notes", "data": nil})
        }
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


func DeleteNote(c *fiber.Ctx) error{
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