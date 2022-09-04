package search

type Query struct {
	Q          string `schema:"q,omitempty,required"`
	Targets    string `schema:"targets,omitempty,required"`
	Fields     string `schema:"fields,omitempty"`
	Filters    string `schema:"filters,omitempty"`
	JsonFilter string `schema:"jsonFilter,omitempty"`
	Sort       string `schema:"_sort,omitempty,required"`
	Offset     int    `schema:"_offset,omitempty"`
	Limit      int    `schema:"_limit,omitempty"`
	Context    string `schema:"_context,omitempty,required"`
}

type Meta struct {
	Status     int    `json:"status"`
	ID         string `json:"id"`
	TotalCount int    `json:"totalCount"`
}

type Data struct {
	ContentID       string `json:"contentId"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	UserID          int    `json:"userId"`
	ChannelID       int    `json:"channelId"`
	ViewCounter     int    `json:"viewCounter"`
	MyListCounter   int    `json:"mylistCounter"`
	LikeCounter     int    `json:"likeCounter"`
	LengthSeconds   int    `json:"lengthSeconds"`
	ThumbnailURL    string `json:"thumbnailUrl"`
	StartTime       string `json:"startTime"`
	LastResBody     string `json:"lastResBody"`
	CommentCounter  int    `json:"commentCounter"`
	LastCommentTime string `json:"lastCommentTime"`
	CategoryTags    string `json:"categoryTags"`
	Tags            string `json:"tags"`
	TagsExact       string `json:"tagsExact"`
	Genre           string `json:"genre"`
	GenreKeyword    string `json:"genre.keyword"`
}

type Response struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}
