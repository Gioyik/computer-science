#include <stdio.h>

int main() {
  int x = 5;
  int z[7] = {1,2,3,4,5,6,7};
  int i;

  for (i=0; i<x; i++) {
    if (x == z[i]) {
      printf("%d\n", i);
    }
  }
}
