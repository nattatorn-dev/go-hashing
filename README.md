# Bcrypt

## bcrypt no need to store salt?

The salt is automatically included in the output string which means there is no need to add it by yourself.

## Run

```bash
$ go run main.go
Password: secret
Hash:     $2a$14$98XP6TZdyN6ir/q1I7yQLeQQvb60gajYQ9fdaNmb51Wgz4nSLB6oS
Match:    true
```

## Unit test

```bash
$ go test
```

## Benchmark

BenchmarkHashPassword/password-16                    475          54780659 ns/op

ns/op = Operation per second
total 475 operation

```bash
$ go test -bench=. -benchtime=20s

BenchmarkHashPassword/password-16                    475          54780659 ns/op
BenchmarkHashPassword/password1234password99-16                      428          53780665 ns/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99-16                      452          53281311 ns/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+-16                      448          53139408 ns/op
BenchmarkHashPassword/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-16                450          52942657 ns/op
BenchmarkHashPassword/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-16                 447          53347687 ns/op
BenchmarkCheckPasswordHash/password-16         453          52211613 ns/op
BenchmarkCheckPasswordHash/password1234password99-16         458          55415961 ns/op
BenchmarkCheckPasswordHash/pa$$w@rd1234[assw@rd99-16         414          53107599 ns/op
BenchmarkCheckPasswordHash/-Ey2@bV=uTK6K%Np@sQCM+-16         468          51518254 ns/op
BenchmarkCheckPasswordHash/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-16         456          52328948 ns/op
BenchmarkCheckPasswordHash/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-16            459          50759361 ns/op
PASS
ok      github.com/nattatorn-dev/go-hashing     351.591s

```
