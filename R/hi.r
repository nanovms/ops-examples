say_hello_arg <- function(greet,  name) {
  args <- commandArgs(TRUE)
  name <- args[2]
  greet <- args[1]
  message  <- paste(greet,  name, sep = " ")
  print(message)
}
say_hello_arg(greet,  name)
