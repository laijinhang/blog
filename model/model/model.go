package model

type ArchiverModel struct {
	Cid          int32 // 用户id
	Title        string
	Slug         string
	Text         string
	Order        int32
	AuthorId     int32
	Template     string
	Type         string
	Statue       string
	Password     string
	CommentsNum  int32
	AllowComment string
	AllowPing    string
	AllowFeed    string
	Parent       int32
	Views        int32
	Model
}

func (*ArchiverModel) TableName() string {
	return "typecho_contents"
}

type Model struct {
	Created  int64
	Modified int64
}
