package blog_test

import (
	"content/business"
	"content/business/blog"
	blogMock "content/business/blog/mocks"
	"os"
	"testing"

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
			err: business.ErrDatabase,
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
