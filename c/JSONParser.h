#include <string.h>

// No general parsing of this JSON - just for this context

struct parsedScopeEntry{
  char *hostname;
  char *IP;
  char *port;
};    

struct parsedScopeEntry parseEntries(char *entry){
    char *entryItem;
    struct parsedScopeEntry parsedScopeEntryItem;
    entryItem = strtok(entry, "\"");
    entryItem = strtok(NULL, "\"");
    if(strcmp(entryItem,"name") == 0){
        entryItem = strtok(NULL, "\"");
        entryItem = strtok(NULL, "\"");
        parsedScopeEntryItem.hostname = entryItem;
    } else {
        printf("Item invalid - hostname\n");
    }
    entryItem = strtok(NULL, "\"");
    entryItem = strtok(NULL, "\"");
    if(strcmp(entryItem,"IP") == 0){
        entryItem = strtok(NULL, "\"");
        entryItem = strtok(NULL, "\"");
        parsedScopeEntryItem.IP = entryItem;
    } else {
        printf("Item invalid - IP\n");
    }
    entryItem = strtok(NULL, "\"");
    entryItem = strtok(NULL, "\"");
    if(strcmp(entryItem,"Port") == 0){
        entryItem = strtok(NULL, "\"");
        entryItem = strtok(NULL, "\"");
        parsedScopeEntryItem.port = entryItem;
    } else {
        printf("Item invalid - port\n");
    }
    return parsedScopeEntryItem;
}

void parseJSON(char *scopeFileBuffer){
    char *scopeEntry;
    char *arrayEntries[5];
    struct parsedScopeEntry parsedScopeEntryArray[5];
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
            parsedScopeEntryArray[displayArrayCounter] = parseEntries(arrayEntries[displayArrayCounter]);
    }
    for(int displayArrayCounter = 0; 
        displayArrayCounter < entryCounter; 
        ++displayArrayCounter
        ) {
            printf("---%d %s %s %s\n", 
                displayArrayCounter,
                parsedScopeEntryArray[displayArrayCounter].hostname,
                parsedScopeEntryArray[displayArrayCounter].IP,
                parsedScopeEntryArray[displayArrayCounter].port
                );
    }
}

