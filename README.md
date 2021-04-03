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
$ go test -bench=. -benchtime=20s -benchmem -cpu 1,2,4,8

BenchmarkHashPassword/8_Charaters           	      22	  53655842 ns/op	    5200 B/op	      11 allocs/op
BenchmarkHashPassword/8_Charaters-2         	      22	  49841241 ns/op	    5216 B/op	      11 allocs/op
BenchmarkHashPassword/8_Charaters-4         	      24	  51418096 ns/op	    5215 B/op	      11 allocs/op
BenchmarkHashPassword/8_Charaters-8         	      24	  50001819 ns/op	    5237 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters          	      22	  49836426 ns/op	    5232 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters-2        	      19	  53050153 ns/op	    5251 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters-4        	      22	  53042847 ns/op	    5249 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters-8        	      24	  51316475 ns/op	    5278 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols           	      20	  52089776 ns/op	    5233 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols-2         	      24	  50659190 ns/op	    5246 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols-4         	      20	  52768444 ns/op	    5262 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols-8         	      22	  51694814 ns/op	    5283 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower           	      22	  49555180 ns/op	    5232 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2         	      21	  52086177 ns/op	    5249 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4         	      22	  50001451 ns/op	    5249 B/op	      11 allocs/op
BenchmarkHashPassword/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8         	      22	  51993856 ns/op	    5283 B/op	      11 allocs/op
BenchmarkHashPassword/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower           	      24	  51957422 ns/op	    5327 B/op	      12 allocs/op
BenchmarkHashPassword/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2         	      21	  50177124 ns/op	    5334 B/op	      12 allocs/op
BenchmarkHashPassword/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4         	      24	  51489449 ns/op	    5353 B/op	      12 allocs/op
BenchmarkHashPassword/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8         	      22	  50173670 ns/op	    5369 B/op	      12 allocs/op
BenchmarkHashPassword/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower          	      24	  51771974 ns/op	    5567 B/op	      12 allocs/op
BenchmarkHashPassword/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2        	      24	  51305974 ns/op	    5573 B/op	      12 allocs/op
BenchmarkHashPassword/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4        	      21	  50731874 ns/op	    5597 B/op	      12 allocs/op
BenchmarkHashPassword/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8        	      24	  51334670 ns/op	    5614 B/op	      12 allocs/op
BenchmarkCheckPasswordHash/8_Charaters                                              	      22	  49810842 ns/op	    5284 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/8_Charaters-2                                            	      24	  51628031 ns/op	    5289 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/8_Charaters-4                                            	      24	  51398645 ns/op	    5299 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/8_Charaters-8                                            	      21	  50068013 ns/op	    5337 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters                                             	      24	  51236630 ns/op	    5315 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters-2                                           	      24	  50352176 ns/op	    5330 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters-4                                           	      22	  51120266 ns/op	    5344 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters-8                                           	      21	  51271494 ns/op	    5369 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols                             	      22	  51474334 ns/op	    5316 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols-2                           	      24	  51648477 ns/op	    5330 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols-4                           	      22	  52980756 ns/op	    5344 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols-8                           	      24	  53102654 ns/op	    5362 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower      	      22	  51436490 ns/op	    5316 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2    	      24	  51329015 ns/op	    5330 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4    	      22	  51403930 ns/op	    5344 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/22_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8    	      24	  52862452 ns/op	    5353 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower      	      22	  51471613 ns/op	    5412 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2    	      24	  51590277 ns/op	    5426 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4    	      24	  51024950 ns/op	    5437 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/44_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8    	      22	  51819826 ns/op	    5463 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower     	      22	  52945964 ns/op	    5652 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-2   	      21	  52481942 ns/op	    5658 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-4   	      22	  51192954 ns/op	    5680 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/128_Charaters_Include_Symbols,_Numbers,_Upper,_Lower-8   	      24	  51115729 ns/op	    5698 B/op	      15 allocs/op
PASS
ok  	github.com/nattatorn-dev/go-hashing	62.633s


```
