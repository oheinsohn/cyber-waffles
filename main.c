#include <stdio.h>

void readScopeFromJSON()
{
    FILE *scopeFile;
    scopeFile = fopen("scope.json","r");

    char scopeFileBuffer[10000];
    fread(scopeFileBuffer, 10000, 1, scopeFile);
	fclose(scopeFile);

    // No general parsing of this JSON - just for this case
    printf("%s\n", scopeFileBuffer);
}

void displayScope()
{
    char host[10000] = "localhost";
    char IP[10000] = "127.0.0.1";

    readScopeFromJSON();

    printf("Host: %s\n", host);
    printf("IP: %s\n", IP);
}

int main(void)
{
    printf("### CyberWaffles ###\n");
    displayScope();
    return 0;
}

