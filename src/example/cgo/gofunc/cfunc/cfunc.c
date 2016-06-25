#include <stdio.h>
#include "cfunc.h"
#include "_cgo_export.h"

void
Say(char *name, char *msg){
	printf("[C]%s: %s\n", name, msg);
	Show(name,msg);
}
