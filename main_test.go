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

func TestGenerate(t *testing.T) {
	type args struct {
		k string
		t time.Time
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
		wantErr  bool
	}{
		{
			name:     "短",
			args:     args{k: "4OAYE===", t: time.Unix(0, 0)},
			wantCode: "449502",
			wantErr:  false,
		},
		{
			name:     "推荐",
			args:     args{k: "KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", t: time.Unix(0, 0)},
			wantCode: "233869",
			wantErr:  false,
		},
		{
			name:     "长",
			args:     args{k: "KL67YBZBKL67YBZBQJSU6FR7L4HZUYQ5OKKL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2NKWNR2NQJSU6FR7L4HZUYQ5OKKWNR2N", t: time.Unix(0, 0)},
			wantCode: "994273",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, err := Generate(tt.args.k, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("Generate() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	type args struct {
		k    string
		t    time.Time
		code string
	}
	tests := []struct {
		name        string
		args        args
		wantSuccess bool
		wantErr     bool
	}{
		{
			name:        "正确",
			args:        args{k: "KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", t: time.Unix(1614556800, 0), code: "721396"},
			wantSuccess: true,
			wantErr:     false,
		},
		{
			name:        "错误",
			args:        args{k: "KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", t: time.Now(), code: "721396"},
			wantSuccess: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuccess, err := Authenticate(tt.args.k, tt.args.t, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("Authenticate() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
