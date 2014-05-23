#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <getopt.h>

#define PAINT "\033[38;5;%d;m%c\033[0m"

int rainbow[24] = {
154, 184, 214, 208,  209, 203, 204, 198, 199, 164, 129, 93, 99, 63, 69, 33, 39,
44, 49, 48, 84, 83, 119, 118
};

void go_bananas(FILE *fh) {
  if(fh == NULL) {
    return;
  }

  int c;
  int start = 0;
  int pos = 0;
  while( EOF != (c = fgetc(fh)) ) {
    if(c == 0x0a) {
      start++;
      pos = start;
      printf("%c", c);
      continue;
    }
    // TODO: strip out any existing colours first
    printf(PAINT, rainbow[pos], c);
    if(pos <23){
      pos++;
    } else {
      pos = 0;
    }
  }
}

int main(int argc, char *argv[]) {
  FILE *stream_input = NULL;
  int opt;
  if(argc <= 1) {
    stream_input = stdin;
  } else {
    while((opt = getopt(argc, argv, "f::h")) != -1) {
      switch(opt) {
        case 'f':
          stream_input = fopen(optarg, "r");
          break;
        case 'h':
          fprintf(stderr, "Usage: \n\tcat FILE | rnbw\n\trnbw -f FILE\n\n");
          exit(1);
        default:
          break;
      }
    }
  }

  go_bananas(stream_input);

  if(stream_input != stdin) {
    fclose(stream_input);
  }

  return 0;
}
