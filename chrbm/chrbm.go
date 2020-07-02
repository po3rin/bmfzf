package chrbm

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
)

type Bookmark struct {
	Name string
	Path string
	URL  string
}

type Bookmarks struct {
	Roots Roots `json:"roots"`
}

type Roots struct {
	BookmarkBar Node `json:"bookmark_bar"`
}

type Node struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Children []Node `json:"children"`
}

type Visitor interface {
	Visit(n Node, path string) error
}

type bookmarkRecoder struct {
	data []Bookmark
}

func NewBookmarkRecoder() *bookmarkRecoder {
	return &bookmarkRecoder{
		data: make([]Bookmark, 0, 100),
	}
}

func (b *bookmarkRecoder) Visit(n Node, path string) error {
	bm := Bookmark{
		Name: n.Name,
		Path: path,
		URL:  n.URL,
	}
	b.data = append(b.data, bm)
	return nil
}

func WalkEdge(n Node, visitor Visitor) error {
	return walkEdge(n, "/", visitor)
}

func walkEdge(n Node, path string, v Visitor) error {
	switch n.Type {
	case "folder":
		for _, c := range n.Children {
			// Now support bookmark_bar only ...
			p := filepath.Join(path, c.Name)
			if c.Name == "ブックマークバー" {
				p = "/"
			}
			err := walkEdge(c, p, v)
			if err != nil {
				return err
			}
		}
	case "url":
		v.Visit(n, path)
	default:
		return fmt.Errorf("unsupported type: %+v", n.Type)
	}
	return nil
}

func NewBookmark(byteJSON []byte) ([]Bookmark, error) {
	if !json.Valid(byteJSON) {
		return nil, errors.New("input is invalid format (required json only)")
	}

	var m Bookmarks
	err := json.Unmarshal(byteJSON, &m)
	if err != nil {
		return nil, err
	}

	r := NewBookmarkRecoder()
	err = WalkEdge(m.Roots.BookmarkBar, r)
	if err != nil {
		return nil, err
	}

	return r.data, nil
}
