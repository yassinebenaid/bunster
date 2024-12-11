package tst

// TestCase represents a single test case with input and output
type Test struct {
	Label  string
	Input  string
	Output string
}

type parser struct {
	in   []byte
	pos  int
	curr byte
	next byte

	tests []Test
}

func (p *parser) proceed() {
	p.curr = p.next
	if p.pos >= len(p.in) {
		p.next = 0
	}
	p.pos++
}

// Parse reads from an io.Reader and parses the test case format
func Parse(in []byte) ([]Test, error) {
	p := parser{in: in}
	p.proceed()
	p.proceed()

	return p.tests, nil
}
