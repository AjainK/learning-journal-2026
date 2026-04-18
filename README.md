# Learning Journal 2026

## Day 0:
## Resources:
    Go programming language by Donovan
    golang.org
    blog.golang.org
    play.golang.org
    tour.golang.org( Built atop the Playground) (Primary shortcoming - only allow standard libraries to be imported.)
    golang.org/pkg (Read the code for any type or function in the standard library online . Use this to figure out how something works)
    gobyexample.com
    exercism.org
    Go the complete developers guide (Udemy)

## Day 1:

## Topics:
## Names
### Declarations: eg Fahrenheit to Celsius  
### Packages
    Types of packages:
    The name of the package determines whether we are creating an executable or a dependency type package.
    eg. When you run go build main.go we can see a file called main in the package, which we dont see when we run go run main.go. 

    Executable packages
    Generates a file that we can run.
    Must always have a func called main


    Reusable packages
    Code used as helpers . Good place to put the reusable logic.
### import
    Import statement is used to gain access to another package and its exported functions.
### Variables:
    var name type = expression // If expression is omitted , the initial value is the zero value for the type, which is 0 for numbers, false for booleans, "" for strings and nil for intrfaces and refrence types(slice, pointer, map, channel, function)
    var s string
    var i, j, k int //int, int, int
    var b, f, s = trur, 2.3, "four" //bool, float64, string
    var f, err = os.Open(name)
    // In Go there is no such thing as a uninitialized  variable

### Short variable declarations:
    t := 0.0
    freq := rand.Float64() * 3.0
    i,j := 0,1 

    in, err = os.Open(infile)
    out, err = os.Create(infile) //if some of the variables are already declared in the same lexical block , then the short variable declaration acts like an assignment to those variables.

    in, err = os.Open(infile)
    in, err = os.Open(infile) // compile error: no new variables . A short variable declaration must declare atleast one new variable

### Pointers
    x := 1
    p := &x
    fmt.Println(*p) // "1"
    *p = 2. // equivalent to x = 2
    fmt.println(x) // "2"

### The new Function
    The expression new(T) creates a new unnamed variable of type T , initializes it to the zero value of T and returns its address, whuch is a value of type *T
    p := new(T)
    fmt.Println(*p) // "0"
    *p = 2
    fmt.Println(*p) // "2"

    Each call to new() returns a distinct variable with a unique address:
    p := new(int)
    q := new(int)
    fmy.Println(p == q) // "false"

### Lifetime of Variables
    Package level variable - entire execution of the program
    Local variable - dynamic lifetimes
    The lifetime of a variable is determined only by whether or not it is reachable .

    var global *int 
    func f() {
        var x int
        x = 1
        global = &x
    } // Heap allocated because it is still reachable from the variable global after f has returned

    func g(){
        y := new(int)
        *y = 1
    } // becomes unreachable and can be recycled and its for the compiler to allocate *y on the stack

### fmt.Scanln
#### Behavior
    Stops at space
    Reads one word at a time
    Leaves leftover input in buffer (danger ⚠️)
    Think of it like: “Read the next word and stop”

### bufio.Reader
#### Behavior
Reads entire line
Includes spaces
Includes \n (you trim it)
Think of it like: “Read everything until Enter is pressed”

### Time
    time.Now() // 2026-04-17 17:51:55.343189 +0530 IST m=+0.000085667
    time.Now().Date() // 2026 April 17
    time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location()) //2026-04-30 00:00:00 +0530 IST


## Day 2:

### for 
    Go has only one looping construct, the for loop.
    The basic for loop has three components separated by semicolons:
    the init statement: executed before the first iteration
    the condition expression: evaluated before every iteration
    the post statement: executed at the end of every iteration
    The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
    The loop will stop iterating once the boolean condition evaluates to false.

    The init and post statements are optional.
    C's while is spelled for in Go.

    one more example:

    for index, card := range cards {
        fmt.Println(i, card) //index starts from 0
    }
    continue in for loop will restart the loop and break will let you out of the loop

    eg:
    for {
    fmt.Print("\nEnter your guess: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    guess, err := strconv.Atoi(input)
    if err != nil {
    fmt.Println("Please enter a valid number.")
    continue
    }

    if guess == secretNumber {
    fmt.Printf("🎉 Correct! You guessed it in %d attempt(s)!\n", attempts)
    break
    }

### if
    if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
    if sum == 0 {
        fmt.Println(sum)
    } else {
        fmt.Println("Sum is not zero")
    }
    This is a Go language rule—the parser expects else immediately after the closing brace. If you put it on a new line, Go's semicolon insertion rules treat the } as the end of a statement, making else unexpected.

    Like for, the if statement can start with a short statement to execute before the condition.Variables declared by the statement are only in scope until the end of the if.
    Variables declared inside an if short statement are also available inside any of the else blocks.

### Array 
    Fixed length list of things.

### Slice
    An array that can grow or shrink
    Append: To add more items to the Slice 
    eg. cards := []string {"ace of spades","two of hearts"}
        cards = append(cards, "five of diamonds")
    Append function does not modify the existing slice instead it creates a new slice 
    
### OO Approach vs Go Approach
    Classes that contain properties, functions like print, shuffle , saveToFile and class instance 

    * In go we extend a base type and add some extra functionality to it.

    eg type deck []string : Tell go we want to create an array of strings and add a bunch of functions specifically made to work with it.

    Functions with 'deck' as 'receiver'
    A function with a receiver is like a method - a function that belongs to an instance

### Receiver function
    Use a receiver method when the function belongs to a type — it operates on a specific struct's data. The function is part of that type's behaviour.

    Use a regular function when the logic is independent — it doesn't need to live on a type, or it works across multiple types.

    eg. func(d deck) newDeck() deck {
       for i, card := range d {
        fmt.Println(i, card)
       }
    }

### Slice range 
    eg.  0 Ace of Spades
        1 Two of Spades
        2 Three of Spades
        3 Four of Spades
  

        [0:2]  gives Ace of Spades and Two of Spades
        [ where to start : how many items to pick]
        
        [:3] = size (for deck of cards its called handsize(3 cards))
        starts from 0 and number of items
        Ace of Spades, Two of Spades, Three of Spades

        [3:] = starts from 2 and anything after that
        Four of Spades

### Multiple return values
    func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
    }

### WriteFile to the HardDrive
    WriteFile writes data to the named file, creating it if necessary. If the file does not exist, WriteFile creates it with permissions perm
    Uses the io package.
    func WriteFile(name string, data []byte, perm FileMode) error
    Needs a filename , byte slice (convert type string to []byte) and os perm = 0666(read, write access)

### ReadFile from the HardDrive
    ReadFile reads the named file and returns the contents. A successful call returns err == nil
    func ReadFile(name string) ([]byte, error)
     or func ReadFile(name string) ([]byte) where instead of returning the error we can check if error is thrown and handle it inside the function.
     	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

###  Rand.
    func New(src Source) *Rand
    New returns a new Rand that uses random values from src(seed) to generate other random values.

    A Source represents a source of uniformly-distributed pseudo-random int64 values in the range [0, 1<<63).
    func NewSource(seed int64) Source
    NewSource returns a new pseudo-random Source seeded with the given value.
    eg. 	
    source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
    for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}


### String to Int & Int to String
    strconv.Atoi(input)
    strconv.Itoa(input)

### switch {
    case condition1 : {
        return
    }
    case condition2 : {
        return
    }
    default :{
        return
    }
    }
