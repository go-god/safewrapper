# safewrapper
wrap go func and waitgroup exec goroutine In a safe way.
# Why encapsulate this package? 
- if panic occurs during the execution of go func,the entire program may exit if the current program does not recover more than once.
- Because goroutine cannot capture panic across coroutines, panic must be handled well in the current execution of func.
- This approach is `a defensive approach to programming`.
