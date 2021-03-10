// name the package after the subdirectory it lives in so it is exported i.e. clients can
// say import"github.com/hectorMurilloIbarra43/gomodlib/lib1" and are able to use
// the function Bar
package lib1

func Bar() string {
	return "gomodlib/lib1 Bar"
}
