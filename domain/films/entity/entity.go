package entity

type Films struct {
	Id             int64  `db:"id" json:"id"`
	Title          string `db:"title" json:"title"`
	Description    string `db:"description" json:"description"`
	ImageThumbnail string `db:"image_thumbnail" json:"image_thumbnail"`
}
