package totp

import (
	"testing"
	"time"
)

func BenchmarkGenerateTOTP(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Generate("KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", time.Unix(1614556800, 0))
	}
}

func BenchmarkAuthenticate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Authenticate("KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", time.Unix(1614556800, 0), "721396")
	}
}
