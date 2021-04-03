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

BenchmarkHashPassword/password-16                    410          63111774 ns/op            5199 B/op         11 allocs/op

- 410 is the number of iterations for i := 0; i < b.N; i++
- 63111774 ns/op is approximate time it took for one iteration to complete
- 5199 B/op allocs/op means how many distinct memory allocations - occurred per op (single iteration).
- 11 allocs/op B/op is how many bytes were allocated per op.


```bash
$ go test -bench=. -benchtime=20s -benchmem
BenchmarkHashPassword/password-16                    410          63111774 ns/op            5199 B/op         11 allocs/op
BenchmarkHashPassword/password1234password99-16                      372          62160115 ns/op            5231 B/op         11 allocs/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99-16                      374          64401856 ns/op            5232 B/op         11 allocs/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+-16                      373          65280327 ns/op            5232 B/op         11 allocs/op

```
