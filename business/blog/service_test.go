package blog_test

import (
	"content/business/blog"
	blogMock "content/business/blog/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	blogService    blog.Service
	blogRepository blogMock.Repository
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	blogService = blog.NewService(&blogRepository)
}

func Test_longestSubstring(t *testing.T) {
	type parameters struct {
		id int
	}
	type testCase struct {
		name string
		args parameters
		want blog.Blog
		err  error
	}
	tests := []testCase{
		{
			name: "test 0 get content",
			args: parameters{1},
			want: blog.Blog{
				ID:      1,
				Title:   "Hello world",
				Content: "Hello world dang dang",
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := blogService.FindContentById(tt.args.id)

			assert.NotNil(t, got)
			assert.Nil(t, err)

		})
	}
}
