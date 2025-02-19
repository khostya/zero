package pgmodel

//go:generate reform

// NewsCategory represents a row in news_category table.
//
//reform:news_category
type NewsCategory struct {
	NewsID     int32 `reform:"news_id"`
	CategoryID int32 `reform:"category_id"`
}
