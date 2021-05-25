package snippet

import (
	"testing"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/datatype"
	"github.com/stretchr/testify/assert"
)

func TestRequest_List(t *testing.T) {
	cases := []struct {
		name    string
		request list
		fail    bool
	}{
		{
			"success",
			list{Favorite: false,
				AccessLevel: -1,
				Title:       "", SortBy: "id",
				OrderBy: "desc",
				Page:    1,
				Limit:   50,
			},
			false,
		},
		{
			"invalid_access_level",
			list{
				Favorite:    false,
				AccessLevel: 4,
				Title:       "",
				SortBy:      "id",
				OrderBy:     "desc",
				Page:        1,
				Limit:       50},
			true,
		},
		{
			"invalid_title",
			list{
				Favorite:    false,
				AccessLevel: -1,
				Title:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
				SortBy:      "id",
				OrderBy:     "desc",
				Page:        1,
				Limit:       50,
			},
			true,
		},
		{
			"invalid_sort_by",
			list{
				Favorite:    false,
				AccessLevel: 4,
				Title:       "",
				SortBy:      "title",
				OrderBy:     "desc",
				Page:        1,
				Limit:       50,
			},
			true,
		},
		{
			"invalid_order_by",
			list{
				Favorite:    false,
				AccessLevel: 4,
				Title:       "",
				SortBy:      "title",
				OrderBy:     "ascc",
				Page:        1,
				Limit:       50,
			},
			true,
		},
		{
			"invalid_limit",
			list{
				Favorite:    false,
				AccessLevel: 4,
				Title:       "",
				SortBy:      "title",
				OrderBy:     "ascc",
				Page:        1,
				Limit:       500,
			},
			true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}
func TestRequest_Upsert(t *testing.T) {
	cases := []struct {
		name    string
		request snippet
		fail    bool
	}{
		{
			"success",
			snippet{
				Favorite:            datatype.FlexibleBool{Value: true, String: "1"},
				AccessLevel:         0,
				Title:               "test",
				Content:             "",
				Language:            "",
				CustomEditorOptions: entity.CustomEditorOptions{},
			},
			false,
		},
		{
			"invalid_access_level",
			snippet{
				Favorite:            datatype.FlexibleBool{Value: true, String: "1"},
				AccessLevel:         432,
				Title:               "test",
				Content:             "hello world",
				Language:            "go",
				CustomEditorOptions: entity.CustomEditorOptions{},
			},
			true,
		},
		{
			"empty_title",
			snippet{
				Favorite:            datatype.FlexibleBool{Value: true, String: "1"},
				AccessLevel:         0,
				Title:               "",
				Content:             "hello world",
				Language:            "go",
				CustomEditorOptions: entity.CustomEditorOptions{},
			},
			true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}
