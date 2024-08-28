@7  // A <- 7
D=A // D <- 7
@SP // M <- &SP
A=M // A <- M: アドレス値を覚える
M=D // &SP <- 7
@SP // M <- &SP
M=M+1 // SP++

@8
D=A
@SP
A=M
M=D
@SP
M=M+1

// pop to M to D
@SP   // M <- &SP
M=M-1 // SP--
A=M   // A <- SP

D=M   // D  <- M[SP]

@SP
M=M-1
A=M

D=D+M

// push from D to SP
@SP   // M  <- &SP
A=M   // A  <- SP
M=D   // &SP<- D: result (15)
@SP   // M  <- &SP
M=M+1 // SP <- SP+1
