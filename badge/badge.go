package badge

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// TotalCoverage runs 'go tool cover' on the provided profile path and returns
// the coverage percentage as a string and a float64.
func TotalCoverage(profilePath string) (string, float64, error) {
	cmd := exec.Command("go", "tool", "cover", "-func="+profilePath)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", 0, fmt.Errorf("go tool cover failed: %v\n%s", err, string(b))
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	if len(lines) == 0 {
		return "", 0, fmt.Errorf("empty output from go tool cover")
	}

	last := lines[len(lines)-1]
	fields := strings.Fields(last)
	if len(fields) == 0 {
		return "", 0, fmt.Errorf("unexpected final line: %q", last)
	}

	raw := fields[len(fields)-1]
	raw = strings.TrimSuffix(raw, "%")

	value, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return "", 0, fmt.Errorf("invalid percentage: %q", raw)
	}

	return raw, value, nil
}

// BadgeColor returns a color name based on the coverage percentage.
func BadgeColor(p float64) string {
	switch {
	case p >= 90:
		return "brightgreen"
	case p >= 80:
		return "green"
	case p >= 70:
		return "yellowgreen"
	case p >= 60:
		return "yellow"
	case p >= 50:
		return "orange"
	default:
		return "red"
	}
}

// Download fetches a URL and saves the response body to the specified filepath.
func Download(url, filepath string) error {
	resp, err := http.Get(url)
	fmt.Println("Downloading", url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

// Escape replaces special characters in a string to be used in a shields.io URL.
func Escape(s string) string {
	return strings.NewReplacer("-", "--", "_", "__", " ", "_").Replace(s)
}

// Fail prints an error message to stderr and exits the program with status 1.
func Fail(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

// ExecuteBadge parses command line flags and generates the coverage badge.
func ExecuteBadge() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: gonest-tools badge <pattern1> [pattern2] ...")
		os.Exit(0)
	}

	in := flag.String("in", "coverage.out", "input coverage profile path")
	out := flag.String("out", ".public/coverage.svg", "output badge SVG path")
	label := flag.String("label", "coverage", "label for the badge")
	flag.Parse()

	percentText, percentValue, err := TotalCoverage(*in)
	if err != nil {
		Fail(err)
	}

	color := BadgeColor(percentValue)

	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		Fail(err)
	}

	url := fmt.Sprintf("https://img.shields.io/badge/%s-%s%%25-%s.svg", Escape(*label), Escape(percentText), Escape(color))
	if err := Download(url, *out); err != nil {
		Fail(err)
	}
}
