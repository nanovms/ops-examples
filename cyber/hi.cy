import m 'math'

worlds = ['World', '世界', 'दुनिया']
worlds.append(m.random())
for worlds each w:
    print 'Hello, {w}!'

func fib(n int) int:
    coyield
    if n < 2:
        return n
    return fib(n - 1) + fib(n - 2)

count = 0    -- Counts iterations.
fiber = coinit fib(30)
while fiber.status() != #done:
    res = coresume fiber
    count += 1
print '{res} {count}'
