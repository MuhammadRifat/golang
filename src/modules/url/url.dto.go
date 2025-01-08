package url

type URLDto struct {
	OriginalUrl string `binding:"required"`
	UserId      int
}
