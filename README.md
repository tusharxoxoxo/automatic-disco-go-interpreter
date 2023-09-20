# automatic-disco-go-interpreter
another interpreter in go

In this, i have built my own lexer, my own parser, my own tree representation, and my very own evaluator. The interpreter i have built is the famous tree-walking interpreter

we are going to have our very own programming language for our interpreter

Expressed as a list of features, Donkey has the following:
• C-like syntax
• variable bindings
• integers and booleans
• arithmetic expressions
• built-in functions
• first-class and higher-order functions • closures
• a string data structure
• an array data structure
• a hash data structure

here is some code example for our donkey programming language, like the syntax it's going to follow
Here is how we bind values to names in Donkey:
```let age = 11;```
```let name = "Donkey";```
```let result = 1000 * (20 / 2);```

Besides integers, booleans, and strings, the Donkey interpreter we’re going to build will also support arrays and hashes.
Here’s what binding an array of integers to a name looks like:
```let myArray = [1, 2, 3, 4, 5];```

And here is a hash, where values are associated with keys:
```let Tushar = {"name": "Tushar", "age": 21};```

Accessing the elements in arrays and hashes is done with index expressions:
```myArray[0] // => 1 tushar["name"] // => "Tushar"```

Donkey not only supports return statements. Implicit return values are also possible,
which means we can leave out the return if we want to:
```let add = fn(a, b) { a + b; };```
```add(1, 2);```

Donkey also supports a special type of function, called higher-order functions. These are
functions that take other functions as arguments. Here is an example:
