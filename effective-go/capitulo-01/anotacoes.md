# Intro

## Objetivos do livro

Principal objetivo do livro: ensinar como escrever código com bom design e
manutenível em Go.

Código bom "bom design" é simples, mas simples não é fácil.
Código com "bom design" é fácil de entender, de mudar, confiável, com menos bugs
e ajuda a evitar surpresas.

Normalmente há uma forma correta de fazer coisas em Go - Go Idiomático.

    Simples - Direto, fácil de ler e entender. Você pode entender o custo de hardware de quase tudo que é feito
    Adaptável - Go segue a filosofia UNIX é é feito com compatibilidade em mente ao invés de herdar comportamentos
    Testável - Código fácil de testar


## Código testável

Testes são cidadãos de primeira classe em Go, e a stdlib tem bom suporte.

Código deveria se adaptar a nocos requisitos e resistir ao teste do tempo. Com
testes é mais fácil atingir essa qualidade, pois permite saber se o código continua
fazendo o que deveria fazer após mudanças.

Testes manuais são fáceis de errar e díficil de executar em códigos grandes.

Testes automatizados resolvem isso, mas é preciso escrever código testável.

#### Benefícios:
* Confiança;
* Desig modular
* Menos bugs
* Fácil debuggar
* Documentação

#### Desvantagens
* Mais código
* Mais trabalho (apesar de passar menos tempo procurando erros e testatando manualmente)
* Testes podem se tornam o objetivo principal


## Motivações da linguagem

Robert Griesemer, Rob Pike, e Ken Thompson trabalhavam no Google e estavam lidando
com tempos de compilação longos, características complexas de linguagem e código 
difícil de entender.

C,C++ e Java são performaticas, mas não são amigáveis. Pytho, Ruby e PHP são amigáveis,
mas não são performatícas.

Influências:


    C-like statement and expression syntax.
    Pascal-like declaration syntax.
    Oberon-like packaging system. Instead of using public, private, and protected keywords to manage access to an identifier, Go and Oberon use a simple mechanism to export an identifier from a package. Oberon, like Go, when you import a package, you need to qualify the package's name to access the exported identifiers. Go exports when you capitalize the first letter, and Oberon does so when adding an asterisk.
    Smalltalk-like object-oriented programming style. Developers from other object-oriented programming languages to Go are often surprised when they can't see any classes. There is no concept of class: Data and behavior are two distinct concepts in Go.
    Smalltalk-like duck-typing style in which you can pass a value to any type that expects a set of behaviors. You can see the same feature in other popular languages like Ruby and Python. But what makes Go different in this case is that Go provides type-safety and duck-typing at the same time.
    Newsqueak-like concurrency features. Newsqueak was another language created by Rob Pike.
    An object file format from Modula.

### The reasons behind Go's success
#### Opinionated

There is often one right way to do things in Go. There are no tabs vs. spaces arguments in Go. It formats code in a standard style. Refuses to compile when there are unused variables and packages. Encourages packages to be simple and coherent units that mimic the Unix way of building software. Refuses to compile when there is a cyclic dependency between packages. Its type system is strict and does not allow inheritance, and the list goes on.

#### Simplicity

The language is easy to work with, concise, explicit, and easy to read and understand. It's minimal and easy to learn in a week or so. There is a 50-page long specification that defines the mechanics of the Go language. Whenever confusion about some language feature occurs, you can get an authoritative answer from the spec. The backward compatibility of Go guarantees that even though Go evolves each day, the code you wrote ten years ago still works today.

#### Type system and concurrency

Go is a strongly and statically typed programming language and takes the best tenets of object-oriented programming like composition. The compiler knows every value type and warns you if you make a mistake. Go is a modern language with built-in support for concurrency, and it's ready to tackle today's large-scale distributed applications.

#### Built-in packages and tools

Maybe newcomers believe that the Go Standard Library—stdlib—lacks features and depends on third-party code. But in reality, Go comes with a rich set of packages. For example, there are packages for writing command-line tools, http servers/clients, network programs, JSON encoders/decoders, file management, etc. Once newcomers have had enough experience in Go, most get rid of the third-party packages and prefer using the Standard Library instead.

When you install Go, it comes with many built-in tools that help you develop Go programs effectively: A compiler, tester, package manager, code formatters, static code analyzers, linters, test coverage, documentation, refactoring, performance optimization tools, and more.
#### Go compiler

There is no intermediary execution environment such as an interpreter or a virtual machine. Go compiles code directly to fast native machine code. It's also cross-platform: You can compile and run your code on major operating systems like OS X, Linux, Windows, and more.

One of the design goals of Go from the beginning of its design was a fast compilation. Go compiles so fast that you may think you're not even compiling your code. It feels as if you're working in an interpreted language like Python. A fast compiler makes you productive by quickly writing and testing your code. So how does it compile so fast:

    The language grammar is simple and easier to parse.
    Each source code file tells the compiler what the code should import at the top. So the compiler does not need to parse the rest of the file to find out what the file is importing.
    The compilation ends if a source code file does not use an imported package.
    The compiler runs faster because there are no cyclic dependencies. For example, if package A imports package B, package B cannot import package A.
    During compilation, the compiler records the dependencies of a package and the package's dependencies on an object file. Then the compiler uses the object file as a caching mechanism and compiles progressively faster for subsequent packages.



    |                Go Program               |
    | +-------------------------------------+ |
    | |   Statically Linked Packages        | |
    | +-------------------------------------+ |
    |                                         | <-----> SO
    | +-------------------------------------+ |
    | |              Go Runtime             | |
    | |  (Scheduler)     (Garbage Colector) | |
    | |       (OS Specific Packages )       | |
    | +-------------------------------------+ |
    |                                         |
    +-----------------------------------------+


#### Sistema de tipos

Orientação a objetos sem classes. Não tem classe nem herança. Cada tipo pode ter
seu comportamento. Ao invés de criar classes grandes que herdam, em go são criados
tipos compostos de outros. O foco é em passar mensagens entre os objetos (tipos)


