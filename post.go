package main

import (
	dbstorage "github.com/nektro/go.dbstorage"
)

type Post struct {
	ID        int64 `json:"id"`
	IDS       string
	Subreddit string `json:"subreddit" sqlite:"text"`
	PostID    string `json:"post_id" sqlite:"text"`
	Title     string `json:"title" sqlite:"text"`
	PostJson  string `json:"json" sqlite:"text"`
	Link      string `json:"link" sqlite:"text"`
}

func InsertPost(sub, post, title, pjson, link string) {
	id := db.QueryNextID("posts")
	db.QueryPrepared(true, "insert into posts values (?, ?, ?, ?, ?, ?)", id, sub, post, title, pjson, link)
}

func DoesPostExist(post string) bool {
	return dbstorage.QueryHasRows(db.Build().Se("*").Fr("posts").Wh("post_id", post).Exe())
}
