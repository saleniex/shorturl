package handlers

type ShortIdUri struct {
	ShortId string `uri:"id" binding:"required"`
}
