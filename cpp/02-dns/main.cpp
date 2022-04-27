#include <sys/types.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>

int main(int argc, char *argv[])
{

  struct hostent *hp;
  struct in_addr ip_addr;

  const  char *host = "dns.google";

  hp = gethostbyname(host);
  if (!hp) {
    printf("was not resolved\n");
  }

  ip_addr = *(struct in_addr *)(hp->h_addr);
  printf("Hostname: dns.google, was resolved to: %s\n",
  inet_ntoa(ip_addr));

  exit(EXIT_SUCCESS);

}
