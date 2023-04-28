# 1. Examples
The Go package sources are intended to serve not only as the core library but also as examples of how to use the language.  
Go 包的源码不仅是核心库，同时也是学习如何使用 Go 语言的示例代码。

# 2. Formatting
Formatting issues are the most contentious but the least consequential.  
All Go code in the standard packages has been formated with **gofmt**.  
Some formatting details remain, Very briefly:
## Indentation
We use tabs for indetation and gofmt emits them by default. Use spaces if you must.
## Line length 
Go has no line lenth limit.
## Parentheses
Go needs fewer parentheses than C and Java;

# 3. Commentary 
   1. Go provides C-style ```/* */ ```block comments and C++-style ```//``` comments. Line comments are the norm, block comments appear mostly as package comments.
   2. Every package should have a package  comment, a block comment preceding the package clause. For multi-file package, the package comment only needs to be present in one file, and anyone will do.
   3. Doc comments work best as complete sentence， which allow a wide variety of automated presentations. The first sentence should be a one-sentence summary that starts with the name being declared.

# 4. Names
Names are important in Go as in any other language. The even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case.  

## Package Name
1. When a package is imported, the package name becomes an accessor for the contents, after ```import "bytes"``` ,the import package can talk about ```bytes.Buffer```. It`s helpful if everyone using the package use the same name to refer ti uts contents, which implies that package name should be good: short, concise, evcative.  
2. By convenntion, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. Package name need not be unique across all source code, and in the rare case of a collision the importing package can choose a different name to use locally.
3. Another convention is that the package name is the base name of  its source directory; the package in ```src/encoding/base64``` is imported as ```"encoding/base64"```  

## Getters  
Go doesn't provide automatic support for gettes and setters. You can providing getters and setters yourself, but it's neither idiomatic nor necessary to put a Get into the getter's name.
If you have a field called owner(lower case, unexported), the getter method should be called Owner (Upper case, exported), not GetOwner. The use of upper-case names for export provides the hook to discriminate the filed from the method. A setter function, if needed, will likely be called SetOwner.

```owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```  

## Interface names
By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to sonstruct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

## MixedCaps  
In Go is to use **MixedCaps** or **mixedCaps** rather than underscores to write multiword names.  

# Semicolons
Like C, Go's formal grammar uses semicolons to terminate statements, but unlike C, those semicolons do not appear in the source. Instead the lexer uses a simple rlue to insert semicolons automatically as it scans, so the input text is mostly free of them.
The lexer inserts a semicolons after the token, This could be summarized as, "if the newline comes after a token that could end a statement, insert a semicolon".
Idiomatic Go programs have semicolons only in palces such as **for** loop clauses, to separate the initializer, contition, and continuation elements.

# Control structures
The control structures of Go are related to those of C but differ in important ways. 
There is no **do** or **while** loop, only a slightly generalized **for**;
**switch** is more flexible;
**if** and **switch** accept an optinal initialization statement like that of **for**;
**break** and **continue** statements take an optional label to identify what to break or continue;
and there are new control structures including a type switch and a multiway communications multiplexer, **select**.
The syntax is also slightly different: there are no parentheses the bodies must always be brace-delimited.

## if

Mandatory braces encourage writing simple **if** statement on multiple lines. A simple **if** looks like this:
```
if x > 0 {
    return y
}
```

This is an example of a common situation where code must guard against a sequence ofn error conditions. The code reads well if the successful flow of control runs down the 


page, eliminating error cases as they arise. Since error cases tend to end in **return** statements, the resulting code needs no **else** statements.
```
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    return err
}
codeUsing(f, d)
```
The previous example demonstrates a detail of how the := short declaration form works.

## Redeclaration and reassignment
The last example in the previous setion, We Notice that err appears in both statements. This is duplication is legal: **err** is declared by the first statement, but only re-assigned in the second. This means that the call to **f.Stat** uses the existing err variable declared above, and just gives it a new value. 
For example, in long **if-else** chain. You'll see it used often. This unusual property is pure pragmatism.

It's worth noting here that in Go the scope of funtion parameters and return values in the same as the function body, even though they appear lexically outside the braces that enclose the body.

## for
