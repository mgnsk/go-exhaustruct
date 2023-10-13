//nolint:all
package exported

type Test struct {
	A string
	b string
}

func shouldPassFullyDefined() {
	_ = Test{
		A: "",
	}
}
