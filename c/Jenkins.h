#include <curl/curl.h>

// CVE-2018-1000861 Jenkins auth bypass
// check if Jenkins target is vulnerable
void checkJenkins(){
    CURL *curl = curl_easy_init();
    CURLcode responseCode;
    if(curl){
        // TODO: url handling / work endpoint into it like url + "securityRealm/user/admin/" + endpoint
        char *endpoint = "descriptorByName/org.jenkinsci.plugins.scriptsecurity.sandbox.groovy.SecureGroovyScript/checkScript";
        printf("Checking Jenkins endpoint: %s\n",endpoint);
        // get the linker working
        //  brew install curl
        //  export LDFLAGS="-L/opt/homebrew/opt/curl/lib"
        //  export CPPFLAGS="-I/opt/homebrew/opt/curl/include"
        curl_easy_setopt(curl, CURLOPT_URL, "http://127.0.0.1:10217");
        curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
        responseCode = curl_easy_perform(curl);
        printf("response code %u\n",responseCode);
        curl_easy_cleanup(curl);
        // TODO: if entry found then vulnerable
    }
}