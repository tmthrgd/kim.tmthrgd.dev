package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	mrand "math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
	"go.tmthrgd.dev/refresh"
)

var postTmpl = newTemplate("post.tmpl")

func fetchPosts(blog *tumblr.BlogRef) (*tumblr.Blog, []tumblr.PostInterface, error) {
	var posts []tumblr.PostInterface

	params := url.Values{
		"limit":  {"20"},
		"offset": {"0"},

		"filter": {"text"},
		"type":   {"photo"},
	}
	for {
		postsResp, err := blog.GetPosts(params)
		if err != nil {
			return nil, nil, &tumblrAPIError{err}
		}

		postsAll, err := postsResp.All()
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, postsAll...)

		if len(postsResp.Posts) <= 20 {
			break
		}

		params["offset"][0] = strconv.Itoa(len(posts))
	}

	blogInfo, err := blog.GetInfo()
	if err != nil {
		return nil, nil, err
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].GetSelf().Id <
			posts[j].GetSelf().Id
	})

	return blogInfo, posts, nil
}

func postHandler() http.HandlerFunc {
	blog := tumblrclient.NewClient(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET")).
		GetBlog("kimjongunlookingatthings.tumblr.com")

	type result struct {
		blogInfo *tumblr.Blog
		posts    []tumblr.PostInterface
	}
	posts := refresh.New(time.Hour, func() (interface{}, error) {
		blogInfo, posts, err := fetchPosts(blog)
		posts = filterPosts(posts)
		return result{blogInfo, posts}, err
	})
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		postID := chi.URLParam(r, "postID")
		postID64, err := strconv.ParseUint(postID, 10, 64)
		if postID != "" && err != nil {
			return os.ErrNotExist
		}

		v, err := posts.Load()
		blogInfo, posts := v.(result).blogInfo, v.(result).posts
		if err != nil {
			return err
		} else if len(posts) == 0 {
			return os.ErrNotExist
		}

		var post tumblr.PostInterface
		if postID == "" {
			post = posts[mrand.Intn(len(posts))]
		} else {
			idx := sort.Search(len(posts), func(i int) bool {
				return posts[i].GetSelf().Id >= postID64
			})
			if idx >= len(posts) || posts[idx].GetSelf().Id != postID64 {
				return os.ErrNotExist
			}
			post = posts[idx]
		}

		imagePost, ok := post.(*tumblr.PhotoPost)
		if !ok {
			return errors.New("internal error: post is not a *tumblr.PhotoPost")
		} else if len(imagePost.Photos) == 0 {
			return errors.New("upstream error: photo post has no photos")
		}

		caption, ok := captionReplacement[imagePost.Id]
		if !ok {
			caption = imagePost.Caption
		}
		if caption == "" {
			caption = "looking at things"
		}

		nonce := make([]byte, 24)
		if _, err := rand.Read(nonce); err != nil {
			return err
		}
		b64Nonce := base64.StdEncoding.EncodeToString(nonce)

		w.Header().Set("Content-Security-Policy", strings.ReplaceAll(contentSecurityPolicy,
			"style-src", "style-src 'nonce-"+b64Nonce+"' 'unsafe-inline'"))

		return templateExecute(w, postTmpl, &struct {
			CSPNonce string
			IsIndex  bool

			PostID  uint64
			Caption string
			Image   string

			SourceURL   string
			SourceTitle string
		}{
			b64Nonce,
			postID == "",

			imagePost.Id,
			caption,
			imagePost.Photos[0].OriginalSize.Url,

			blogInfo.Url,
			blogInfo.Title,
		})
	})
}
