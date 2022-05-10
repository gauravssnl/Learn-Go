#include <stdio.h>
#include "callC.h"

void cHello() {
    printf("Hello from C!\n");
}

void printMessage(char *message) {
    printf("Go sent me %s\n", message);
}