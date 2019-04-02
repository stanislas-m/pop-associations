package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/stanislas-m/pop-associations/onetomany/models"
)

func createEntities(c *pop.Connection) error {
	for i := 0; i < 3; i++ {
		// Create variable amount of fruits
		fruits := make([]models.Fruit, i)
		for j := 0; j < i; j++ {
			fruits[j] = models.Fruit{}
		}
		b := &models.Tree{
			Name:   fmt.Sprintf("Tree %d", i),
			Fruits: fruits,
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

	trees := &models.Trees{}
	if err := c.All(trees); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("basic fetch trees: %v", trees)

	trees = &models.Trees{}
	if err := c.Eager().All(trees); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("eager fetch trees: %v", trees)

	fruits := &models.Fruits{}
	if err := c.All(fruits); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("basic fetch fruits: %v", fruits)

	fruits = &models.Fruits{}
	if err := c.Eager().All(fruits); err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Printf("eager fetch fruits: %v", fruits)
}
