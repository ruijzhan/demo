package contract

import "net/http"

const KernelKey = "my:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
