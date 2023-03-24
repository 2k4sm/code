#include <stdio.h>


int main(){
    int number;
    printf("Enter your Number: ");
    scanf("%i",&number);
    if (number > 0){
        printf("%i\n",number);
    }else{
        printf("Invalid Number..\n");
    }
}