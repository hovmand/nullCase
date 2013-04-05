#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv) {
	FILE *file;
	char temp[512];
	
	printf("HEJ\n");
	
	file = popen("php -f ../package/php/test.php", "r");
	
	if (file == NULL) {
		printf("Can't open file\n");
		return 1;
	}
	
	while (fgets(temp, sizeof(temp), file) != NULL) {
		printf("!\n");
		printf("%s\n", temp);
	}
	
	fclose(file);
	
	printf("HEJ\n");
	
	return 0;
}
