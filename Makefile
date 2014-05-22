CFLAGS += -std=c99 -pedantic-errors -Wall -Wextra -Werror 

default:
	$(CC) $(CFLAGS) $(LDFLAGS) rnbw.c -o rnbw 
