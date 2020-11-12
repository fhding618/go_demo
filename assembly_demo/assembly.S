"".g STEXT size=35 args=0x8 locals=0x0
	0x0000 00000 (assembly.go:3)	TEXT	"".g(SB), ABIInternal, $0-8
	0x0000 00000 (assembly.go:3)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:3)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:3)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:3)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:3)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:3)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:3)	JLS	28
	0x0012 00018 (assembly.go:3)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:3)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:3)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:4)	PCDATA	$0, $0
	0x0012 00018 (assembly.go:4)	PCDATA	$1, $0
	0x0012 00018 (assembly.go:4)	MOVL	"".p+4(SP), AX
	0x0016 00022 (assembly.go:4)	INCL	AX
	0x0017 00023 (assembly.go:4)	MOVL	AX, "".~r1+8(SP)
	0x001b 00027 (assembly.go:4)	RET
	0x001c 00028 (assembly.go:4)	NOP
	0x001c 00028 (assembly.go:3)	PCDATA	$1, $-1
	0x001c 00028 (assembly.go:3)	PCDATA	$0, $-2
	0x001c 00028 (assembly.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0021 00033 (assembly.go:3)	PCDATA	$0, $-1
	0x0021 00033 (assembly.go:3)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 0a 8b 44 24 04 40 89 44 24 08 c3 e8 00 00 00  v..D$.@.D$......
	0x0020 00 eb dd                                         ...
	rel 9+4 t=17 TLS+0
	rel 29+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=26 args=0x0 locals=0x0
	0x0000 00000 (assembly.go:7)	TEXT	"".main(SB), ABIInternal, $0-0
	0x0000 00000 (assembly.go:7)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:7)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:7)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:7)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:7)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:7)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:7)	JLS	19
	0x0012 00018 (assembly.go:7)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:7)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:7)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:7)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:10)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:10)	PCDATA	$1, $-1
	0x0012 00018 (assembly.go:10)	RET
	0x0013 00019 (assembly.go:10)	NOP
	0x0013 00019 (assembly.go:7)	PCDATA	$1, $-1
	0x0013 00019 (assembly.go:7)	PCDATA	$0, $-2
	0x0013 00019 (assembly.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0018 00024 (assembly.go:7)	PCDATA	$0, $-1
	0x0018 00024 (assembly.go:7)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 01 c3 e8 00 00 00 00 eb e6                    v.........
	rel 9+4 t=17 TLS+0
	rel 20+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".g$abstract SDWARFINFO dupok size=15
	0x0000 04 2e 67 00 01 01 11 70 00 00 00 00 00 00 00     ..g....p.......
	rel 10+4 t=29 go.info.int+0
go.loc."".g SDWARFLOC size=27
	0x0000 ff ff ff ff 00 00 00 00 00 00 00 00 23 00 00 00  ............#...
	0x0010 01 00 9c 00 00 00 00 00 00 00 00                 ...........
	rel 4+4 t=1 "".g+0
go.info."".g SDWARFINFO size=37
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 13  ................
	0x0010 00 00 00 00 00 00 00 00 0f 7e 72 31 00 01 03 00  .........~r1....
	0x0020 00 00 00 00 00                                   .....
	rel 0+0 t=24 type.int+0
	rel 1+4 t=29 go.info."".g$abstract+0
	rel 5+4 t=1 "".g+0
	rel 9+4 t=1 "".g+35
	rel 16+4 t=29 go.info."".g$abstract+6
	rel 20+4 t=29 go.loc."".g+0
	rel 31+4 t=29 go.info.int+0
go.range."".g SDWARFRANGE size=0
go.debuglines."".g SDWARFMISC size=13
	0x0000 04 02 11 c4 06 37 06 4a 04 01 03 7e 01           .....7.J...~.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=25
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 01 9c 00 00 00 00 01 00                       .........
	rel 9+4 t=1 "".main+0
	rel 13+4 t=1 "".main+26
	rel 19+4 t=30 gofile../Users/fhding/zhihu/code/go_mod/go_demo/assembly_demo/assembly.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=12
	0x0000 04 02 03 01 14 c6 16 04 01 03 7a 01              ..........z.
""..inittask SNOPTRDATA size=12
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00              ............
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".g STEXT size=35 args=0x8 locals=0x0
	0x0000 00000 (assembly.go:3)	TEXT	"".g(SB), ABIInternal, $0-8
	0x0000 00000 (assembly.go:3)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:3)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:3)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:3)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:3)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:3)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:3)	JLS	28
	0x0012 00018 (assembly.go:3)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:3)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:3)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:4)	PCDATA	$0, $0
	0x0012 00018 (assembly.go:4)	PCDATA	$1, $0
	0x0012 00018 (assembly.go:4)	MOVL	"".p+4(SP), AX
	0x0016 00022 (assembly.go:4)	INCL	AX
	0x0017 00023 (assembly.go:4)	MOVL	AX, "".~r1+8(SP)
	0x001b 00027 (assembly.go:4)	RET
	0x001c 00028 (assembly.go:4)	NOP
	0x001c 00028 (assembly.go:3)	PCDATA	$1, $-1
	0x001c 00028 (assembly.go:3)	PCDATA	$0, $-2
	0x001c 00028 (assembly.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0021 00033 (assembly.go:3)	PCDATA	$0, $-1
	0x0021 00033 (assembly.go:3)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 0a 8b 44 24 04 40 89 44 24 08 c3 e8 00 00 00  v..D$.@.D$......
	0x0020 00 eb dd                                         ...
	rel 9+4 t=17 TLS+0
	rel 29+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=26 args=0x0 locals=0x0
	0x0000 00000 (assembly.go:7)	TEXT	"".main(SB), ABIInternal, $0-0
	0x0000 00000 (assembly.go:7)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:7)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:7)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:7)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:7)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:7)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:7)	JLS	19
	0x0012 00018 (assembly.go:7)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:7)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:7)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:7)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:10)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:10)	PCDATA	$1, $-1
	0x0012 00018 (assembly.go:10)	RET
	0x0013 00019 (assembly.go:10)	NOP
	0x0013 00019 (assembly.go:7)	PCDATA	$1, $-1
	0x0013 00019 (assembly.go:7)	PCDATA	$0, $-2
	0x0013 00019 (assembly.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0018 00024 (assembly.go:7)	PCDATA	$0, $-1
	0x0018 00024 (assembly.go:7)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 01 c3 e8 00 00 00 00 eb e6                    v.........
	rel 9+4 t=17 TLS+0
	rel 20+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".g$abstract SDWARFINFO dupok size=15
	0x0000 04 2e 67 00 01 01 11 70 00 00 00 00 00 00 00     ..g....p.......
	rel 10+4 t=29 go.info.int+0
go.loc."".g SDWARFLOC size=27
	0x0000 ff ff ff ff 00 00 00 00 00 00 00 00 23 00 00 00  ............#...
	0x0010 01 00 9c 00 00 00 00 00 00 00 00                 ...........
	rel 4+4 t=1 "".g+0
go.info."".g SDWARFINFO size=37
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 13  ................
	0x0010 00 00 00 00 00 00 00 00 0f 7e 72 31 00 01 03 00  .........~r1....
	0x0020 00 00 00 00 00                                   .....
	rel 0+0 t=24 type.int+0
	rel 1+4 t=29 go.info."".g$abstract+0
	rel 5+4 t=1 "".g+0
	rel 9+4 t=1 "".g+35
	rel 16+4 t=29 go.info."".g$abstract+6
	rel 20+4 t=29 go.loc."".g+0
	rel 31+4 t=29 go.info.int+0
go.range."".g SDWARFRANGE size=0
go.debuglines."".g SDWARFMISC size=13
	0x0000 04 02 11 c4 06 37 06 4a 04 01 03 7e 01           .....7.J...~.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=25
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 01 9c 00 00 00 00 01 00                       .........
	rel 9+4 t=1 "".main+0
	rel 13+4 t=1 "".main+26
	rel 19+4 t=30 gofile../Users/fhding/zhihu/code/go_mod/go_demo/assembly_demo/assembly.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=12
	0x0000 04 02 03 01 14 c6 16 04 01 03 7a 01              ..........z.
""..inittask SNOPTRDATA size=12
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00              ............
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".g STEXT size=35 args=0x8 locals=0x0
	0x0000 00000 (assembly.go:8)	TEXT	"".g(SB), ABIInternal, $0-8
	0x0000 00000 (assembly.go:8)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:8)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:8)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:8)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:8)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:8)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:8)	JLS	28
	0x0012 00018 (assembly.go:8)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:8)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:8)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:8)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:9)	PCDATA	$0, $0
	0x0012 00018 (assembly.go:9)	PCDATA	$1, $0
	0x0012 00018 (assembly.go:9)	MOVL	"".p+4(SP), AX
	0x0016 00022 (assembly.go:9)	INCL	AX
	0x0017 00023 (assembly.go:9)	MOVL	AX, "".~r1+8(SP)
	0x001b 00027 (assembly.go:9)	RET
	0x001c 00028 (assembly.go:9)	NOP
	0x001c 00028 (assembly.go:8)	PCDATA	$1, $-1
	0x001c 00028 (assembly.go:8)	PCDATA	$0, $-2
	0x001c 00028 (assembly.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0021 00033 (assembly.go:8)	PCDATA	$0, $-1
	0x0021 00033 (assembly.go:8)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 0a 8b 44 24 04 40 89 44 24 08 c3 e8 00 00 00  v..D$.@.D$......
	0x0020 00 eb dd                                         ...
	rel 9+4 t=17 TLS+0
	rel 29+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=26 args=0x0 locals=0x0
	0x0000 00000 (assembly.go:12)	TEXT	"".main(SB), ABIInternal, $0-0
	0x0000 00000 (assembly.go:12)	MOVL	TLS, CX
	0x0007 00007 (assembly.go:12)	PCDATA	$0, $-2
	0x0007 00007 (assembly.go:12)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (assembly.go:12)	PCDATA	$0, $-1
	0x000d 00013 (assembly.go:12)	CMPL	SP, 8(CX)
	0x0010 00016 (assembly.go:12)	PCDATA	$0, $-2
	0x0010 00016 (assembly.go:12)	JLS	19
	0x0012 00018 (assembly.go:12)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:12)	PCDATA	$0, $-2
	0x0012 00018 (assembly.go:12)	PCDATA	$1, $-2
	0x0012 00018 (assembly.go:12)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:12)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:12)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (assembly.go:15)	PCDATA	$0, $-1
	0x0012 00018 (assembly.go:15)	PCDATA	$1, $-1
	0x0012 00018 (assembly.go:15)	RET
	0x0013 00019 (assembly.go:15)	NOP
	0x0013 00019 (assembly.go:12)	PCDATA	$1, $-1
	0x0013 00019 (assembly.go:12)	PCDATA	$0, $-2
	0x0013 00019 (assembly.go:12)	CALL	runtime.morestack_noctxt(SB)
	0x0018 00024 (assembly.go:12)	PCDATA	$0, $-1
	0x0018 00024 (assembly.go:12)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 01 c3 e8 00 00 00 00 eb e6                    v.........
	rel 9+4 t=17 TLS+0
	rel 20+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".g$abstract SDWARFINFO dupok size=15
	0x0000 04 2e 67 00 01 01 11 70 00 00 00 00 00 00 00     ..g....p.......
	rel 10+4 t=29 go.info.int+0
