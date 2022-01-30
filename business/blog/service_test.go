package blog_test

import (
	"content/business"
	"content/business/blog"
	blogMock "content/business/blog/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func Test_FindContentById(t *testing.T) {
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
			name: "test 0 expected got content",
			args: parameters{1},
			want: blog.Blog{
				ID:      1,
				Title:   "Hello world",
				Content: "Hello world dang dang",
			},
			err: nil,
		},
		{
			name: "test 1 expected err database",
			args: parameters{1},
			err:  business.ErrDatabase,
		},
		{
			name: "test 2 expected err not found",
			args: parameters{1},
			want: blog.Blog{
				ID:      0,
				Title:   "",
				Content: "",
			},
			err: business.ErrNotFound,
		},
	}
	t.Run(tests[0].name, func(t *testing.T) {
		blogRepository.On("FindContentById", mock.AnythingOfType("int")).Return(&tests[0].want, nil).Once()
		got, err := blogService.FindContentById(tests[0].args.id)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, tests[0].args.id, got.ID)

	})

	t.Run(tests[1].name, func(t *testing.T) {
		blogRepository.On("FindContentById", mock.AnythingOfType("int")).Return(nil, business.ErrDatabase).Once()
		got, err := blogService.FindContentById(tests[1].args.id)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, err, business.ErrDatabase)

	})

	t.Run(tests[2].name, func(t *testing.T) {
		blogRepository.On("FindContentById", mock.AnythingOfType("int")).Return(&tests[2].want, nil).Once()
		got, err := blogService.FindContentById(tests[2].args.id)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, err, business.ErrNotFound)

	})

}

func Test_FindAllContent(t *testing.T) {

	type testCase struct {
		name string
		want []blog.Blog
		err  error
	}
	tests := []testCase{
		{
			name: "test 0 expected got all content",
			want: []blog.Blog{
				{
					ID:      1,
					Title:   "Hello World",
					Content: "Hello World",
				},
				{
					ID:      2,
					Title:   "Number 2",
					Content: "Number 2 general",
				},
			},
			err: nil,
		},
		{
			name: "test 1 expected err database",
			err:  business.ErrDatabase,
		},
		{
			name: "test 2 expected err not found",
			err:  business.ErrNotFound,
			want: []blog.Blog{},
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		blogRepository.On("FindAllContent").Return(tests[0].want, nil).Once()
		got, err := blogService.FindAllContent()

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, len(tests[0].want), len(got))
		assert.Equal(t, tests[0].want[0].Title, got[0].Title)
		assert.Equal(t, tests[0].want[1].Title, got[1].Title)

	})

	t.Run(tests[1].name, func(t *testing.T) {
		blogRepository.On("FindAllContent").Return(nil, tests[1].err).Once()
		got, err := blogService.FindAllContent()

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[1].err, err)

	})

	t.Run(tests[2].name, func(t *testing.T) {
		blogRepository.On("FindAllContent").Return(tests[2].want, nil).Once()
		got, err := blogService.FindAllContent()

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[2].err, err)

	})
}

func Test_InsertContent(t *testing.T) {
	type parameters struct {
		blog blog.Blog
	}
	type testCase struct {
		name     string
		arg      parameters
		want     blog.Blog
		wantFind int
		err      error
	}
	tests := []testCase{
		{
			name: "test 0 expected got all content",
			arg: parameters{
				blog.Blog{
					Title:   "Hello world cui",
					Content: "Hello world 3",
				},
			},
			wantFind: 3,
			want: blog.Blog{
				ID:          3,
				Title:       "Hello world cui",
				Content:     "Hello world 3",
				PublishedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			},
			err: nil,
		},
		{
			name: "test 1 expected err database",
			err:  business.ErrDatabase,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		blogRepository.On("InsertContent").Return(tests[0].wantFind, nil).Once()
		blogRepository.On("FindContentById", mock.AnythingOfType("int")).Return(&tests[0].want, nil).Once()

		got, err := blogService.InsertContent(tests[0].arg.blog)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, tests[0].want.ID, got.ID)
		assert.Equal(t, tests[0].want.Title, got.Title)
		assert.Equal(t, tests[0].want.Content, got.Content)

	})

	t.Run(tests[1].name, func(t *testing.T) {
		blogRepository.On("InsertContent").Return(nil, tests[1].err).Once()

		got, err := blogService.InsertContent(tests[0].arg.blog)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[1].err, err)

	})

}

func Test_UpdateContent(t *testing.T) {
	type parameters struct {
		blog blog.Blog
	}
	type testCase struct {
		name     string
		arg      parameters
		want     blog.Blog
		wantFind int
		err      error
	}
	tests := []testCase{
		{
			name: "test 0 expected got all content",
			arg: parameters{
				blog.Blog{
					ID:      3,
					Title:   "Hello world cui update",
					Content: "Hello world 3",
				},
			},
			want: blog.Blog{
				ID:        3,
				Title:     "Hello world cui update",
				Content:   "Hello world 3",
				UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			},
			err: nil,
		},
		{
			name: "test 1 expected err database",
			err:  business.ErrDatabase,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		blogRepository.On("UpdateContent").Return(nil).Once()
		blogRepository.On("FindContentById", mock.AnythingOfType("int")).Return(&tests[0].want, nil).Once()

		got, err := blogService.UpdateContent(tests[0].arg.blog)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, tests[0].want.ID, got.ID)
		assert.Equal(t, tests[0].want.Title, got.Title)
		assert.Equal(t, tests[0].want.Content, got.Content)

	})

	t.Run(tests[1].name, func(t *testing.T) {
		blogRepository.On("UpdateContent").Return(tests[1].err).Once()

		got, err := blogService.UpdateContent(tests[0].arg.blog)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[1].err, err)

	})

}

func Test_DeleteContent(t *testing.T) {
	type parameters struct {
		id int
	}
	type testCase struct {
		name string
		arg  parameters
		want blog.Blog
		err  error
	}
	tests := []testCase{
		{
			name: "test 0 expected got all content",
			arg:  parameters{3},
			want: blog.Blog{
				ID:        3,
				Title:     "Hello world cui update",
				Content:   "Hello world 3",
				UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			},
			err: nil,
		},
		{
			name: "test 1 expected err database",
			err:  business.ErrDatabase,
		},
		{
			name: "test 2 expected err not found",
			arg:  parameters{3},
			want: blog.Blog{
				ID:      0,
				Title:   "",
				Content: "",
			},
			err: business.ErrNotFound,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		blogRepository.On("DeleteContent").Return(&tests[0].want, nil).Once()

		got, err := blogService.DeleteContent(tests[0].arg.id)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, tests[0].want.ID, got.ID)
		assert.Equal(t, tests[0].want.Title, got.Title)
		assert.Equal(t, tests[0].want.Content, got.Content)

	})

	t.Run(tests[1].name, func(t *testing.T) {
		blogRepository.On("DeleteContent").Return(nil, tests[1].err).Once()

		got, err := blogService.DeleteContent(tests[0].arg.id)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[1].err, err)

	})

	t.Run(tests[2].name, func(t *testing.T) {
		blogRepository.On("DeleteContent").Return(&tests[2].want, nil).Once()
		got, err := blogService.DeleteContent(tests[2].arg.id)

		assert.NotNil(t, err)
		assert.Nil(t, got)
		assert.Equal(t, tests[2].err, err)

	})

}
