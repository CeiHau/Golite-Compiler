        .arch armv8-a
        .comm globalInit,8,8
        .comm Point,8,8
        .comm .nilPoint2D,8,8
        .text
        .type Init, %function
        .global Init
        .p2align 2
Init:
.L0:
        
// Prologue
        sub sp, sp, #160
        stp x29, x30, [sp]
        mov x29, sp
        
// store i64 %initVal, i64* %P_initVal 
        str x0, [sp, #152]
        
// %r0 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
        adrp x8, .nilPoint2D
        add x8, x8, :lo12:.nilPoint2D
        ldr x8, [x8]
        str x8, [sp, #32]
        
// store %struct.Point2D* %r0, %struct.Point2D** %newPt 
        ldr x8, [sp, #32]
        str x8, [sp, #24]
        
// br label %L1
        b .L1
.L1:
        
// %r1 = load i64, i64* %P_initVal
        ldr x8, [sp, #152]
        str x8, [sp, #40]
        
// %r1 = icmp sgt i64 %r1, 0
        ldr x8, [sp, #40]
        mov x9, #0
        cmp x9, x8
        
// br i1 %r2 label %L2, label %L3
        b.gt .L2
        b .L3
.L2:
        
// %r3 = call i8* @malloc(i32 16)
        mov x0, #16
        bl malloc
        str x0, [sp, #64]
        
// %r4 = bitcast i8* %r3 to %struct.Point2D*
        
// store %struct.Point2D* %r4, %struct.Point2D** %newPt 
        ldr x8, [sp, #64]
        str x8, [sp, #24]
        
// %r5 = load i64, i64* %P_initVal
        ldr x8, [sp, #152]
        str x8, [sp, #72]
        
// %r6 = load %struct.Point2D*, %struct.Point2D** %newPt
        ldr x8, [sp, #24]
        str x8, [sp, #80]
        
// store i64 %r5, i64* %r7 
        ldr x8, [sp, #72]
        str x8, [sp, #88]
        
// %r8 = load i64, i64* %P_initVal
        ldr x8, [sp, #152]
        str x8, [sp, #96]
        
// %r9 = load %struct.Point2D*, %struct.Point2D** %newPt
        ldr x8, [sp, #24]
        str x8, [sp, #104]
        
// store i64 %r8, i64* %r10 
        ldr x8, [sp, #96]
        str x8, [sp, #112]
        
// %r11 = load %struct.Point2D*, %struct.Point2D** %newPt
        ldr x8, [sp, #24]
        str x8, [sp, #120]
        
// store %struct.Point2D* %r11, %struct.Point2D** %_retval 
        ldr x8, [sp, #120]
        str x8, [sp, #16]
        
// %r12 = load %struct.Point2D*, %struct.Point2D** %_retval
        ldr x8, [sp, #16]
        str x8, [sp, #128]
        
// ret %struct.Point2D* %r12
        ldr x0, [sp, #128]
.L3:
        
// br label %L4
        b .L4
.L4:
        
// %r13 = load %struct.Point2D*, %struct.Point2D** %newPt
        ldr x8, [sp, #24]
        str x8, [sp, #136]
        
// store %struct.Point2D* %r13, %struct.Point2D** %_retval 
        ldr x8, [sp, #136]
        str x8, [sp, #16]
        
// %r14 = load %struct.Point2D*, %struct.Point2D** %_retval
        ldr x8, [sp, #16]
        str x8, [sp, #144]
        
// ret %struct.Point2D* %r14
        ldr x0, [sp, #144]
// Epilogue
        ldp x29, x30, [sp]
        add sp, sp, #160
        ret
        .size Init, (.-Init)
        .type main, %function
        .global main
        .p2align 2
main:
.L5:
        
// Prologue
        sub sp, sp, #272
        stp x29, x30, [sp]
        mov x29, sp
        
// store i64 5, i64* %a 
        mov x8, 5
        str x8, [sp, #24]
        
// %r15 = load i64, i64* %a
        ldr x8, [sp, #24]
        str x8, [sp, #56]
        
// %r16 = add i64 %r15, 7
        ldr x8, [sp, #56]
        mov x9, #7
        add x10, x8, x9
        str x10, [sp, #64]
        
// %r17 = mul i64 %r16, 3
        ldr x8, [sp, #64]
        mov x9, #3
        mul x10, x9, x8
        str x10, [sp, #72]
        
// store i64 %r17, i64* %b 
        ldr x8, [sp, #72]
        str x8, [sp, #32]
        
// %r18 = call i8* @malloc(i32 16)
        mov x0, #16
        bl malloc
        str x0, [sp, #88]
        
// %r19 = bitcast i8* %r18 to %struct.Point2D*
        
// store %struct.Point2D* %r19, %struct.Point2D** %pt1 
        ldr x8, [sp, #88]
        str x8, [sp, #40]
        
// %r20 = load i64, i64* %a
        ldr x8, [sp, #24]
        str x8, [sp, #96]
        
// %r21 = load %struct.Point2D*, %struct.Point2D** %pt1
        ldr x8, [sp, #40]
        str x8, [sp, #104]
        
// store i64 %r20, i64* %r22 
        ldr x8, [sp, #96]
        str x8, [sp, #112]
        
// %r23 = load i64, i64* %b
        ldr x8, [sp, #32]
        str x8, [sp, #120]
        
// %r24 = load %struct.Point2D*, %struct.Point2D** %pt1
        ldr x8, [sp, #40]
        str x8, [sp, #128]
        
// store i64 %r23, i64* %r25 
        ldr x8, [sp, #120]
        str x8, [sp, #136]
        
// store i64 1, i64* @globalInit 
        mov x8, 1
        adrp x9, globalInit
        add x9, x9, :lo12:globalInit
        str x8, [x9]
        
// %r26 = load i64, i64* @globalInit
        adrp x8, globalInit
        add x8, x8, :lo12:globalInit
        ldr x8, [x8]
        str x8, [sp, #144]
        
// %r27 = call %struct.Point2D* @Init(i64 %r26)
        ldr x0, [sp, #144]
        bl Init
        str x0, [sp, #152]
        
// %r28 = load i64, i64* @globalInit
        adrp x8, globalInit
        add x8, x8, :lo12:globalInit
        ldr x8, [x8]
        str x8, [sp, #160]
        
// %r29 = load %struct.Point2D*, %struct.Point2D** %pt2
        ldr x8, [sp, #48]
        str x8, [sp, #168]
        
// %r31 = load i64, i64* %r30
        ldr x8, [sp, #176]
        str x8, [sp, #184]
        
// %r32 = load %struct.Point2D*, %struct.Point2D** %pt2
        ldr x8, [sp, #48]
        str x8, [sp, #192]
        
// %r34 = load i64, i64* %r33
        ldr x8, [sp, #200]
        str x8, [sp, #208]
        adrp x8, .FMTSTR0
        add x8, x8, :lo12:.FMTSTR0
        mov x0, x8
        ldr x1, [sp, #160]
        ldr x2, [sp, #184]
        ldr x3, [sp, #208]
        bl printf
        
// %r36 = load %struct.Point2D*, %struct.Point2D** %pt1
        ldr x8, [sp, #40]
        str x8, [sp, #224]
        
// %r35 = load %struct.Point2D*, %struct.Point2D** %r36
        ldr x8, [sp, #224]
        str x8, [sp, #216]
        
// %r37 = bitcast %struct.Point2D* %r35 to i8*
        str x0, [sp, #232]
        bl free
        
// %r39 = load %struct.Point2D*, %struct.Point2D** %pt2
        ldr x8, [sp, #48]
        str x8, [sp, #248]
        
// %r38 = load %struct.Point2D*, %struct.Point2D** %r39
        ldr x8, [sp, #248]
        str x8, [sp, #240]
        
// %r40 = bitcast %struct.Point2D* %r38 to i8*
        str x0, [sp, #256]
        bl free
        
// br label %L6
        b .L6
.L6:
        
// store i64 0, i64* %_retval 
        mov x8, 0
        str x8, [sp, #16]
        
// %r41 = load i64, i64* %_retval
        ldr x8, [sp, #16]
        str x8, [sp, #264]
        
// ret i64 %r41
        ldr x0, [sp, #264]
// Epilogue
        ldp x29, x30, [sp]
        add sp, sp, #272
        ret
        .size main, (.-main)
.FMTSTR0:
        .asciz "offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00"
        .size .FMTSTR0, 32
