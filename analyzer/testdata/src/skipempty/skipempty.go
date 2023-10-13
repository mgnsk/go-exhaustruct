//nolint:all
package skipempty

type Test struct {
	A string
}

func shouldPassFullyDefined() {
	_ = Test{}
}
