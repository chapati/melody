package melody

type envelope struct {
	t      int
	msg    []byte
	filter filterFunc
	msgf   msgFunc
}

type envelopeex struct {
	t     int
	msgf  msgFunc
}
