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