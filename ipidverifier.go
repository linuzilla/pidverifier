package pidverifier

type IPidVerifier interface {
	Verify(str string) bool
}
