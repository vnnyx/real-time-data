package domain

import (
	"encoding/json"
	"io"

	"github.com/vnnyx/real-time-data/internal/utils"
)

// type ResultType string

// const (
// 	ArticlesResultType               ResultType = "articles"
// 	URIWgtListResultType             ResultType = "uriWgtList"
// 	LangAggrResultType               ResultType = "langAggr"
// 	TimeAggrResultType               ResultType = "timeAggr"
// 	SourceAggrResultType             ResultType = "sourceAggr"
// 	SourceExAggrResultType           ResultType = "sourceExAggr"
// 	AuthorAggrResultType             ResultType = "authorAggr"
// 	KeywordAggrResultType            ResultType = "keywordAggr"
// 	LocAggrResultType                ResultType = "locAggr"
// 	ConceptAggrResultType            ResultType = "conceptAggr"
// 	ConceptGraphResultType           ResultType = "conceptGraph"
// 	CategoryAggrResultType           ResultType = "categoryAggr"
// 	DateMentionAggrResultType        ResultType = "dateMentionAggr"
// 	SentimentAggrResultType          ResultType = "sentimentAggr"
// 	RecentActivityArticlesResultType ResultType = "recentActivityArticles"
// )

// type DataType string

// const (
// 	Article DataType = "article"
// 	PR      DataType = "pr"
// 	Blog    DataType = "blog"
// )

// type DataTimeWindows uint64

// const (
// 	OneWeek  DataTimeWindows = 7
// 	OneMonth DataTimeWindows = 31
// )

// type SortBy string

// const (
// 	Relevance              SortBy = "rel"
// 	Date                   SortBy = "date"
// 	SourceImportance       SortBy = "sourceImportance"
// 	SourceImportanceRank   SortBy = "sourceImportanceRank"
// 	SourceAlexaGlobalRank  SortBy = "sourceAlexaGlobalRank"
// 	SourceAlexaCountryRank SortBy = "sourceAlexaCountryRank"
// 	SocialStore            SortBy = "socialScore"
// 	FacebookShares         SortBy = "facebookShares"
// )

// type NewsAPIRequest struct {
// 	Action                 string          `json:"action"`
// 	Keyword                string          `json:"keyword"`
// 	ArticlePage            uint64          `json:"articlePage"`
// 	ArticleSortBy          SortBy          `json:"articleSortBy"`
// 	ArticlesSortByAsc      bool            `json:"articlesSortByAsc"`
// 	ArticlesArticleBodyLen int64           `json:"articlesArticleBodyLen"`
// 	ResultType             string          `json:"resultType"`
// 	DataTypes              []DataType      `json:"dataTypes"`
// 	APIKey                 string          `json:"apiKey`
// 	ForceMaxDataTimeWindow DataTimeWindows `json:"forceMaxDataTimeWindow"`
// }

// type NewsAPIResponse struct {
// 	Articles     Articles `json:"articles"`
// 	TotalResults uint64   `json:"totalResults"`
// 	Page         uint64   `json:"page"`
// 	Count        uint64   `json:"count"`
// 	Pages        uint64   `json:"pages"`
// }

// type Articles struct {
// 	Results []struct {
// 		URI         string `json:"uri"`
// 		Lang        string `json:"lang"`
// 		IsDuplicate bool   `json:"isDuplicate"`
// 		Date        string `json:"date"`
// 		Time        string `json:"time"`
// 		DateTime    string `json:"dateTime"`
// 		DateTimePub string `json:"dateTimePub"`
// 		DataType    string `json:"dataType"`
// 		SIM         string `json:"sim"`
// 		URL         string `json:"url"`
// 		Title       string `json:"title"`
// 		Body        string `json:"body"`
// 		Source      struct {
// 			URI      string `json:"uri"`
// 			DataType string `json:"dataType"`
// 			Title    string `json:"title"`
// 		}
// 		Authors []struct {
// 			URI      string `json:"uri"`
// 			Name     string `json:"name"`
// 			Type     string `json:"type"`
// 			IsAgency bool   `json:"isAgency"`
// 		}
// 		Links  []string `json:"links"`
// 		Videos []struct {
// 			URI   string `json:"uri"`
// 			Label string `json:"label"`
// 		}
// 	}
// }

type NewsAPIResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Article
}

type Article struct {
	Source      Source  `json:"source"`
	Author      *string `json:"author"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	URL         string  `json:"url"`
	URLToImage  *string `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
	Content     *string `json:"content"`
}

type Source struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

func (n *NewsAPIResponse) FromJSON(data io.Reader) error {
	return json.NewDecoder(data).Decode(n)
}

func (a *Article) ToNewsEvent() *NewsEvent {
	return &NewsEvent{
		Source:      a.Source,
		Author:      utils.NonNilValue(a.Author, ""),
		Title:       a.Title,
		Description: utils.NonNilValue(a.Description, ""),
		URL:         a.URL,
		URLToImage:  utils.NonNilValue(a.URLToImage, ""),
		PublishedAt: a.PublishedAt,
		Content:     utils.NonNilValue(a.Content, ""),
	}
}
