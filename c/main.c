#include <stdio.h>
#include <CL/cl.h>

int main() {
  cl_int CL_err = CL_SUCCESS;
  cl_uint numPlatforms = 0;

  CL_err = clGetPlatformIDs(0, NULL, &numPlatforms);

  if (CL_err == CL_SUCCESS) {
    printf("%u platform(s) found\n", numPlatforms);
  } else {
    printf("clGetPlatformIDs(%i)\n", CL_err);
  }

  return 0;
}
