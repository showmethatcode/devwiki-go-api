package controllers

import (
	"devwiki/db"
	"devwiki/ent/term"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateTerm(c *gin.Context) error {
	var body RequestBody

	d, err := db.Conn()
	if err != nil {
		log.Println(err)
	}

	tx, err := d.Tx(c)
	if err != nil {
		log.Println(err)
	}

	if err = c.BindJSON(&body); err != nil {
		log.Println(err)
	}

	// 중복되는 용어가 있는지 체크하고 Boolean을 리턴함
	hasTerm, err := tx.Term.
		Query().
		Where(term.NameEQ(body.Name)).
		Exist(c)
	if err != nil {
		log.Println(err)
		return tx.Rollback()
	}

	if hasTerm == true {
		existedTerm, err := tx.Term.
			Query().
			Where(term.NameEQ(body.Name)).
			Only(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		newRevision, err := tx.TermRevision.
			Create().
			SetDescription(body.Description).
			SetTermID(existedTerm.ID).
			Save(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		_, err = tx.TermPointer.
			Update().
			SetTermID(existedTerm.ID).
			SetRevisionID(newRevision.ID).
			Save(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		c.JSON(http.StatusCreated, gin.H{
			"term":     existedTerm.Name,
			"revision": newRevision,
		})

		return tx.Commit()
	}

	if hasTerm == false {
		newRevision, err := tx.TermRevision.
			Create().
			SetDescription(body.Description).
			Save(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		newTerm, err := tx.Term.
			Create().
			SetName(body.Name).
			AddRevisions(newRevision).
			Save(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		_, err = tx.TermPointer.
			Create().
			SetTermID(newTerm.ID).
			SetRevisionID(newRevision.ID).
			Save(c)
		if err != nil {
			log.Println(err)
			return tx.Rollback()
		}

		c.JSON(http.StatusOK, gin.H{
			"term":     newTerm,
			"revision": newRevision,
		})
	}

	return tx.Commit()
}
