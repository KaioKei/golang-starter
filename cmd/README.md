# CMD DIRECTORY

It contains the main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects,
then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that
code in the /internal directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and
nothing else.

See the /cmd directory for examples.

## Install commands


From the command line in the hello directory, run the go build command to compile the code into an executable.

$ go build

From the command line in the hello directory, run the new hello executable to confirm that the code works.

Note that your result might differ depending on whether you changed your greetings.go code after testing it.
On Linux or Mac:

$ ./hello
map[Darrin:Great to see you, Darrin! Gladys:Hail, Gladys! Well met! Samantha:Hail, Samantha! Well met!]

On Windows:

$ hello.exe
map[Darrin:Great to see you, Darrin! Gladys:Hail, Gladys! Well met! Samantha:Hail, Samantha! Well met!]

You've compiled the application into an executable so you can run it. But to run it currently, your prompt needs either to be in the executable's directory, or to specify the executable's path.

Next, you'll install the executable so you can run it without specifying its path.
Discover the Go install path, where the go command will install the current package.

You can discover the install path by running the go list command, as in the following example:

$ go list -f '{{.Target}}'

For example, the command's output might say /home/gopher/bin/hello, meaning that binaries are installed to /home/gopher/bin. You'll need this install directory in the next step.
Add the Go install directory to your system's shell path.

That way, you'll be able to run your program's executable without specifying where the executable is.
On Linux or Mac, run the following command:

$ export PATH=$PATH:/path/to/your/install/directory

On Windows, run the following command:

$ set PATH=%PATH%;C:\path\to\your\install\directory

As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

$ go env -w GOBIN=/path/to/your/bin

or

$ go env -w GOBIN=C:\path\to\your\bin

Once you've updated the shell path, run the go install command to compile and install the package.

$ go install

Run your application by simply typing its name. To make this interesting, open a new command prompt and run the hello executable name in some other directory.

$ hello
map[Darrin:Hail, Darrin! Well met! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met!]
