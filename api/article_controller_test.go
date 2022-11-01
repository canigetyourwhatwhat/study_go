package api

import (
	"net/http"
	"net/http/httptest"
	"path"
	"practice_go/interfaces"
	"reflect"
	"testing"
)

const DOMAIN = "http://localhost:8080"

func TestArticleController_GetArticle(t *testing.T) {

	urlPath := path.Join(DOMAIN, "article", "all")
	req := httptest.NewRequest(http.MethodGet, urlPath, nil)
	res := httptest.NewRecorder()

	type fields struct {
		service interfaces.ArticleService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"1", fields{ser}, args{res, req}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ArticleController{
				service: tt.fields.service,
			}
			c.GetArticle(tt.args.w, tt.args.r)
		})
	}
}

func TestArticleController_ListArticles(t *testing.T) {
	type fields struct {
		service interfaces.ArticleService
	}
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ArticleController{
				service: tt.fields.service,
			}
			c.ListArticles(tt.args.w, tt.args.in1)
		})
	}
}

func TestArticleController_PostArticle(t *testing.T) {
	type fields struct {
		service interfaces.ArticleService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ArticleController{
				service: tt.fields.service,
			}
			c.PostArticle(tt.args.w, tt.args.r)
		})
	}
}

func TestArticleController_PostNice(t *testing.T) {
	type fields struct {
		service interfaces.ArticleService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ArticleController{
				service: tt.fields.service,
			}
			c.PostNice(tt.args.w, tt.args.r)
		})
	}
}

func TestNewMyAppController(t *testing.T) {
	type args struct {
		s interfaces.ArticleService
	}
	tests := []struct {
		name string
		args args
		want *ArticleController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleController(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleController() = %v, want %v", got, tt.want)
			}
		})
	}
}
