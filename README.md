# Language features

- Ruby-like syntax + static typing
```
def print(): void
	puts("Haj 0.1")
	puts(">>> ")
end

def add(a: Integer, b: Integer): Integer
	a + b
end
```

- First-class functions
```
def map(array: Array<Any>, mapFn: Function): Array<Any>
	...
end

arry = [0, 1, 2]

def double(term: Integer): Integer
	term * 2
end

map(arry, double)
```
	- Notice here that `double` is referenced _without_ parentheses. This is
		how we reference a function in Haj (without invoking it)

- Anonymous functions
_TBD_
