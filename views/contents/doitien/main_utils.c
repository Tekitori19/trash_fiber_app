#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
int UCLN(int x, int y) {
    int temp = 0;
    int min = MIN(x, y);
    int max = MAX(x, y);
    
    for (int i = 1; i < max; i++)
    {
        if (min % i == 0 && max % i ==0)
        {
            temp = i;
        }   
    }
    return temp;
}

EMSCRIPTEN_KEEPALIVE
int BCNN(int x, int y, int ucln) {
    return (x*y)/ucln;
}