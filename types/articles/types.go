package articles

type Tag string

const (
	LIFESTYLE     Tag = "Lifestyle"
	PROGRAMMING   Tag = "Programming"
	TECHNOLOGY    Tag = "Technology"
	BUSINESS      Tag = "Business"
	ENTERTAINMENT Tag = "Entertainment"
	EDUCATION     Tag = "Education"
	ENVIRONMENT   Tag = "Environment"
	DESIGN        Tag = "Design"
	PERSONAL      Tag = "Personal"
	FINANCE       Tag = "Finance"
	NEWS_POLITICS Tag = "News & Politics"
	SPORTS        Tag = "Sports"
)

func ParseTags(dbArray []string) []Tag {
	tags := make([]Tag, len(dbArray))
	for i, v := range dbArray {
		tags[i] = Tag(v)
	}
	return tags
}

type Article struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	OwnerID     int64  `json:"owner_id"`
	OwnerName   string `json:"owner_name"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Tags        []Tag  `json:"tags"`
	CreatedAt   string `json:"created_at"`
}

type ArticlePublishPayload struct {
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
	OwnerID     int64  `json:"owner_id" validate:"required"`
	OwnerName   string `json:"owner_name" validate:"required"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Tags        []Tag  `json:"tags"`
}
