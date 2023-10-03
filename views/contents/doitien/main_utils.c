#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
void DoiTien(int *arr ,int tien_can_doi) {
    int tien = tien_can_doi;

     while (tien > 0)
    {
        if (tien >= 500)
        {
            arr[0] = tien/500;
            tien %= 500;
        } else if (tien >= 200)
        {
            arr[1] = tien/200;
            tien %= 200;
        } else if (tien >= 100)
        {
            arr[2] = tien/100;
            tien %= 100;
        } else if (tien >= 50)
        {
            arr[3] = tien/50;
            tien %= 50;
        } else if (tien >= 20)
        {
            arr[4] = tien/20;
            tien %= 20;
        } else if (tien >= 10)
        {
            arr[5] = tien/10;
            tien %= 10;
        } else if (tien >= 5)
        {
            arr[6] = tien/5;
            tien %= 5;
        } else if (tien >= 2)
        {
            arr[7] = tien/2;
            tien %= 2;
        }else
        {
            tien -= 1;
            arr[8]=0;
            arr[8]++;
        }
    }
}