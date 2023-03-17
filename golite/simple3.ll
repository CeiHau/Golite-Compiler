source_filename = "simple3"
target triple = "x86_64-linux-gnu"
%struct.Point2D = type{i64, i64}
@.nilPoint2D = common global %struct.Point2D* null
@globalInit = common global i64 0
@Point = common global %struct.Point2D* null

define %struct.Point2D* @Init(i64% initVal)
{
L0:
    %_retval = alloca %struct.Point2D*
    %P_initVal = alloca i64
    %newPt = alloca %struct.Point2D*
    store i64 %initVal, i64* %P_initVal 
    %r0 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
    store %struct.Point2D* %r0, %struct.Point2D** %newPt 
    br label %L1
L1:
    %r1 = load i64, i64* %P_initVal
    %r1 = icmp sgt i64 %r1, 0
    br i1 %r2 label %L2, label %L3
L2:
    %r3 = call i8* @malloc(i32 16)
    %r4 = bitcast i8* %r3 to %struct.Point2D*
    store %struct.Point2D* %r4, %struct.Point2D** %newPt 
    %r5 = load i64, i64* %P_initVal
    %r6 = load %struct.Point2D*, %struct.Point2D** %newPt
    %r7 = getelementptr %struct.Point2D, %struct.Point2D* %r6, i32 0, i32 0
    store i64 %r5, i64* %r7 
    %r8 = load i64, i64* %P_initVal
    %r9 = load %struct.Point2D*, %struct.Point2D** %newPt
    %r10 = getelementptr %struct.Point2D, %struct.Point2D* %r9, i32 0, i32 1
    store i64 %r8, i64* %r10 
    %r11 = load %struct.Point2D*, %struct.Point2D** %newPt
    store %struct.Point2D* %r11, %struct.Point2D** %_retval 
    %r12 = load %struct.Point2D*, %struct.Point2D** %_retval
    ret %struct.Point2D* %r12
L3:
    br label %L4
L4:
    %r13 = load %struct.Point2D*, %struct.Point2D** %newPt
    store %struct.Point2D* %r13, %struct.Point2D** %_retval 
    %r14 = load %struct.Point2D*, %struct.Point2D** %_retval
    ret %struct.Point2D* %r14
}

define i64 @main()
{
L5:
    %_retval = alloca i64
    %a = alloca i64
    %b = alloca i64
    %pt1 = alloca %struct.Point2D*
    %pt2 = alloca %struct.Point2D*
    store i64 5, i64* %a 
    %r15 = load i64, i64* %a
    %r16 = add i64 %r15, 7
    %r17 = mul i64 %r16, 3
    store i64 %r17, i64* %b 
    %r18 = call i8* @malloc(i32 16)
    %r19 = bitcast i8* %r18 to %struct.Point2D*
    store %struct.Point2D* %r19, %struct.Point2D** %pt1 
    %r20 = load i64, i64* %a
    %r21 = load %struct.Point2D*, %struct.Point2D** %pt1
    %r22 = getelementptr %struct.Point2D, %struct.Point2D* %r21, i32 0, i32 0
    store i64 %r20, i64* %r22 
    %r23 = load i64, i64* %b
    %r24 = load %struct.Point2D*, %struct.Point2D** %pt1
    %r25 = getelementptr %struct.Point2D, %struct.Point2D* %r24, i32 0, i32 1
    store i64 %r23, i64* %r25 
    store i64 1, i64* @globalInit 
    %r26 = load i64, i64* @globalInit
    %r27 = call %struct.Point2D* @Init(i64 %r26)
    %r28 = load i64, i64* @globalInit
    %r29 = load %struct.Point2D*, %struct.Point2D** %pt2
    %r30 = getelementptr %struct.Point2D, %struct.Point2D* %r29, i32 0, i32 0
    %r31 = load i64, i64* %r30
    %r32 = load %struct.Point2D*, %struct.Point2D** %pt2
    %r33 = getelementptr %struct.Point2D, %struct.Point2D* %r32, i32 0, i32 1
    %r34 = load i64, i64* %r33
    call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fSTR0, i32 0, i32 0), i64 %r28, i64 %r31, i64 %r34)
    %r36 = load %struct.Point2D*, %struct.Point2D** %pt1
    %r35 = load %struct.Point2D*, %struct.Point2D** %r36
    %r37 = bitcast %struct.Point2D* %r35 to i8*
    call void @free(i8* %r37)
    %r39 = load %struct.Point2D*, %struct.Point2D** %pt2
    %r38 = load %struct.Point2D*, %struct.Point2D** %r39
    %r40 = bitcast %struct.Point2D* %r38 to i8*
    call void @free(i8* %r40)
    br label %L6
L6:
    store i64 0, i64* %_retval 
    %r41 = load i64, i64* %_retval
    ret i64 %r41
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
@.fSTR0 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1
