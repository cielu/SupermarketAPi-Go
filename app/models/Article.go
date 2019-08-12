package models

import "time"

// ArticleModel 文章
type Article struct {
	ID           uint       `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
	Name         string     `json:"name"`
	BrowseCount  uint       `json:"browseCount"`
	CommentCount uint       `json:"commentCount"`
	CollectCount uint       `json:"collectCount"`
	Status       int        `json:"status"`
	Content      string     `json:"content"`
	HTMLContent  string     `json:"htmlContent"`
	ContentType  int        `json:"contentType"`
	// Categories    []Category `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
	// Comments      []Comment  `gorm:"ForeignKey:SourceID" json:"comments"`
	UserID uint `json:"userID"`
	// User          User       `json:"user"`
	LastUserID uint `json:"lastUserID"` //最后一个回复话题的人
	// LastUser      User       `json:"lastUser"`
	LastCommentAt *time.Time `json:"lastCommentAt"`
}
