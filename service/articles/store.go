package articles

import (
	"database/sql"
	"fmt"
	"github.com/Darthex/ink-golang/types"
	"github.com/Darthex/ink-golang/types/articles"
	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewArticleStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type ArticleOut map[string]interface{}

func (s *Store) GetArticles(p types.Pagination) (ArticleOut, error) {
	orderClause := "ORDER BY " + p.Field + " " + string(p.Sort)
	query := `
    SELECT id, title, content, owner_id, owner_name, description, cover, tags, created_at
    FROM articles 
    WHERE title ILIKE $1 OR content ILIKE $1 OR owner_name ILIKE $1
    ` + orderClause + `
    LIMIT $2 OFFSET $3`

	searchTerm := "%" + p.Search + "%"

	rows, err := s.db.Query(query, searchTerm, p.Take, p.Skip)
	if err != nil {
		return nil, err
	}
	all := new([]articles.Article)
	for rows.Next() {
		a := new(articles.Article)
		a, err := scanRowIntoArticle(rows)
		if err != nil {
			return nil, err
		}
		*all = append(*all, *a)
	}
	if len(*all) == 0 {
		all = &[]articles.Article{}
	}

	var totalCount int
	countQuery := `
    SELECT COUNT(*) 
    FROM articles 
    WHERE title ILIKE $1 OR content ILIKE $1 OR owner_name ILIKE $1`

	err = s.db.QueryRow(countQuery, searchTerm).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"result": all,
		"total":  totalCount,
	}, nil
}

func (s *Store) GetArticleById(id int64) (*articles.Article, error) {
	rows, err := s.db.Query("SELECT id, title, content, owner_id, owner_name, description, cover, tags, created_at FROM articles WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	a := new(articles.Article)
	for rows.Next() {
		a, err = scanRowIntoArticle(rows)
		if err != nil {
			return nil, err
		}
	}
	if a.ID == 0 {
		return nil, fmt.Errorf("article not found")
	}
	return a, nil
}

func (s *Store) CreateNewArticle(a articles.ArticlePublishPayload) error {
	if _, err := s.db.Exec(
		"INSERT INTO articles (title, content, owner_id, owner_name, description, cover, tags) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		a.Title,
		a.Content,
		a.OwnerID,
		a.OwnerName,
		a.Description,
		a.Cover,
		pq.Array(a.Tags),
	); err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateArticle(a articles.ArticlePublishPayload, id int64) error {
	if _, err := s.db.Exec(
		"UPDATE articles SET title=$2, content=$3, owner_id=$4, owner_name=$5, description=$6, cover=$7, tags=$8 where id=$1",
		id,
		a.Title,
		a.Content,
		a.OwnerID,
		a.OwnerName,
		a.Description,
		a.Cover,
		pq.Array(a.Tags),
	); err != nil {
		return err
	}
	return nil
}

func scanRowIntoArticle(row *sql.Rows) (*articles.Article, error) {
	article := new(articles.Article)
	var dbTags []string
	if err := row.Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.OwnerID,
		&article.OwnerName,
		&article.Description,
		&article.Cover,
		pq.Array(&dbTags),
		&article.CreatedAt,
	); err != nil {
		return nil, err
	}
	article.Tags = articles.ParseTags(dbTags)
	return article, nil
}
