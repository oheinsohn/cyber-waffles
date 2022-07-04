#include <stdio.h>

void displayScope()
{
    char host[10000] = "localhost";
    char IP[10000] = "127.0.0.1";

    printf("Host: %s\n", host);
    printf("IP: %s\n", IP);
}

int main(void)
{
    printf("### CyberWaffles ###\n");
    displayScope();
    return 0;
}

