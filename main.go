package main

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/stanislas-m/pop-associations/models"
)

func createEntities(c *pop.Connection) error {
	for i := 0; i < 3; i++ {
		b := &models.Body{
			Head: models.Head{},
		}

		if err := c.Eager().Create(b); err != nil {
			return err
		}
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

	bodies := &models.Bodies{}
	if err := c.All(bodies); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("basic fetch: %v", bodies)

	bodies = &models.Bodies{}
	if err := c.Eager().All(bodies); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("eager fetch: %v", bodies)
}
