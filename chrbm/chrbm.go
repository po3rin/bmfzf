package chrbm

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

// Bookmark contains bookmark data that have url type
type Bookmark struct {
	Name string
	Path string
	URL  string
}

// BookmarkTree is organaized by nodes that has bookmark data.
type BookmarkTree struct {
	Roots Roots `json:"roots"`
}

// Roots has kind of Chrome Bbookmarks.
type Roots struct {
	BookmarkBar Node `json:"bookmark_bar"`
	Other       Node `json:"other"`
	Synced      Node `json:"synced"`
}

// Node has Bookmark info.
// folder type Node has children Node.
type Node struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Children []Node `json:"children"`
}

// Visitor visits Nodes.
type Visitor interface {
	Visit(n Node, path string) error
}

type bookmarkRecoder struct {
	data []Bookmark
}

func newBookmarkRecoder() *bookmarkRecoder {
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

// WalkEdge walk edge nodes that has url type.
func WalkEdge(n Node, visitor Visitor) error {
	return walkEdge(n, "/", visitor)
}

func walkEdge(n Node, path string, v Visitor) error {
	switch n.Type {
	case "folder":
		for _, c := range n.Children {
			p := filepath.Join(path, c.Name)
			err := walkEdge(c, p, v)
			if err != nil {
				return err
			}
		}
	case "url":
		v.Visit(n, path)
	case "":
		// empty type
		return nil
	default:
		return fmt.Errorf("unsupported type: %+v", n.Type)
	}
	return nil
}

// ListBookmarks return Chrome Bookmark List that has no hierarchy.
// byteJSON needs Chrome Bookmark JSON format.
func ListBookmarks(byteJSON []byte) ([]Bookmark, error) {
	if !json.Valid(byteJSON) {
		return nil, errors.New("input is invalid format (required json only)")
	}

	var m BookmarkTree
	err := json.Unmarshal(byteJSON, &m)
	if err != nil {
		return nil, err
	}

	r := newBookmarkRecoder()

	// TODO: consider order
	var eg errgroup.Group
	eg.Go(func() error {
		return WalkEdge(m.Roots.Other, r)
	})
	eg.Go(func() error {
		return WalkEdge(m.Roots.Synced, r)
	})
	eg.Go(func() error {
		return WalkEdge(m.Roots.BookmarkBar, r)
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return r.data, nil
}
