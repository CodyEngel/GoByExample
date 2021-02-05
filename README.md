# Go by Example

This goes through the different samples which are available at [Go by Example](https://gobyexample.com/)

## Progress

This outlines the examples which have already been completed and are in order of completion (oldest -> newest).
1. [Hello World](https://gobyexample.com/hello-world)
2. [Values](https://gobyexample.com/values)
3. [Variables](https://gobyexample.com/variables)
4. [Constants](https://gobyexample.com/constants)
5. [For](https://gobyexample.com/for)
6. [If/Else](https://gobyexample.com/if-else)
7. [Switch](https://gobyexample.com/switch)
8. [Arrays](https://gobyexample.com/arrays)
9. [Slices](https://gobyexample.com/slices)
10. [Maps](https://gobyexample.com/maps)
11. [Range](https://gobyexample.com/range)
12. [Functions](https://gobyexample.com/functions)
13. [Multiple Return Values](https://gobyexample.com/multiple-return-values)
14. [Closures](https://gobyexample.com/closures)
15. [Recursion](https://gobyexample.com/recursion)
16. [Pointers](https://gobyexample.com/pointers)
17. [Interfaces](https://gobyexample.com/interfaces)
18. [Errors](https://gobyexample.com/errors)
19. [Goroutines](https://gobyexample.com/goroutines)
20. [Channels](https://gobyexample.com/channels) are the pipes that connected concurrent goroutines. You can send values
into channels from one goroutine and receive those values into another goroutine. [Source](/main/channels.go)
21. [Channel Buffering](https://gobyexample.com/channel-buffering) by default channels are unbuffered, meaning that they 
will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered 
channels accept a limited number of values without a corresponding receiver for those values. [Source](/main/channel-buffering.go)
22. [Channel Synchronization](https://gobyexample.com/channel-synchronization) we can use channels to synchronize execution 
across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for 
multiple goroutines to finish, you may prefer to use a WaitGroup. [Source](/main/channel-synchronization.go)
23. [Channel Directions](https://gobyexample.com/channel-directions) when using channels as function parameters, you can 
specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.
[Source](/main/channel-directions.go)
24. [Select](https://gobyexample.com/select) Go’s select lets you wait on multiple channel operations. Combining goroutines 
and channels with select is a powerful feature of Go. [Source](/main/select.go)
25. [Timeouts](https://gobyexample.com/timeouts) are important for programs that connect to external resources or that 
otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.
[Source](/main/timeouts.go)
26. [Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations) basic sends and receives 
on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, 
and even non-blocking multi-way selects. [Source](/main/non-blocking-channel-operations.go)
27. [Closing Channels](https://gobyexample.com/closing-channels) closing a channel indicates that no more values will be 
sent on it. This can be useful to communicate completion to the channel’s receivers. [Source](/main/closing-channels.go)
28. [Range over Channels](https://gobyexample.com/range-over-channels) in a previous example we saw how for and range 
provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.
[Source](/main/range-over-channels.go)
29. [Timers](https://gobyexample.com/timers) we often want to execute Go code at some point in the future, or repeatedly 
at some interval. Go’s built-in timer and ticker features make both of these tasks easy. We’ll look first at timers and 
then at tickers. [Source](/main/timers.go)
                                                                       
