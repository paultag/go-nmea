package nmea

func GSAFixStatus(s GPGSASentence) string {
	switch s.Fix.Status {
	case GSAFixStatusNoFix:
		return "No fix"
	case GSAFixStatus2DFix:
		return "2-D Fix"
	case GSAFixStatus3DFix:
		return "3-D Fix"
	default:
		return "Unknown"
	}
}
