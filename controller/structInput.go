package controller

type createPostInput struct {
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
}

type createUserInput struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Privillege int64  `json:"privillege"`
}

type createCommentInput struct {
	Content string `json:"content" binding:"required"`
}

type createLabelInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
