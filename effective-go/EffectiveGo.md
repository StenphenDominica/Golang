1. Examples
    Go 包的源码不仅是核心库，同时也是学习如何使用 Go 语言的示例代码。
    The Go package sources are intended to serve not only as the core library but also as examples of how to use the language.

2. Formatting
   Formatting issues are the most contentious but the least consequential.
   All Go code in the standard packages has been formated with **gofmt**.
   Some formatting details remain, Very briefly:
   ## Indentation
    We use tabs for indetation and gofmt emits them by default. Use spaces if you must.
   ## Line length 
    Go has no line lenth limit.
   ## Parentheses
    Go needs fewer parentheses than C and Java;

3. Commentary 
   1. Go provides C-style /* */ block comments and C++-style // comments. Line comments are the norm, block comments appear mostly as package comments.
   2. Every package should have a package  comment, a block comment preceding the package clause. For multi-file package, the package comment only needs to be present in one file, and anyone will do.
   3. Doc comments work best as complete sentence， which allow a wide variety of automated presentations. The first sentence should be a one-sentence summary that starts with the name being declared.

4. Names
   1. Names are important in Go as in any other language. The even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case.
   2. **Package Name**
      1. When a package is imported, the package name becomes an accessor for the contents, after ```import "bytes"``` ,the import package can talk about ```bytes.Buffer```.It`s helpful if everyone using the package use the same name to refer ti uts contents, which implies that package name should be good: short, concise, evcative.
      2. By convenntion, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. Package name need not be unique across all source code, and in the rare case of a collision the importing package can choose a different name to use locally.
      3. Another convention is that the package name is the base name of  its source directory; the package in ```src/encoding/base64``` is imported as ```"encoding/base64"```
5. 
6. 
7. 
8. 
