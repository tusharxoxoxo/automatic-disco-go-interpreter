# automatic-disco-go-interpreter
another tree walking interpreter in go

Interpreter take source code and evaluate it without producing some visible, intermediate result that can later be executed while a compilers, which take source code and produce output in another language that the underlying system can understand.
 
In this, i have built my own lexer, my own parser, my own tree representation, and my very own evaluator. The interpreter i have built is the famous tree-walking interpreter

we are going to have our very own programming language for our interpreter
![th-3839655494](https://github.com/tusharxoxoxo/automatic-disco-go-interpreter/assets/79051850/5b980d07-e0eb-43eb-95c6-e6bcc6248c0b)

GoblinShark (the language) has the following features:
• C-like syntax
• variable bindings
• integers and booleans
• arithmetic expressions
• built-in functions
• first-class and higher-order functions • closures
• a string data structure
• an array data structure
• a hash data structure

here is some code example for our GoblinShark programming language, like the syntax it's going to follow
Here is how we bind values to names in GoblinShark:

```let age = 11;```<br />
```let name = "GoblinShark";```<br />
```let result = 1000 * (20 / 2);```<br />

Besides integers, booleans, and strings, the GoblinShark interpreter we’re going to build will also support arrays and hashes.
Here’s what binding an array of integers to a name looks like:

```let myArray = [1, 2, 3, 4, 5];```

And here is a hash, where values are associated with keys:

```let Tushar = {"name": "Tushar", "age": 21};```

Accessing the elements in arrays and hashes is done with index expressions:

```myArray[0] // => 1 tushar["name"] // => "Tushar"```

GoblinShark not only supports return statements. Implicit return values are also possible,
which means we can leave out the return if we want to:

```let add = fn(a, b) { a + b; };```
```add(1, 2);```

GoblinShark also supports a special type of function, called higher-order functions. These are
functions that take other functions as arguments. Here is an example:
```
let twice = fn(f, x){
    return f(f(x));
  };
let addFive = fn(x) {
    return x + 5;  
  };
twice(addTwo, 2); // => 12
```

Functions in GoblinShark are just values, like integers or strings. That feature is called “first class functions”.
The following are the features our interpreter for the GoblinShark language is going to have: 
• the lexer -> takes source code as input and output the tokens that represent source code
• the parser
• the Abstract Syntax Tree (AST) 
• the internal object system
• the evaluator

The actual interpreter will attach filenames and line number to tokens, for tracking down lexing and parsing errors but we just creating a simple interpreter just to learn stuff so we are excluding this part to reduce complexity of our interpreter.
So we are just using strings and ignoring filenames and line numbers.

