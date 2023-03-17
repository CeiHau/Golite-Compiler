source_filename = "simple1"
target triple = "x86_64-linux-gnu"

define i64 @fact(i64% x)
{
L0:
    %_retval = alloca i64
    %P_x = alloca i64
    store i64 %x, i64* %P_x 
    br label %L1
L1:
    %r0 = load i64, i64* %P_x
    %r0 = icmp sle i64 %r0, 1
    br i1 %r1 label %L2, label %L3
L2:
    store i64 1, i64* %_retval 
    %r2 = load i64, i64* %_retval
    ret i64 %r2
L3:
    %r3 = load i64, i64* %P_x
    %r4 = load i64, i64* %P_x
    %r5 = sub i64 %r4, 1
    %r6 = call i64 @fact(i64 %r5)
    %r7 = mul i64 %r3, %r6
    store i64 %r7, i64* %_retval 
    %r8 = load i64, i64* %_retval
    ret i64 %r8
L4:
    store i64 0, i64* %_retval 
    %r9 = load i64, i64* %_retval
    ret i64 %r9
}

define i64 @main()
{
L5:
    %_retval = alloca i64
    %stop = alloca i64
    %factor = alloca i64
    %toStop = alloca i64
    %temp = alloca i64
    store i64 0, i64* %stop 
    store i64 0, i64* %factor 
    call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @factor)
    %r10 = load i64, i64* %factor
    %r11 = call i64 @fact(i64 %r10)
    store i64 %r11, i64* %temp 
    %r12 = load i64, i64* %temp
    call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.fSTR0, i32 0, i32 0), i64 %r12)
    br label %L6
L6:
    store i64 0, i64* %_retval 
    %r13 = load i64, i64* %_retval
    ret i64 %r13
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
@.read = private unnamed_addr constant [4 x i8] c%l\00", align 1
@.fSTR0 = private unnamed_addr constant [3 x i8] c"%ld\00", align 1
