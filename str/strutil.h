// Func that reverse string presented as char*
// Estimate time: O(n) Estimate required memory: O(1)
void reverse(char* str){
    char* end = str;
    char tmp;
    if (str) {
    	while (*end) {
    		++end;
    	}
    	--end;
    	while (str < end){
    		tmp = *str;
    		*str++ = *end;
    		*end-- = tmp;
    	}
    }
}
