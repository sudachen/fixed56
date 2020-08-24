import os
import math
import scipy.stats
import random

prec = 56

def mul(a, b):
    z = int(a * (1 << prec)) * int(b * (1 << prec))
    return ((z >> (prec-1)) + 1) >> 1


def div(a, b):
    a = int(a * (1 << prec))
    b = int(b * (1 << prec))
    z = (a * (1 << prec) * 2) // b
    return (z + 1) >> 1


def fixed(x):
    return int(x * (1 << prec))


def float64(x):
    return x / (1 << prec)


def ceil(x):
    return math.ceil(x)


def floor(x):
    return math.floor(x)


def round(x):
    return float(int(x + (0.5 if x > -1 else -0.5)))


def string(x):
    # formats x as fixed with precision prec
    #maxval = (1 << (63-prec)) - 1
    v = abs(x)
    #if (v >> prec) >= maxval:
    #    return "overflow"
    fmt = "%%s%%d+%%0%dx/%d" % (prec / 4, prec)
    return fmt % (("", "-")[x < 0], v >> prec, v & ((1 << prec) - 1))


def xstring(a, b, v):
    #maxval = (1 << (63-prec)) - 1
    #if abs(a) >= maxval:
    #    return "overflow"
    #if abs(b) >= maxval:
    #    return "overflow"
    return string(v)


def gen_mul_tests():
    f = open('mul_ts_test.go1', 'w+')
    f.write('''
package fixed56
    
var mulTestCases = []struct {
    x float64
    y float64
    z float64
    s string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: {:3},
    s: "{:4}",
}},'''.format(a, b, float64(mul(a, b)), xstring(a, b, mul(a, b))))

    case(0, 1.5)
    case(1.25, 4)
    case(1.25, -4)
    case(-1.25, 4)
    case(-1.25, -4)
    case(1.25, 1.5)
    case(1234.5, -8888.875)
    case(1.515625, 1.531250)
    case(0.500244140625, 0.500732421875)
    case(0.015625, 0.000244140625)
    case(1.44140625, 1.44140625)
    case(1.44140625, 1.441650390625)
    f.write('\n}\n')
    f.close()
    if os.path.isfile('mul_ts_test.go'):
        os.remove('mul_ts_test.go')
    os.rename('mul_ts_test.go1', 'mul_ts_test.go')


def gen_div_tests():
    f = open('div_ts_test.go1', 'w+')
    f.write('''
package fixed56
    
var divTestCases = []struct {
    x int
    y int
    z float64
    s string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: {:3},
    s: "{:4}",
}},'''.format(int(a), int(b), a/b, xstring(a, b, div(a, b))))

    case(2, 3)
    case(1, 3)
    case(10, 7)
    case(18, 7)
    f.write('\n}\n')
    f.close()
    if os.path.isfile('div_ts_test.go'):
        os.remove('div_ts_test.go')
    os.rename('div_ts_test.go1', 'div_ts_test.go')


def gen_fixed_tests():
    f = open('fixed_ts_test.go1', 'w+')
    f.write('''
package fixed56
    
var testCases = []struct {
    x float64
    s string
    floor int
    round int
    ceil int
}{''')

    def case(a):
        a = float(a)
        f.write('''
{{
    x: {:1},
    s: "{:2}",
    floor: {:3},
    round: {:4},
    ceil:  {:5}, 
}},'''.format(a,string(fixed(a)), floor(a), round(a), ceil(a)))

    case(0)
    case(1)
    case(1.25)
    case(2.5)
    case(63/64)
    case(-0.5)
    case(-4.125)
    case(-7.75)

    f.write('\n}\n')
    f.close()
    if os.path.isfile('fixed_ts_test.go'):
        os.remove('fixed_ts_test.go')
    os.rename('fixed_ts_test.go1', 'fixed_ts_test.go')


def gen_bincdf_tests():

    f = open('bincdf_ts_test.go1', 'w+')
    f.write('''
package fixed56

var bincdfTestCases = []struct {
    n,x int64
    p,cdf float64
    s string
}{''')

    def case(n,x,p):
        n = int(n)
        x = int(x)
        p = float(p)

        cdf = scipy.stats.binom.cdf(x, n, p)
        if abs(cdf) < 1e-16:
            cdf = 0
        #print(cdf, x,n,p)
        if math.isnan(cdf):
            return

        f.write('''
{{
    n: {:1},
    x: {:2},
    p: {:3},
    cdf: {:4},
    s: "{:5}",
}},'''.format(n,x,p,cdf,string(fixed(cdf))))

    random.seed(42)

    def q(n):
        x = n//10 + random.randint(0,n//10)
        case(n, x, 0.2 + random.random()*0.2 - 0.1)
        x = n//2 + random.randint(0,n//2-1)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n//3+ random.randint(0,n//3)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n//3*2 + random.randint(0,n//3-1)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n-n//10 - random.randint(0,n//10)
        case(n, x, 0.8 + random.random()*0.2 - 0.1)

    #N = 1 << 32
    #q(N)
    for k in range(50):
        n = random.randint(30,1000)
        q(n)
    for k in range(50):
        n = random.randint(300,10000)
        q(n)
    for k in range(50):
        n = random.randint(3000,100000)
        q(n)
    for k in range(50):
        n = random.randint(30000,1000000)
        q(n)
    for k in range(50):
        n = random.randint(300,10000000)
        q(n)

    f.write('\n}\n')
    f.close()
    if os.path.isfile('bincdf_ts_test.go'):
        os.remove('bincdf_ts_test.go')
    os.rename('bincdf_ts_test.go1', 'bincdf_ts_test.go')


if __name__ == '__main__':
    gen_mul_tests()
    gen_div_tests()
    gen_fixed_tests()
    gen_bincdf_tests()
