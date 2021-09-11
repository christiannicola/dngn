package graphics

type CharCode int

const (
	// 1 - 15.
	WhiteSmilingFace   CharCode = 0x263a
	BlackSmilingFace   CharCode = 0x263b
	BlackHeartSuit     CharCode = 0x2665
	BlackDiamondSuit   CharCode = 0x2666
	BlackClubSuit      CharCode = 0x2663
	BlackSpadeSuit     CharCode = 0x2660
	Bullet             CharCode = 0x2022
	InverseBullet      CharCode = 0x25d8
	WhiteCircle        CharCode = 0x25cb
	InverseWhiteCircle CharCode = 0x25d9
	MaleSign           CharCode = 0x2642
	FemaleSign         CharCode = 0x2640
	EighthNote         CharCode = 0x266a
	BeamedEighthNotes  CharCode = 0x266b
	WhiteSunWithRays   CharCode = 0x263c

	// 16 - 31.
	BlackRightPointingPointer CharCode = 0x25ba
	BlackLeftPointingPointer  CharCode = 0x25c4
	UpDownArrow               CharCode = 0x2195
	DoubleExclamationMark     CharCode = 0x203c
	Pilcrow                   CharCode = 0x00b6
	SectionSign               CharCode = 0x00a7
	BlackRectangle            CharCode = 0x25ac
	UpDownArrowWithBase       CharCode = 0x21a8
	UpwardsArrow              CharCode = 0x2191
	DownwardsArrow            CharCode = 0x2193
	RightwardsArrow           CharCode = 0x2192
	LeftwardsArrow            CharCode = 0x2190
	RightAngle                CharCode = 0x221f
	LeftRightArrow            CharCode = 0x2194
	BlackUpPointingTriangle   CharCode = 0x25b2
	BlackDownPointingTriangle CharCode = 0x25bc

	// 32 - 47.
	Space            CharCode = 0x0020
	ExclamationPoint CharCode = 0x0021
	DoubleQuote      CharCode = 0x0022
	NumberSign       CharCode = 0x0023
	DollarSign       CharCode = 0x0024
	Percent          CharCode = 0x0025
	Ampersand        CharCode = 0x0026
	Apostrophe       CharCode = 0x0027
	LeftParenthesis  CharCode = 0x0028
	RightParenthesis CharCode = 0x0029
	Asterisk         CharCode = 0x002a
	Plus             CharCode = 0x002b
	Comma            CharCode = 0x002c
	Minus            CharCode = 0x002d
	Period           CharCode = 0x002e
	Slash            CharCode = 0x002f

	// 48 - 63.
	Zero         CharCode = 0x0030
	One          CharCode = 0x0031
	Two          CharCode = 0x0032
	Three        CharCode = 0x0033
	Four         CharCode = 0x0034
	Five         CharCode = 0x0035
	Six          CharCode = 0x0036
	Seven        CharCode = 0x0037
	Eight        CharCode = 0x0038
	Nine         CharCode = 0x0039
	Colon        CharCode = 0x003a
	Semicolon    CharCode = 0x003b
	LessThan     CharCode = 0x003c
	Equals       CharCode = 0x003d
	GreaterThan  CharCode = 0x003e
	QuestionMark CharCode = 0x003f

	// 64 - 95.
	At           CharCode = 0x0040
	AUpper       CharCode = 0x0041
	BUpper       CharCode = 0x0042
	CUpper       CharCode = 0x0043
	DUpper       CharCode = 0x0044
	EUpper       CharCode = 0x0045
	FUpper       CharCode = 0x0046
	GUpper       CharCode = 0x0047
	HUpper       CharCode = 0x0048
	IUpper       CharCode = 0x0049
	JUpper       CharCode = 0x004a
	KUpper       CharCode = 0x004b
	LUpper       CharCode = 0x004c
	MUpper       CharCode = 0x004d
	NUpper       CharCode = 0x004e
	OUpper       CharCode = 0x004f
	PUpper       CharCode = 0x0050
	QUpper       CharCode = 0x0051
	RUpper       CharCode = 0x0052
	SUpper       CharCode = 0x0053
	TUpper       CharCode = 0x0054
	UUpper       CharCode = 0x0055
	VUpper       CharCode = 0x0056
	WUpper       CharCode = 0x0057
	XUpper       CharCode = 0x0058
	YUpper       CharCode = 0x0059
	ZUpper       CharCode = 0x005a
	LeftBracket  CharCode = 0x005b
	BackSlash    CharCode = 0x005c
	RightBracket CharCode = 0x005d
	Caret        CharCode = 0x005e
	Underscore   CharCode = 0x005f

	// 96 - 127.
	Accent     CharCode = 0x0060
	ALower     CharCode = 0x0061
	BLower     CharCode = 0x0062
	CLower     CharCode = 0x0063
	DLower     CharCode = 0x0064
	ELower     CharCode = 0x0065
	FLower     CharCode = 0x0066
	GLower     CharCode = 0x0067
	HLower     CharCode = 0x0068
	ILower     CharCode = 0x0069
	JLower     CharCode = 0x006a
	KLower     CharCode = 0x006b
	LLower     CharCode = 0x006c
	MLower     CharCode = 0x006d
	NLower     CharCode = 0x006e
	OLower     CharCode = 0x006f
	PLower     CharCode = 0x0070
	QLower     CharCode = 0x0071
	RLower     CharCode = 0x0072
	SLower     CharCode = 0x0073
	TLower     CharCode = 0x0074
	ULower     CharCode = 0x0075
	VLower     CharCode = 0x0076
	WLower     CharCode = 0x0077
	XLower     CharCode = 0x0078
	YLower     CharCode = 0x0079
	ZLower     CharCode = 0x007a
	LeftBrace  CharCode = 0x007b
	Pipe       CharCode = 0x007c
	RightBrace CharCode = 0x007d
	Tilde      CharCode = 0x007e
	House      CharCode = 0x2302

	// 128 - 143.
	LatinCapitalLetterCWithCedilla   CharCode = 0x00c7
	LatinSmallLetterUWithDiaeresis   CharCode = 0x00fc
	LatinSmallLetterEWithAcute       CharCode = 0x00e9
	LatinSmallLetterAWithCircumflex  CharCode = 0x00e2
	LatinSmallLetterAWithDiaeresis   CharCode = 0x00e4
	LatinSmallLetterAWithGrave       CharCode = 0x00e0
	LatinSmallLetterAWithRingAbove   CharCode = 0x00e5
	LatinSmallLetterCWithCedilla     CharCode = 0x00e7
	LatinSmallLetterEWithCircumflex  CharCode = 0x00ea
	LatinSmallLetterEWithDiaeresis   CharCode = 0x00eb
	LatinSmallLetterEWithGrave       CharCode = 0x00e8
	LatinSmallLetterIWithDiaeresis   CharCode = 0x00ef
	LatinSmallLetterIWithCircumflex  CharCode = 0x00ee
	LatinSmallLetterIWithGrave       CharCode = 0x00ec
	LatinCapitalLetterAWithDiaeresis CharCode = 0x00c4
	LatinCapitalLetterAWithRingAbove CharCode = 0x00c5

	// 144 - 159.
	LatinCapitalLetterEWithAcute     CharCode = 0x00c9
	LatinSmallLetterAe               CharCode = 0x00e6
	LatinCapitalLetterAe             CharCode = 0x00c6
	LatinSmallLetterOWithCircumflex  CharCode = 0x00f4
	LatinSmallLetterOWithDiaeresis   CharCode = 0x00f6
	LatinSmallLetterOWithGrave       CharCode = 0x00f2
	LatinSmallLetterUWithCircumflex  CharCode = 0x00fb
	LatinSmallLetterUWithGrave       CharCode = 0x00f9
	LatinSmallLetterYWithDiaeresis   CharCode = 0x00ff
	LatinCapitalLetterOWithDiaeresis CharCode = 0x00d6
	LatinCapitalLetterUWithDiaeresis CharCode = 0x00dc
	CentSign                         CharCode = 0x00a2
	PoundSign                        CharCode = 0x00a3
	YenSign                          CharCode = 0x00a5
	PesetaSign                       CharCode = 0x20a7
	LatinSmallLetterFWithHook        CharCode = 0x0192

	// 160 - 175.
	LatinSmallLetterAWithAcute            CharCode = 0x00e1
	LatinSmallLetterIWithAcute            CharCode = 0x00ed
	LatinSmallLetterOWithAcute            CharCode = 0x00f3
	LatinSmallLetterUWithAcute            CharCode = 0x00fa
	LatinSmallLetterNWithTilde            CharCode = 0x00f1
	LatinCapitalLetterNWithTilde          CharCode = 0x00d1
	FeminineOrdinalIndicator              CharCode = 0x00aa
	MasculineOrdinalIndicator             CharCode = 0x00ba
	InvertedQuestionMark                  CharCode = 0x00bf
	ReversedNotSign                       CharCode = 0x2310
	NotSign                               CharCode = 0x00ac
	VulgarFractionOneHalf                 CharCode = 0x00bd
	VulgarFractionOneQuarter              CharCode = 0x00bc
	InvertedExclamationMark               CharCode = 0x00a1
	LeftPointingDoubleAngleQuotationMark  CharCode = 0x00ab
	RightPointingDoubleAngleQuotationMark CharCode = 0x00bb

	// 176 - 191.
	LightShade                             CharCode = 0x2591
	MediumShade                            CharCode = 0x2592
	DarkShade                              CharCode = 0x2593
	BoxDrawingsLightVertical               CharCode = 0x2502
	BoxDrawingsLightVerticalAndLeft        CharCode = 0x2524
	BoxDrawingsVerticalSingleAndLeftDouble CharCode = 0x2561
	BoxDrawingsVerticalDoubleAndLeftSingle CharCode = 0x2562
	BoxDrawingsDownDoubleAndLeftSingle     CharCode = 0x2556
	BoxDrawingsDownSingleAndLeftDouble     CharCode = 0x2555
	BoxDrawingsDoubleVerticalAndLeft       CharCode = 0x2563
	BoxDrawingsDoubleVertical              CharCode = 0x2551
	BoxDrawingsDoubleDownAndLeft           CharCode = 0x2557
	BoxDrawingsDoubleUpAndLeft             CharCode = 0x255d
	BoxDrawingsUpDoubleAndLeftSingle       CharCode = 0x255c
	BoxDrawingsUpSingleAndLeftDouble       CharCode = 0x255b
	BoxDrawingsLightDownAndLeft            CharCode = 0x2510

	// 192 - 207.
	BoxDrawingsLightUpAndRight              CharCode = 0x2514
	BoxDrawingsLightUpAndHorizontal         CharCode = 0x2534
	BoxDrawingsLightDownAndHorizontal       CharCode = 0x252c
	BoxDrawingsLightVerticalAndRight        CharCode = 0x251c
	BoxDrawingsLightHorizontal              CharCode = 0x2500
	BoxDrawingsLightVerticalAndHorizontal   CharCode = 0x253c
	BoxDrawingsVerticalSingleAndRightDouble CharCode = 0x255e
	BoxDrawingsVerticalDoubleAndRightSingle CharCode = 0x255f
	BoxDrawingsDoubleUpAndRight             CharCode = 0x255a
	BoxDrawingsDoubleDownAndRight           CharCode = 0x2554
	BoxDrawingsDoubleUpAndHorizontal        CharCode = 0x2569
	BoxDrawingsDoubleDownAndHorizontal      CharCode = 0x2566
	BoxDrawingsDoubleVerticalAndRight       CharCode = 0x2560
	BoxDrawingsDoubleHorizontal             CharCode = 0x2550
	BoxDrawingsDoubleVerticalAndHorizontal  CharCode = 0x256c
	BoxDrawingsUpSingleAndHorizontalDouble  CharCode = 0x2567

	// 208 - 223.
	BoxDrawingsUpDoubleAndHorizontalSingle       CharCode = 0x2568
	BoxDrawingsDownSingleAndHorizontalDouble     CharCode = 0x2564
	BoxDrawingsDownDoubleAndHorizontalSingle     CharCode = 0x2565
	BoxDrawingsUpDoubleAndRightSingle            CharCode = 0x2559
	BoxDrawingsUpSingleAndRightDouble            CharCode = 0x2558
	BoxDrawingsDownSingleAndRightDouble          CharCode = 0x2552
	BoxDrawingsDownDoubleAndRightSingle          CharCode = 0x2553
	BoxDrawingsVerticalDoubleAndHorizontalSingle CharCode = 0x256b
	BoxDrawingsVerticalSingleAndHorizontalDouble CharCode = 0x256a
	BoxDrawingsLightUpAndLeft                    CharCode = 0x2518
	BoxDrawingsLightDownAndRight                 CharCode = 0x250c
	FullBlock                                    CharCode = 0x2588
	LowerHalfBlock                               CharCode = 0x2584
	LeftHalfBlock                                CharCode = 0x258c
	RightHalfBlock                               CharCode = 0x2590
	UpperHalfBlock                               CharCode = 0x2580

	// 224 - 239.
	GreekSmallLetterAlpha   CharCode = 0x03b1
	LatinSmallLetterSharpS  CharCode = 0x00df
	GreekCapitalLetterGamma CharCode = 0x0393
	GreekSmallLetterPi      CharCode = 0x03c0
	GreekCapitalLetterSigma CharCode = 0x03a3
	GreekSmallLetterSigma   CharCode = 0x03c3
	MicroSign               CharCode = 0x00b5
	GreekSmallLetterTau     CharCode = 0x03c4
	GreekCapitalLetterPhi   CharCode = 0x03a6
	GreekCapitalLetterTheta CharCode = 0x0398
	GreekCapitalLetterOmega CharCode = 0x03a9
	GreekSmallLetterDelta   CharCode = 0x03b4
	Infinity                CharCode = 0x221e
	GreekSmallLetterPhi     CharCode = 0x03c6
	GreekSmallLetterEpsilon CharCode = 0x03b5
	Intersection            CharCode = 0x2229

	// 240 - 254.
	IdenticalTo                  CharCode = 0x2261
	PlusMinusSign                CharCode = 0x00b1
	GreaterThanOrEqualTo         CharCode = 0x2265
	LessThanOrEqualTo            CharCode = 0x2264
	TopHalfIntegral              CharCode = 0x2320
	BottomHalfIntegral           CharCode = 0x2321
	DivisionSign                 CharCode = 0x00f7
	AlmostEqualTo                CharCode = 0x2248
	DegreeSign                   CharCode = 0x00b0
	BulletOperator               CharCode = 0x2219
	MiddleDot                    CharCode = 0x00b7
	SquareRoot                   CharCode = 0x221a
	SuperscriptLatinSmallLetterN CharCode = 0x207f
	SuperscriptTwo               CharCode = 0x00b2
	BlackSquare                  CharCode = 0x25a0
)
