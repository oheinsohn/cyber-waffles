#include <string.h>

// No general parsing of this JSON - just for this context    
void parseJSON(char *scopeFileBuffer){
    
    // show all of the JSON
    // printf("%s\n", scopeFileBuffer);
   
    char *scopeEntry;
    char *arrayEntries[3];
    // first entry is just "scope" - no real entry
    scopeEntry = strtok(scopeFileBuffer, "{");
    scopeEntry = strtok(NULL, "{");
    int entryCounter = 0;
    while(scopeEntry != NULL) {
        arrayEntries[entryCounter++] = scopeEntry;    
        scopeEntry = strtok(NULL, "{");
    }
    for(int displayArrayCounter = 0; 
        displayArrayCounter < entryCounter; 
        ++displayArrayCounter
        ) {
            printf("%s\n", arrayEntries[displayArrayCounter]);
    }
}