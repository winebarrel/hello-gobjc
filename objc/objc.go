package objc

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>

void hello();
*/
import "C"

// ShowNotification shows a notification to the user.
func Hello() {
	C.hello()
}
