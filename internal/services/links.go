package services

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/tallquist10/linkslasher/internal/links"
)

type LinksService struct {
	db                                          *sql.DB
	generator                                   *links.Generator
	createLink, updateLink, getLink, deleteLink *sql.Stmt
}

type CustomHash struct {
	Hash string
}

// type CustomExpiration struct {
// 	Expiration int64
// }

// func NewLinksService(file string) (*LinksService, error) {
func NewLinksService(
	db *sql.DB,
	createQuery string,
	readQuery string,
	// updateQuery string,
	deleteQuery string,
) (*LinksService, error) {
	createLink, err := db.Prepare(createQuery)
	if err != nil {
		return nil, err
	}
	// updateLink, err := db.Prepare(updateQuery)
	// if err != nil {
	// 	return nil, err
	// }
	getLink, err := db.Prepare(readQuery)
	if err != nil {
		return nil, err
	}
	deleteLink, err := db.Prepare(deleteQuery)
	if err != nil {
		return nil, err
	}
	return &LinksService{
		db:         db,
		generator:  links.NewGenerator(),
		createLink: createLink,
		// updateLink: updateLink,
		getLink:    getLink,
		deleteLink: deleteLink,
	}, nil
}

func (ls *LinksService) GetLink(hash string) (*links.Link, error) {
	var link *links.Link
	err := ls.getLink.QueryRow(hash).Scan(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (ls *LinksService) CreateLink(link *links.Link) (*links.Link, error) {
	if link.Original == "" {
		return nil, errors.New("please enter a valid link")
	}

	var hash string
	if link.Hash == "" {
		for {
			generatedHash, err := ls.generator.GeneratePath(link.Original)
			if err != nil {
				return nil, err
			}
			_, err = ls.GetLink(generatedHash)
			if err == sql.ErrNoRows {
				hash = generatedHash
				break
			}
			if err != nil {
				return nil, err
			}
			if hash != "" {
				break
			}
		}
		link.Hash = hash
	}

	result, err := ls.createLink.Exec(link.Original, hash)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		panic("No errors found, but no rows inserteds")
	}
	return link, nil
}

// func (ls *LinksService) UpdateLink(hash string) (*links.Link, error) {
// 	if time.Now().Unix() < expiration {
// 		expiration = time.Now().Unix()
// 	}
// 	result, err := ls.updateLink.Exec(hash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if rows == 0 {
// 		return nil, errors.New(fmt.Sprintf("no link found to be updated for hash %s", hash))
// 	}
// 	return ls.GetLink(hash)
// }

// func (ls *LinksService) DeleteLink(hash string) error {
// 	result, err := ls.deleteLink.Exec(hash)
// }

func (ls *LinksService) Close() (int, error) {
	err := ls.createLink.Close()
	if err != nil {
		return 1, err
	}
	err = ls.getLink.Close()
	if err != nil {
		return 1, err
	}
	err = ls.updateLink.Close()
	if err != nil {
		return 1, err
	}
	err = ls.deleteLink.Close()
	if err != nil {
		return 1, err
	}
	err = ls.db.Close()
	if err != nil {
		return 1, err
	}
	return 0, nil
}
