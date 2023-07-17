module example.com/hello

go 1.20

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.11.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace example.com/greetings => ./greetings