go.loc."".g SDWARFLOC size=27
	0x0000 ff ff ff ff 00 00 00 00 00 00 00 00 23 00 00 00  ............#...
	0x0010 01 00 9c 00 00 00 00 00 00 00 00                 ...........
	rel 4+4 t=1 "".g+0
go.info."".g SDWARFINFO size=37
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 13  ................
	0x0010 00 00 00 00 00 00 00 00 0f 7e 72 31 00 01 08 00  .........~r1....
	0x0020 00 00 00 00 00                                   .....
	rel 0+0 t=24 type.int+0
	rel 1+4 t=29 go.info."".g$abstract+0
	rel 5+4 t=1 "".g+0
	rel 9+4 t=1 "".g+35
	rel 16+4 t=29 go.info."".g$abstract+6
	rel 20+4 t=29 go.loc."".g+0
	rel 31+4 t=29 go.info.int+0
go.range."".g SDWARFRANGE size=0
go.debuglines."".g SDWARFMISC size=15
	0x0000 04 02 03 02 14 c4 06 37 06 4a 04 01 03 79 01     .......7.J...y.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=25
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 01 9c 00 00 00 00 01 00                       .........
	rel 9+4 t=1 "".main+0
	rel 13+4 t=1 "".main+26
	rel 19+4 t=30 gofile../Users/fhding/zhihu/code/go_mod/go_demo/assembly_demo/assembly.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=12
	0x0000 04 02 03 06 14 c6 16 04 01 03 75 01              ..........u.
""..inittask SNOPTRDATA size=12
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00              ............
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
