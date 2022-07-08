package tutorial

import (
	// Built-in packages have import names and may have nested packages
	"log"
	"math/big"
	// Local packages have import paths relative to the module's path (root of the go.mod file)
	"golang_starter/pkg/tutorial"
	"golang_starter/pkg/tutorial/nested"
	// Third-party packages have import url
	// to install them, use 'go get -u <package url>
	// they also have import nested packages
	"rsc.io/quote"
)

func packages() {
	// built-in
	log.Printf("packages tutorial")
	ten := big.NewInt(10)
	log.Printf("int64: %d", ten)

	// local
	// this function is Visible outside the 'tutorial' package because its name starts with an uppercase
	tutorial.Visible()
	// two functions may be called the same way from different packages
	// even if the packages are nested
	nested.Visible()
	// this function uses a random algorithm using a different seed each time the program is launched
	// thanks to the 'init' function of the package, which is called implicitly during import
	tutorial.GetRandomNumber(10)

	// third-party
	log.Println(quote.Go())
}
