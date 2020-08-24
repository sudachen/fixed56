
package fixed56
    
var divTestCases = []struct {
    x int
    y int
    z float64
    s string
}{
{
    x: 2,
    y:  3,
    z: 0.6666666666666666,
    s: "0+aaaaaaaaaaaaab/56",
},
{
    x: 1,
    y:  3,
    z: 0.3333333333333333,
    s: "0+55555555555555/56",
},
{
    x: 10,
    y:  7,
    z: 1.4285714285714286,
    s: "1+6db6db6db6db6e/56",
},
{
    x: 18,
    y:  7,
    z: 2.5714285714285716,
    s: "2+92492492492492/56",
},
}
