package main

import (
	"testing"
)

func Test(t *testing.T) {
	password := "password"
	hash, _ := HashPassword(password)
	match := CheckPasswordHash(password, hash)
	if !match {
		t.Error()
	}
}

func BenchmarkHashPassword(b *testing.B) {
	benchmarks := []struct {
		useCase  string
		password string
	}{
		{"8 Charaters", "password"},
		{"22 Charaters", "password1234password99"},
		{"22 Charaters Include Symbols", "pa$$w@rd1234[assw@rd99"},
		{"22 Charaters Include Symbols, Numbers, Upper, Lower", "-Ey2@bV=uTK6K%Np@sQCM+"},
		{"44 Charaters Include Symbols, Numbers, Upper, Lower", "e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ"},
		{"128 Charaters Include Symbols, Numbers, Upper, Lower", "c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j"},
	}
	for _, benchmark := range benchmarks {
		funcuntion := GetHashPassword(benchmark.password)
		b.Run(benchmark.password, funcuntion)
	}
}

func GetHashPassword(password string) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HashPassword(password)
		}
	}
}

func BenchmarkCheckPasswordHash(b *testing.B) {
	benchmarks := []struct {
		useCase  string
		password string
		hash     string
	}{
		{"8 Charaters", "password", "$2a$10$l9TQ3eKJ25XfnfcQsoLzyOSZx29/M2A.q8tx4UACWmwSxxPj1c996"},
		{"22 Charaters", "password1234password99", "$2a$10$NMW5h./p5vb0o937CW7tCe.nigL/3RS0AwWX21Mn9AHV1U9rrBY52"},
		{"22 Charaters Include Symbols", "pa$$w@rd1234[assw@rd99", "$2a$10$6h.ZzMUH69sVeCC6S1P2HufhNfXFYrafCNamgG4wz3DfHFF6RAXiq"},
		{"22 Charaters Include Symbols, Numbers, Upper, Lower", "-Ey2@bV=uTK6K%Np@sQCM+", "$2a$10$4TWMJQDiuZddpRnWwlk74O84PXOylTsZzA8NWidIU1DLfSKh8eVyW"},
		{"44 Charaters Include Symbols, Numbers, Upper, Lower", "e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ", "$2a$10$tP.9/8Gb70L5lCD6ZFU2S.S9/shFfObYkeQ8GkkfUtb2.wjXccv.O"},
		{"128 Charaters Include Symbols, Numbers, Upper, Lower", "c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j", "$2a$10$MI.zSJaxe8hSS6mapmXPUu3AtyTGjVlcd2S75OBWpKWR5FiOQqrv."},
	}
	for _, benchmark := range benchmarks {
		funcuntion := GetPasswordHashChecking(benchmark.password, benchmark.hash)
		b.Run(benchmark.password, funcuntion)
	}
}

func GetPasswordHashChecking(password, hash string) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CheckPasswordHash(password, hash)
		}
	}
}
