module day1

go 1.23.0

require github.com/adam-lavrik/go-imath v0.0.0-20210910152346-265a42a96f0b // indirect
fun myFold (default, func, nil) = default | 
  myFold (default, func, head::tail) = func(head, myFold (default, func, tail));