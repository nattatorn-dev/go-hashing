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

BenchmarkHashPassword/password           	      21	  55764515 ns/op	    5200 B/op	      11 allocs/op
BenchmarkHashPassword/password-2         	      19	  64726420 ns/op	    5219 B/op	      11 allocs/op
BenchmarkHashPassword/password-4         	      18	  60560375 ns/op	    5223 B/op	      11 allocs/op
BenchmarkHashPassword/password-8         	      19	  55912817 ns/op	    5248 B/op	      11 allocs/op
BenchmarkHashPassword/password1234password99           	      22	  59337395 ns/op	    5232 B/op	      11 allocs/op
BenchmarkHashPassword/password1234password99-2         	      18	  60515238 ns/op	    5252 B/op	      11 allocs/op
BenchmarkHashPassword/password1234password99-4         	      20	  57137659 ns/op	    5262 B/op	      11 allocs/op
BenchmarkHashPassword/password1234password99-8         	      19	  60524345 ns/op	    5291 B/op	      11 allocs/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99           	      19	  58750307 ns/op	    5233 B/op	      11 allocs/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99-2         	      18	  59794117 ns/op	    5252 B/op	      11 allocs/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99-4         	      19	  61186354 ns/op	    5253 B/op	      11 allocs/op
BenchmarkHashPassword/pa$$w@rd1234[assw@rd99-8         	      19	  62356671 ns/op	    5291 B/op	      11 allocs/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+           	      18	  57546408 ns/op	    5234 B/op	      11 allocs/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+-2         	      20	  55025752 ns/op	    5250 B/op	      11 allocs/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+-4         	      20	  58902678 ns/op	    5262 B/op	      11 allocs/op
BenchmarkHashPassword/-Ey2@bV=uTK6K%Np@sQCM+-8         	      19	  62219079 ns/op	    5291 B/op	      11 allocs/op
BenchmarkHashPassword/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ           	      19	  60840218 ns/op	    5329 B/op	      12 allocs/op
BenchmarkHashPassword/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-2         	      21	  61964186 ns/op	    5345 B/op	      12 allocs/op
BenchmarkHashPassword/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-4         	      19	  59514040 ns/op	    5349 B/op	      12 allocs/op
BenchmarkHashPassword/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-8         	      18	  56320785 ns/op	    5380 B/op	      12 allocs/op
BenchmarkHashPassword/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j           	      22	  55823934 ns/op	    5568 B/op	      12 allocs/op
BenchmarkHashPassword/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-2         	      22	  52731598 ns/op	    5574 B/op	      12 allocs/op
BenchmarkHashPassword/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-4         	      22	  53199160 ns/op	    5585 B/op	      12 allocs/op
BenchmarkHashPassword/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-8         	      21	  52678357 ns/op	    5622 B/op	      12 allocs/op
BenchmarkCheckPasswordHash/password                                                                                                                              	      22	  53254083 ns/op	    5284 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password-2                                                                                                                            	      21	  53406320 ns/op	    5290 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password-4                                                                                                                            	      21	  53007317 ns/op	    5302 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password-8                                                                                                                            	      24	  53065103 ns/op	    5330 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password1234password99                                                                                                                	      22	  52736721 ns/op	    5316 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password1234password99-2                                                                                                              	      22	  53628764 ns/op	    5322 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password1234password99-4                                                                                                              	      20	  52612469 ns/op	    5347 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/password1234password99-8                                                                                                              	      22	  52502132 ns/op	    5367 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/pa$$w@rd1234[assw@rd99                                                                                                                	      22	  52465976 ns/op	    5316 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/pa$$w@rd1234[assw@rd99-2                                                                                                              	      22	  52995401 ns/op	    5332 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/pa$$w@rd1234[assw@rd99-4                                                                                                              	      22	  53317756 ns/op	    5344 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/pa$$w@rd1234[assw@rd99-8                                                                                                              	      20	  55654915 ns/op	    5373 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/-Ey2@bV=uTK6K%Np@sQCM+                                                                                                                	      20	  54785142 ns/op	    5317 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/-Ey2@bV=uTK6K%Np@sQCM+-2                                                                                                              	      22	  52649042 ns/op	    5322 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/-Ey2@bV=uTK6K%Np@sQCM+-4                                                                                                              	      22	  52964064 ns/op	    5344 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/-Ey2@bV=uTK6K%Np@sQCM+-8                                                                                                              	      20	  62068443 ns/op	    5362 B/op	      14 allocs/op
BenchmarkCheckPasswordHash/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ                                                                                          	      21	  54603081 ns/op	    5412 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-2                                                                                        	      19	  57787895 ns/op	    5420 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-4                                                                                        	      21	  58602505 ns/op	    5441 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/e%4a-DdEtfJbfxfZnjw46#Q8T5M5#8!Njj%Ak%XSM=vQ-8                                                                                        	      19	  61529464 ns/op	    5472 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j      	      22	  55766077 ns/op	    5652 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-2    	      20	  56247692 ns/op	    5670 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-4    	      19	  57270642 ns/op	    5674 B/op	      15 allocs/op
BenchmarkCheckPasswordHash/c7vESMnyTm57y429vrHSk4*?b2cysBqTTZVf67EP!_#5_P!efAbd2haWM-Sd6AU5TxxFV5aZAc5QJsh=VX^dqA2SNe$CpkaPgE33pSZGcSjV-U?ksJEm5uH6Mg@qDf2j-8    	      21	  55931019 ns/op	    5705 B/op	      15 allocs/op
PASS
ok  	github.com/nattatorn-dev/go-hashing	62.713s

```
