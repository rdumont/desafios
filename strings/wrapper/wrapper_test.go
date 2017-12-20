package wrapper

import "testing"

const (
	Justify     = true
	DontJustify = false
)

func TestWrap(t *testing.T) {
	testCase := func(text string, limit int, justify bool, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			result := Wrap(text, limit, justify)

			if result != expected {
				t.Errorf("\n--- expected:\n%v\n--- but got:\n%v\n", expected, result)
			}
		}
	}

	t.Run("empty text", testCase(
		"", 20, DontJustify,
		"\n",
	))

	t.Run("one line", testCase(
		"foo bar", 20, DontJustify,
		"foo bar\n",
	))

	t.Run("one justified line", testCase(
		"foo bar", 20, Justify,
		"foo              bar\n",
	))

	t.Run("wrap exceeding words", testCase(
		"this is a single line of text that should be wrapped", 20, DontJustify,
		`this is a single
line of text that
should be wrapped
`,
	))

	t.Run("retain simple line breaks", testCase(
		"i break\nthis line", 10, DontJustify,
		`i break
this line
`,
	))

	t.Run("retain consecutive line breaks", testCase(
		"i break\n\nthis line", 10, DontJustify,
		`i break

this line
`,
	))

	t.Run("word that exceeds limit", testCase(
		"whatever happens thisisonehelluvabigword for me to wrap", 20, DontJustify,
		`whatever happens
thisisonehelluvabigword
for me to wrap
`,
	))
}

func TestPrintLine(t *testing.T) {
	testCase := func(line []string, limit, length int, justify bool, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			result := printLine(line, limit, length, justify)

			if result != expected {
				t.Errorf("\n--- expected:\n%v\n--- but got:\n%v\n", expected, result)
			}
		}
	}

	t.Run("empty line", testCase(
		[]string{}, 20, 0, Justify,
		"\n",
	))

	t.Run("one word", testCase(
		[]string{"foo"}, 20, 3, Justify,
		"foo\n",
	))

	t.Run("many words evenly spaced", testCase(
		[]string{"foo", "bar", "haha"}, 20, 10, Justify,
		"foo     bar     haha\n",
	))

	t.Run("many words unevenly spaced", testCase(
		[]string{"foo", "bar", "hah"}, 20, 9, Justify,
		"foo      bar     hah\n",
	))

	t.Run("many words without justification", testCase(
		[]string{"foo", "bar", "hah"}, 20, 9, DontJustify,
		"foo bar hah\n",
	))
}
