#include "/usr/local/go/src/pkg/runtime/runtime.h"
#include "/usr/local/go/src/cmd/ld/textflag.h"

// see also: http://dave.cheney.net/2013/09/07/how-to-include-c-code-in-your-go-package

#pragma textflag NOSPLIT
void
·C1(int64 a, int64 b)
{
	b = a+1;
	FLUSH(&b);
}
