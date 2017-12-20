package wrapper

import "testing"

func TestWrap(t *testing.T) {
	testCase := func(text string, limit int, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			result := Wrap(text, limit)

			if result != expected {
				t.Errorf("\n--- expected:\n%v\n--- but got:\n%v\n", expected, result)
			}
		}
	}

	t.Run("empty text", testCase(
		"", 20,
		"\n",
	))

	t.Run("one line", testCase(
		"foo bar", 20,
		"foo              bar\n",
	))

	t.Run("wrap exceeding words", testCase(
		"this is a single line of text that should be wrapped", 20,
		`this   is  a  single
line  of  text  that
should   be  wrapped
`,
	))

	t.Run("retain simple line breaks", testCase(
		"i break\nthis line", 10,
		`i    break
this  line
`,
	))

	t.Run("retain consecutive line breaks", testCase(
		"i break\n\nthis line", 10,
		`i    break

this  line
`,
	))

	t.Run("word that exceeds limit", testCase(
		"whatever happens thisisonehelluvabigword for me to wrap", 20,
		`whatever     happens
thisisonehelluvabigword
for   me   to   wrap
`,
	))
}

func TestPrintLine(t *testing.T) {
	testCase := func(line []string, limit, length int, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			result := printLine(line, limit, length)

			if result != expected {
				t.Errorf("\n--- expected:\n%v\n--- but got:\n%v\n", expected, result)
			}
		}
	}

	t.Run("empty line", testCase(
		[]string{}, 20, 0,
		"\n",
	))

	t.Run("one word", testCase(
		[]string{"foo"}, 20, 3,
		"foo\n",
	))

	t.Run("many words evenly spaced", testCase(
		[]string{"foo", "bar", "haha"}, 20, 10,
		"foo     bar     haha\n",
	))

	t.Run("many words unevenly spaced", testCase(
		[]string{"foo", "bar", "hah"}, 20, 9,
		"foo      bar     hah\n",
	))
}
