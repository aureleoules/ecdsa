# ECDSA in Go
This repository contains a set of tools to compute operations on an elliptic curve.  

## Elliptic Curve
An **elliptic curve** is defined by the equation:



![Elliptic Curve Equation](https://math.now.sh?from=y%5E2%20%3D%20x%5E3%2Bax%2Bb)  



It looks something like this: 
![secp256k1 curve](https://en.bitcoin.it/w/images/en/b/bf/Secp256k1.png)  
*secp256k1 curve used by Bitcoin*
![a](https://math.now.sh?from=a%3D0) 
![b](https://math.now.sh?from=b%3D7) 

## Point operations
[See ECDSA point operations](https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication#Point_operations).  

### Point at infinity
Point at infinity is defined by  
![O](https://math.now.sh?from=O%20%3D%20%5Cpmatrix%7B%200%20%5C%5C%200%7D)

* ![O + O = O](https://math.now.sh?from=O%20%2B%20O%20%3D%20O)  
* ![P + O = P](https://math.now.sh?from=P%20%2B%20O%20%3D%20P)  

### Point negation
* ![P - P = O](https://math.now.sh?from=P%20-%20P%20%3D%20O)
* ![P - Q = 0](https://math.now.sh?from=P%20-%20Q%20%3D%200%20%5CRightarrow%20P%20%3D%20Q)

### Point addition
![P + Q = R](https://math.now.sh?from=P%20%5Cne%20Q%20%5Cne%20O%20%5Cquad%20P%20%2B%20Q%20%3D%20R)
Where `R` is the negation of the intersection of the straight line defined by P and Q, and the curve.  

![Illustration](https://i.imgur.com/iFqWaS6.png)  

![lambda](https://math.now.sh?from=%5Clambda%20%3D%20%5Cfrac%7By_q-y_p%7D%7Bx_q-x_p%7D)  

![R](https://math.now.sh?from=R%20%3D%5Cpmatrix%7B%5Clambda%5E2-x_p%20-%20x_q%20%5C%5C%20%5Clambda(x_p-x_r)-y_q%7D)
 

### Point doubling
![P + P = R](https://math.now.sh?from=P%20%2B%20P%3D%20R)
Where `R` is the negation of the intersection of the curve's tangent at P, and the curve.  

![lambda](https://math.now.sh?from=%5Clambda%20%3D%20%5Cfrac%7B3x_p%5E2%20%2B%20a%7D%7B2y_p%7D)

![R](https://math.now.sh?from=R%20%3D%5Cpmatrix%7B%5Clambda%5E2-x_p%20-%20x_q%20%5C%5C%20%5Clambda(x_p-x_r)-y_q%7D)

### Point multiplication
Point multiplication is the repetition of point addition.
![Equation](https://math.now.sh?from=nP%20%3D%20%5Csum%5En_%7Bk%3D0%7D%7BP_k%7D%20%3D%20P%20%2B%20P%20%2B%20%5Cdots%20%2B%20P)  

**Montgomery ladder algorithm**  
```go
R0 ← 0
  R1 ← P
  for i from m downto 0 do
     if di = 0 then
        R1 ← point_add(R0, R1)
        R0 ← point_double(R0)
     else
        R0 ← point_add(R0, R1)
        R1 ← point_double(R1)
  return R0
```