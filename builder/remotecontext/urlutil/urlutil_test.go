package urlutil // import "github.com/Prakhar-Agarwal-byte/moby/builder/remotecontext/urlutil"

import "testing"

var (
	gitUrls = []string{
		"git://github.com/Prakhar-Agarwal-byte/moby",
		"git@github.com:docker/docker.git",
		"git@bitbucket.org:atlassianlabs/atlassian-docker.git",
		"https://github.com/Prakhar-Agarwal-byte/moby.git",
		"http://github.com/Prakhar-Agarwal-byte/moby.git",
		"http://github.com/Prakhar-Agarwal-byte/moby.git#branch",
		"http://github.com/Prakhar-Agarwal-byte/moby.git#:dir",
	}
	incompleteGitUrls = []string{
		"github.com/Prakhar-Agarwal-byte/moby",
	}
	invalidGitUrls = []string{
		"http://github.com/Prakhar-Agarwal-byte/moby.git:#branch",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}
