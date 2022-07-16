#include <stdio.h>
#include "JSONParser.h"
#include "Jenkins.h"

void readScopeFromJSON()
{
    FILE *scopeFile;
    scopeFile = fopen("./resources/scope.json","r");

    char scopeFileBuffer[1000];
    fread(scopeFileBuffer, 10000, 1, scopeFile);
	fclose(scopeFile);

    parseJSON(scopeFileBuffer);
}

void displayScope()
{
    readScopeFromJSON();
}

int main(void)
{
    printf("### CyberWaffles ###\n");
    displayScope();
    checkJenkins();
    return 0;
}

