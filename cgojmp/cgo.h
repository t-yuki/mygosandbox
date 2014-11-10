
#include <setjmp.h>

static jmp_buf jbuf;

struct GObuf {
	jmp_buf jbuf;
};

static int GOjump() {
	if(setjmp(jbuf)) {
			return 1;
	}
	longjmp(jbuf, 1);
	return 0;
}

__attribute__ ((noinline)) static int call_longjmp(int n) {
	if(n <= 0) {
		longjmp(jbuf, 1);
	}
	return call_longjmp(n-1);
}

static int GOjump_nest1() {
	if(setjmp(jbuf)) {
			return 1;
	}
	call_longjmp(10);
	return 0;
}

static int GOjump_struct(struct GObuf* b) {
	if(setjmp(b->jbuf)) {
			return 1;
	}
	longjmp(b->jbuf, 1);
	return 0;
}


