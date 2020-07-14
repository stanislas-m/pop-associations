package main

import (
	"log"

	"github.com/gobuffalo/pop/v5"
	"github.com/stanislas-m/pop-associations/manytomany/models"
)

func createEntities(c *pop.Connection) error {
	skills := []models.Skill{
		{
			Label: "Math",
		},
		{
			Label: "English",
		},
		{
			Label: "Biology",
		},
	}
	if err := c.Create(skills); err != nil {
		return err
	}

	// Create students with their skills
	bob := &models.Student{
		Name: "Bob",
		Skills: []models.Skill{
			skills[1],
			skills[2],
		},
	}
	if err := c.Eager().Create(bob); err != nil {
		return err
	}

	alice := &models.Student{
		Name: "Alice",
		Skills: []models.Skill{
			skills[0],
			skills[2],
		},
	}
	if err := c.Eager().Create(alice); err != nil {
		return err
	}

	return nil
}

func main() {
	pop.Debug = true
	c, err := pop.Connect("development")
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	defer func() {
		log.Println("Cleanup database")
		c.TruncateAll()
	}()

	if err := createEntities(c); err != nil {
		log.Printf("err: %v", err)
		return
	}

	students := &models.Students{}
	if err := c.All(students); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("basic fetch: %v", students)

	students = &models.Students{}
	if err := c.Eager().All(students); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("eager fetch: %v", students)
}