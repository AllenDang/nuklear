// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 29 Oct 2016 14:52:15 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package nk

/*
#cgo CFLAGS: -DNK_INCLUDE_FIXED_TYPES -DNK_INCLUDE_STANDARD_IO -DNK_INCLUDE_DEFAULT_ALLOCATOR -DNK_INCLUDE_FONT_BAKING -DNK_INCLUDE_DEFAULT_FONT -Wno-implicit-function-declaration
#include "nuklear.h"
#include "nuklear_glfw_gl2.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// IncludeFixedTypes as defined in nuklear/<predefine>:20
	IncludeFixedTypes = 1
	// IncludeStandardIo as defined in nuklear/<predefine>:21
	IncludeStandardIo = 1
	// IncludeDefaultAllocator as defined in nuklear/<predefine>:22
	IncludeDefaultAllocator = 1
	// IncludeFontBaking as defined in nuklear/<predefine>:23
	IncludeFontBaking = 1
	// IncludeDefaultFont as defined in nuklear/<predefine>:24
	IncludeDefaultFont = 1
	// Undefined as defined in nk/nuklear.h:254
	Undefined = (-1.0)
	// UtfInvalid as defined in nk/nuklear.h:255
	UtfInvalid = 0xFFFD
	// UtfSize as defined in nk/nuklear.h:256
	UtfSize = 4
	// InputMax as defined in nk/nuklear.h:258
	InputMax = 16
	// MaxNumberBuffer as defined in nk/nuklear.h:261
	MaxNumberBuffer = 64
	// ScrollbarHidingTimeout as defined in nk/nuklear.h:264
	ScrollbarHidingTimeout = 4.0
	// TexteditUndostatecount as defined in nk/nuklear.h:1267
	TexteditUndostatecount = 99
	// TexteditUndocharcount as defined in nk/nuklear.h:1271
	TexteditUndocharcount = 999
	// ChartMaxSlot as defined in nk/nuklear.h:2430
	ChartMaxSlot = 4
	// WindowMaxName as defined in nk/nuklear.h:2513
	WindowMaxName = 64
	// ButtonBehaviorStackSize as defined in nk/nuklear.h:2599
	ButtonBehaviorStackSize = 8
	// FontStackSize as defined in nk/nuklear.h:2603
	FontStackSize = 8
	// StyleItemStackSize as defined in nk/nuklear.h:2607
	StyleItemStackSize = 16
	// VectorStackSize as defined in nk/nuklear.h:2615
	VectorStackSize = 16
	// FlagsStackSize as defined in nk/nuklear.h:2619
	FlagsStackSize = 32
	// ColorStackSize as defined in nk/nuklear.h:2623
	ColorStackSize = 32
	// Float as defined in nk/nuklear.h:2637
	Float = 0
	// Pi as defined in nk/nuklear.h:2752
	Pi = 3.141592654
)

// Heading as declared in nk/nuklear.h:442
type Heading int32

// Heading enumeration from nk/nuklear.h:442
const (
	Up    = iota
	Right = 1
	Down  = 2
	Left  = 3
)

// ButtonBehavior as declared in nk/nuklear.h:444
type ButtonBehavior int32

// ButtonBehavior enumeration from nk/nuklear.h:444
const (
	ButtonDefault  = iota
	ButtonRepeater = 1
)

// Modify as declared in nk/nuklear.h:445
type Modify int32

// Modify enumeration from nk/nuklear.h:445
const (
	Fixed      = iota
	Modifiable = 1
)

// Orientation as declared in nk/nuklear.h:446
type Orientation int32

// Orientation enumeration from nk/nuklear.h:446
const (
	Vertical   = iota
	Horizontal = 1
)

// CollapseStates as declared in nk/nuklear.h:447
type CollapseStates int32

// CollapseStates enumeration from nk/nuklear.h:447
const (
	Minimized = iota
	Maximized = 1
)

// ShowStates as declared in nk/nuklear.h:448
type ShowStates int32

// ShowStates enumeration from nk/nuklear.h:448
const (
	Hidden = iota
	Shown  = 1
)

// ChartType as declared in nk/nuklear.h:449
type ChartType int32

// ChartType enumeration from nk/nuklear.h:449
const (
	ChartLines  = iota
	ChartColumn = 1
	ChartMax    = 2
)

// ChartEvent as declared in nk/nuklear.h:450
type ChartEvent int32

// ChartEvent enumeration from nk/nuklear.h:450
const (
	ChartHovering = 1
	ChartClicked  = 2
)

// ColorFormat as declared in nk/nuklear.h:451
type ColorFormat int32

// ColorFormat enumeration from nk/nuklear.h:451
const (
	ColorFormatRGB  = iota
	ColorFormatRGBA = 1
)

// PopupType as declared in nk/nuklear.h:452
type PopupType int32

// PopupType enumeration from nk/nuklear.h:452
const (
	PopupStatic  = iota
	PopupDynamic = 1
)

// LayoutFormat as declared in nk/nuklear.h:453
type LayoutFormat int32

// LayoutFormat enumeration from nk/nuklear.h:453
const (
	Dynamic = iota
	Static  = 1
)

// TreeType as declared in nk/nuklear.h:454
type TreeType int32

// TreeType enumeration from nk/nuklear.h:454
const (
	TreeNode = iota
	TreeTab  = 1
)

// AntiAliasing as declared in nk/nuklear.h:455
type AntiAliasing int32

// AntiAliasing enumeration from nk/nuklear.h:455
const (
	AntiAliasingOff = iota
	AntiAliasingOn  = 1
)

// SymbolType as declared in nk/nuklear.h:486
type SymbolType int32

// SymbolType enumeration from nk/nuklear.h:486
const (
	SymbolNone          = iota
	SymbolX             = 1
	SymbolUnderscore    = 2
	SymbolCircleSolid   = 3
	SymbolCircleOutline = 4
	SymbolRectSolid     = 5
	SymbolRectOutline   = 6
	SymbolTriangleUp    = 7
	SymbolTriangleDown  = 8
	SymbolTriangleLeft  = 9
	SymbolTriangleRight = 10
	SymbolPlus          = 11
	SymbolMinus         = 12
	SymbolMax           = 13
)

// Keys as declared in nk/nuklear.h:503
type Keys int32

// Keys enumeration from nk/nuklear.h:503
const (
	KeyNone            = iota
	KeyShift           = 1
	KeyCtrl            = 2
	KeyDel             = 3
	KeyEnter           = 4
	KeyTab             = 5
	KeyBackspace       = 6
	KeyCopy            = 7
	KeyCut             = 8
	KeyPaste           = 9
	KeyUp              = 10
	KeyDown            = 11
	KeyLeft            = 12
	KeyRight           = 13
	KeyTextInsertMode  = 14
	KeyTextReplaceMode = 15
	KeyTextResetMode   = 16
	KeyTextLineStart   = 17
	KeyTextLineEnd     = 18
	KeyTextStart       = 19
	KeyTextEnd         = 20
	KeyTextUndo        = 21
	KeyTextRedo        = 22
	KeyTextWordLeft    = 23
	KeyTextWordRight   = 24
	KeyScrollStart     = 25
	KeyScrollEnd       = 26
	KeyScrollDown      = 27
	KeyScrollUp        = 28
	KeyMax             = 29
)

// Buttons as declared in nk/nuklear.h:541
type Buttons int32

// Buttons enumeration from nk/nuklear.h:541
const (
	ButtonLeft   = iota
	ButtonMiddle = 1
	ButtonRight  = 2
	ButtonMax    = 3
)

// StyleColors as declared in nk/nuklear.h:548
type StyleColors int32

// StyleColors enumeration from nk/nuklear.h:548
const (
	ColorText                  = iota
	ColorWindow                = 1
	ColorHeader                = 2
	ColorBorder                = 3
	ColorButton                = 4
	ColorButtonHover           = 5
	ColorButtonActive          = 6
	ColorToggle                = 7
	ColorToggleHover           = 8
	ColorToggleCursor          = 9
	ColorSelect                = 10
	ColorSelectActive          = 11
	ColorSlider                = 12
	ColorSliderCursor          = 13
	ColorSliderCursorHover     = 14
	ColorSliderCursorActive    = 15
	ColorProperty              = 16
	ColorEdit                  = 17
	ColorEditCursor            = 18
	ColorCombo                 = 19
	ColorChart                 = 20
	ColorChartColor            = 21
	ColorChartColorHighlight   = 22
	ColorScrollbar             = 23
	ColorScrollbarCursor       = 24
	ColorScrollbarCursorHover  = 25
	ColorScrollbarCursorActive = 26
	ColorTabHeader             = 27
	ColorCount                 = 28
)

// StyleCursor as declared in nk/nuklear.h:580
type StyleCursor int32

// StyleCursor enumeration from nk/nuklear.h:580
const (
	CursorArrow                  = iota
	CursorText                   = 1
	CursorMove                   = 2
	CursorResizeVertical         = 3
	CursorResizeHorizontal       = 4
	CursorResizeTopLeftDownRight = 5
	CursorResizeTopRightDownLeft = 6
	CursorCount                  = 7
)

// WidgetLayoutStates as declared in nk/nuklear.h:591
type WidgetLayoutStates int32

// WidgetLayoutStates enumeration from nk/nuklear.h:591
const (
	WidgetInvalid = iota
	WidgetValid   = 1
	WidgetRom     = 2
)

// WidgetStates as declared in nk/nuklear.h:598
type WidgetStates int32

// WidgetStates enumeration from nk/nuklear.h:598
const (
	WidgetStateModified = 2
	WidgetStateInactive = 4
	WidgetStateEntered  = 8
	WidgetStateHover    = 16
	WidgetStateActived  = 32
	WidgetStateLeft     = 64
	WidgetStateHovered  = 18
	WidgetStateActive   = 34
)

// TextAlign as declared in nk/nuklear.h:610
type TextAlign int32

// TextAlign enumeration from nk/nuklear.h:610
const (
	TextAlignLeft     = 1
	TextAlignCentered = 2
	TextAlignRight    = 4
	TextAlignTop      = 8
	TextAlignMiddle   = 16
	TextAlignBottom   = 32
)

// TextAlignment as declared in nk/nuklear.h:618
type TextAlignment int32

// TextAlignment enumeration from nk/nuklear.h:618
const (
	TextLeft     = 17
	TextCentered = 18
	TextRight    = 20
)

// EditFlags as declared in nk/nuklear.h:625
type EditFlags int32

// EditFlags enumeration from nk/nuklear.h:625
const (
	EditDefault            = iota
	EditReadOnly           = 1
	EditAutoSelect         = 2
	EditSigEnter           = 4
	EditAllowTab           = 8
	EditNoCursor           = 16
	EditSelectable         = 32
	EditClipboard          = 64
	EditCtrlEnterNewline   = 128
	EditNoHorizontalScroll = 256
	EditAlwaysInsertMode   = 512
	EditMultiline          = 2048
	EditGotoEndOnActivate  = 4096
)

// EditTypes as declared in nk/nuklear.h:640
type EditTypes int32

// EditTypes enumeration from nk/nuklear.h:640
const (
	EditSimple = 512
	EditField  = 608
	EditBox    = 2664
	EditEditor = 2152
)

// EditEvents as declared in nk/nuklear.h:646
type EditEvents int32

// EditEvents enumeration from nk/nuklear.h:646
const (
	EditActive      = 1
	EditInactive    = 2
	EditActivated   = 4
	EditDeactivated = 8
	EditCommited    = 16
)

// PanelFlags as declared in nk/nuklear.h:654
type PanelFlags int32

// PanelFlags enumeration from nk/nuklear.h:654
const (
	WindowBorder         = 1
	WindowMovable        = 2
	WindowScalable       = 4
	WindowClosable       = 8
	WindowMinimizable    = 16
	WindowNoScrollbar    = 32
	WindowTitle          = 64
	WindowScrollAutoHide = 128
	WindowBackground     = 256
)

// AllocationType as declared in nk/nuklear.h:1131
type AllocationType int32

// AllocationType enumeration from nk/nuklear.h:1131
const (
	BufferFixed   = iota
	BufferDynamic = 1
)

// BufferAllocationType as declared in nk/nuklear.h:1136
type BufferAllocationType int32

// BufferAllocationType enumeration from nk/nuklear.h:1136
const (
	BufferFront = iota
	BufferBack  = 1
	BufferMax   = 2
)

// TextEditType as declared in nk/nuklear.h:1297
type TextEditType int32

// TextEditType enumeration from nk/nuklear.h:1297
const (
	TextEditSingleLine = iota
	TextEditMultiLine  = 1
)

// TextEditMode as declared in nk/nuklear.h:1302
type TextEditMode int32

// TextEditMode enumeration from nk/nuklear.h:1302
const (
	TextEditModeView    = iota
	TextEditModeInsert  = 1
	TextEditModeReplace = 2
)

// FontCoordType as declared in nk/nuklear.h:1430
type FontCoordType int32

// FontCoordType enumeration from nk/nuklear.h:1430
const (
	CoordUv    = iota
	CoordPixel = 1
)

// FontAtlasFormat as declared in nk/nuklear.h:1501
type FontAtlasFormat int32

// FontAtlasFormat enumeration from nk/nuklear.h:1501
const (
	FontAtlasAlpha8 = iota
	FontAtlasRgba32 = 1
)

// CommandType as declared in nk/nuklear.h:1575
type CommandType int32

// CommandType enumeration from nk/nuklear.h:1575
const (
	CommandTypeNop            = iota
	CommandTypeScissor        = 1
	CommandTypeLine           = 2
	CommandTypeCurve          = 3
	CommandTypeRect           = 4
	CommandTypeRectFilled     = 5
	CommandTypeRectMultiColor = 6
	CommandTypeCircle         = 7
	CommandTypeCircleFilled   = 8
	CommandTypeArc            = 9
	CommandTypeArcFilled      = 10
	CommandTypeTriangle       = 11
	CommandTypeTriangleFilled = 12
	CommandTypePolygon        = 13
	CommandTypePolygonFilled  = 14
	CommandTypePolyline       = 15
	CommandTypeText           = 16
	CommandTypeImage          = 17
)

// CommandClipping as declared in nk/nuklear.h:1747
type CommandClipping int32

// CommandClipping enumeration from nk/nuklear.h:1747
const (
	ClippingOff = iota
	ClippingOn  = 1
)

// StyleItemType as declared in nk/nuklear.h:1989
type StyleItemType int32

// StyleItemType enumeration from nk/nuklear.h:1989
const (
	StyleItemColor = iota
	StyleItemImage = 1
)

// StyleHeaderAlign as declared in nk/nuklear.h:2330
type StyleHeaderAlign int32

// StyleHeaderAlign enumeration from nk/nuklear.h:2330
const (
	HeaderLeft  = iota
	HeaderRight = 1
)

// PanelType as declared in nk/nuklear.h:2433
type PanelType int32

// PanelType enumeration from nk/nuklear.h:2433
const (
	PanelWindow     = 1
	PanelGroup      = 2
	PanelPopup      = 4
	PanelContextual = 16
	PanelCombo      = 32
	PanelMenu       = 64
	PanelTooltip    = 128
)

// PanelSet as declared in nk/nuklear.h:2442
type PanelSet int32

// PanelSet enumeration from nk/nuklear.h:2442
const (
	PanelSetNonblock = 240
	PanelSetPopup    = 244
	PanelSetSub      = 246
)

// WindowFlags as declared in nk/nuklear.h:2517
type WindowFlags int32

// WindowFlags enumeration from nk/nuklear.h:2517
const (
	WindowPrivate   = 1024
	WindowDynamic   = 1024
	WindowRom       = 2048
	WindowHidden    = 4096
	WindowClosed    = 8192
	WindowMinimized = 16384
	WindowRemoveRom = 32768
)

// GLFWInitState as declared in nk/nuklear_glfw_gl2.h:18
type GLFWInitState int32

// GLFWInitState enumeration from nk/nuklear_glfw_gl2.h:18
const (
	GLFW3Default          = iota
	GLFW3InstallCallbacks = 1
)

const (
	// False as declared in nk/nuklear.h:430
	False = iota
	// True as declared in nk/nuklear.h:430
	True = 1
)
