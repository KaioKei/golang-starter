# PKG DIRECTORY

It contains the library code that's ok to use by external applications (e.g., /pkg/mypubliclib). 

Other projects will import these libraries expecting them to work, so think twice before you put something here :-)  
Note that the internal directory is a better way to ensure your private packages are not importable because it's 
enforced by Go. The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe
for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and
internal directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and 
directories making it easier to run various Go tools (as mentioned in these talks: Best Practices for Industrial 
Programming from GopherCon EU 2018, GopherCon 2018:  
Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

See the /pkg directory if you want to see which popular Go repos use this project layout pattern. This is a common 
layout pattern, but it's not universally accepted and some in the Go community don't recommend it.

It's ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value 
(unless you really want to :-)). Think about it when it's getting big enough and your root directory gets pretty busy 
(especially if you have a lot of non-Go app components).

The pkg directory origins: The old Go source code used to use pkg for its packages and then various Go projects in the 
community started copying the pattern (see this Brad Fitzpatrick's tweet for more context).

## Create new module

Run :

```sh
cd pkg
mkdir mymodule
cd mymodule
go mod init kaiokei.com/mymodule
```