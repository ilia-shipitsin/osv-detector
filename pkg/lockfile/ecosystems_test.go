package lockfile_test

import (
	"github.com/g-rath/osv-detector/pkg/lockfile"
	"os"
	"strings"
	"testing"
)

func numberOfLockfileParsers(t *testing.T) int {
	t.Helper()

	directories, err := os.ReadDir(".")

	if err != nil {
		t.Fatalf("unable to read current directory: ")
	}

	count := 0

	for _, directory := range directories {
		if strings.HasPrefix(directory.Name(), "parse-") &&
			!strings.HasSuffix(directory.Name(), "_test.go") {
			count++
		}
	}

	return count
}

func TestKnownEcosystems(t *testing.T) {
	t.Parallel()

	expectedCount := numberOfLockfileParsers(t)

	// - npm, yarn, and pnpm,
	// - pip, poetry, and pipenv,
	// - maven and gradle,
	// all use the same ecosystem so "ignore" those parsers in the count
	expectedCount -= 5

	ecosystems := lockfile.KnownEcosystems()

	if knownCount := len(ecosystems); knownCount != expectedCount {
		t.Errorf("Expected to know about %d ecosystems, but knew about %d", expectedCount, knownCount)
	}

	uniq := make(map[lockfile.Ecosystem]int)

	for _, ecosystem := range ecosystems {
		uniq[ecosystem]++

		if uniq[ecosystem] > 1 {
			t.Errorf(`Ecosystem "%s" was listed more than once`, ecosystem)
		}
	}
}
