package extractor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllLinks(t *testing.T) {
	// Test input data
	srcURL := "https://example.com/test"
	body := `
		<html>
			<body>
				<a href="https://example.com/internal1">Internal Link 1</a>
				<a href="https://example.com/internal2">Internal Link 2</a>
				<a href="/internal3">Internal Link 3</a>
				<a href="https://external.com/external1">External Link 1</a>
			</body>
		</html>
	`

	expectedPageLinks := &PageLinks{
		URL:     "https://example.com/test",
		BaseURL: "https://example.com",
		InternalLinks: []string{
			"https://example.com/internal1",
			"https://example.com/internal2",
			"https://example.com/internal3",
		},
		ExternalLinks: []string{
			"https://external.com/external1",
		},
	}

	// Test function call
	actualPageLinks, err := GetAllLinks(srcURL, body)

	// Test assertions
	require.NoError(t, err)
	require.EqualValues(t, expectedPageLinks, actualPageLinks)
}

func TestGetBaseURLWithInvalidURL(t *testing.T) {
	// Test input data
	invalidURL := "invalid-url%"

	// Test function call
	actualBaseURL, err := getBaseURL(invalidURL)

	// Test assertions
	require.Empty(t, actualBaseURL)
	require.Error(t, err)
}

func TestGetBaseURLWithValidURL(t *testing.T) {
	// Test input data
	validURL := "https://example.com/path1/path2"

	// Test function call
	actualBaseURL, err := getBaseURL(validURL)

	// Test assertions
	require.NoError(t, err)
	require.Equal(t, "https://example.com", actualBaseURL)
}

func TestSplitLinks(t *testing.T) {
	baseURL := "https://example.com"
	links := []string{
		"https://example.com/about",
		"https://example.com/contact",
		"https://example.com/products",
		"https://example.com/articles",
		"https://google.com",
		"https://example.org",
	}

	internal, external := splitLinks(baseURL, links)

	require.ElementsMatch(t, []string{
		"https://example.com/about",
		"https://example.com/contact",
		"https://example.com/products",
		"https://example.com/articles",
	}, internal)

	require.ElementsMatch(t, []string{
		"https://google.com",
		"https://example.org",
	}, external)
}

func TestUnique(t *testing.T) {
	strs := []string{"a", "b", "b", "b", "c", "a", "d"}

	uniqueStrs := unique(strs)

	require.ElementsMatch(t, []string{"a", "b", "c", "d"}, uniqueStrs)
}
