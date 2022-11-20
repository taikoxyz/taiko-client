package prover

func (s *ProverTestSuite) TestStartSubscription() {
	s.NotPanics(s.p.startSubscription)
	s.NotPanics(s.p.closeSubscription)
}
