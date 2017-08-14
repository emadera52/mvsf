# mvsf
Test golang methods vs functions on structs, both *local* and *remote* 

### Background
This project came about when I wanted to define methods for a struct that was defined in a different package. The official GO documentation makes it clear that methods must be defined in the package that defines the srtuct. Not so official information confirms that but explains that it is possible to get the desired effect by using something called "embedding" which involves creating an *alias* for the target struct in a different package. Since new methods are defined on the *local alias*, the compiler is happy with them.

For the purposes of this project, the package than defines the struct is called *local* and the package where we want to define methods for it is called *remote*. I was able to confirm that this works, though I never found a good explanation as to why. It seems that this is just one more magical power possessed by empty structs.

After seeing the extra steps involved in setting up the magic and the way the magic is accessed when needed, I was sure there would be a performance hit when using this trick. This program started as a way to find out how much of a hit. Man, was I surprised.

### The Program
I've commented the code such that there is no need for a long explanation here. By the time it was finished, four tests were included as you can see in main.go. These include tests for both *local* and *remote* **methods** and **functions**. The shocker is that all the other options take roughly 35% longer than using *remote* **methods**. I will not be shocked if someone points out that there is a flaw in my tests. Otherwise, I'd love to know why this happens.

### Examples
**mvsf function 200000 loop** runs the *local* **function** test 200,000 times inside a loop that repeats 40 times. Only the average duration for the 40 runs is shown. 

**mvsf method 500000** runs the *local* **method** test 500,000 times. Results are shown at the end of the test including the average duration.

**mvsf func** runs the *remote* **function** test 100,000 times, that being the minimum. Results are printed at the end of the test, including the average duration. Note that **msvf func cat** and **msvf func 42000** will produce the same results.

**mvsf meth 200000 dog** runs the *remote* **method** test 200,000 times inside a loop that repeats 40 times. Only the average duration for the 40 runs is shown. This also illustrates that **any** third parameter triggers the outer loop.
