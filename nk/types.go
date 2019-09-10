// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 10 Sep 2019 22:23:05 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package nk

/*
#include "nk.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Char type as declared in nk/nuklear.h:398
type Char byte

// Uchar type as declared in nk/nuklear.h:399
type Uchar byte

// Byte type as declared in nk/nuklear.h:400
type Byte byte

// Short type as declared in nk/nuklear.h:401
type Short int16

// Ushort type as declared in nk/nuklear.h:402
type Ushort uint16

// Int type as declared in nk/nuklear.h:403
type Int int32

// Uint type as declared in nk/nuklear.h:404
type Uint uint32

// Size type as declared in nk/nuklear.h:405
type Size uint

// Ptr type as declared in nk/nuklear.h:406
type Ptr uint

// Hash type as declared in nk/nuklear.h:408
type Hash uint32

// Flags type as declared in nk/nuklear.h:409
type Flags uint32

// Rune type as declared in nk/nuklear.h:410
type Rune uint32

// Glyph type as declared in nk/nuklear.h:463
type Glyph [4]byte

// Handle as declared in nk/nuklear.h:464
const sizeofHandle = unsafe.Sizeof(C.nk_handle{})

type Handle [sizeofHandle]byte

// PluginAlloc type as declared in nk/nuklear.h:482
type PluginAlloc func(arg0 Handle, old unsafe.Pointer, arg2 Size) unsafe.Pointer

// PluginFree type as declared in nk/nuklear.h:483
type PluginFree func(arg0 Handle, old unsafe.Pointer)

// PluginFilter type as declared in nk/nuklear.h:484
type PluginFilter func(arg0 *TextEdit, unicode Rune) int32

// PluginPaste type as declared in nk/nuklear.h:485
type PluginPaste func(arg0 Handle, arg1 *TextEdit)

// PluginCopy type as declared in nk/nuklear.h:486
type PluginCopy func(arg0 Handle, arg1 string, len int32)

// TextWidthF type as declared in nk/nuklear.h:3896
type TextWidthF func(arg0 Handle, h float32, arg2 string, len int32) float32

// QueryFontGlyphF type as declared in nk/nuklear.h:3897
type QueryFontGlyphF func(handle Handle, fontHeight float32, glyph *UserFontGlyph, codepoint Rune, nextCodepoint Rune)

// CommandCustomCallback type as declared in nk/nuklear.h:4535
type CommandCustomCallback func(canvas unsafe.Pointer, x int16, y int16, w uint16, h uint16, callbackData Handle)

// DrawIndex type as declared in nk/nuklear.h:4668
type DrawIndex uint16

// Allocator as declared in nk/nuklear.h:431
type Allocator C.struct_nk_allocator

// BakedFont as declared in nk/nuklear.h:3936
type BakedFont C.struct_nk_baked_font

// Buffer as declared in nk/nuklear.h:430
type Buffer C.struct_nk_buffer

// BufferMarker as declared in nk/nuklear.h:4112
type BufferMarker C.struct_nk_buffer_marker

// Chart as declared in nk/nuklear.h:5278
type Chart C.struct_nk_chart

// ChartSlot as declared in nk/nuklear.h:5268
type ChartSlot C.struct_nk_chart_slot

// Clipboard as declared in nk/nuklear.h:4247
type Clipboard C.struct_nk_clipboard

// Color as declared in nk/nuklear.h:457
type Color C.struct_nk_color

// Colorf as declared in nk/nuklear.h:458
type Colorf C.struct_nk_colorf

// Command as declared in nk/nuklear.h:4397
type Command C.struct_nk_command

// CommandArc as declared in nk/nuklear.h:4487
type CommandArc C.struct_nk_command_arc

// CommandArcFilled as declared in nk/nuklear.h:4496
type CommandArcFilled C.struct_nk_command_arc_filled

// CommandBuffer as declared in nk/nuklear.h:432
type CommandBuffer C.struct_nk_command_buffer

// CommandCircle as declared in nk/nuklear.h:4472
type CommandCircle C.struct_nk_command_circle

// CommandCircleFilled as declared in nk/nuklear.h:4480
type CommandCircleFilled C.struct_nk_command_circle_filled

// CommandCurve as declared in nk/nuklear.h:4419
type CommandCurve C.struct_nk_command_curve

// CommandCustom as declared in nk/nuklear.h:4537
type CommandCustom C.struct_nk_command_custom

// CommandImage as declared in nk/nuklear.h:4527
type CommandImage C.struct_nk_command_image

// CommandLine as declared in nk/nuklear.h:4411
type CommandLine C.struct_nk_command_line

// CommandPolygon as declared in nk/nuklear.h:4504
type CommandPolygon C.struct_nk_command_polygon

// CommandPolygonFilled as declared in nk/nuklear.h:4512
type CommandPolygonFilled C.struct_nk_command_polygon_filled

// CommandPolyline as declared in nk/nuklear.h:4519
type CommandPolyline C.struct_nk_command_polyline

// CommandRect as declared in nk/nuklear.h:4428
type CommandRect C.struct_nk_command_rect

// CommandRectFilled as declared in nk/nuklear.h:4437
type CommandRectFilled C.struct_nk_command_rect_filled

// CommandRectMultiColor as declared in nk/nuklear.h:4445
type CommandRectMultiColor C.struct_nk_command_rect_multi_color

// CommandScissor as declared in nk/nuklear.h:4405
type CommandScissor C.struct_nk_command_scissor

// CommandText as declared in nk/nuklear.h:4545
type CommandText C.struct_nk_command_text

// CommandTriangle as declared in nk/nuklear.h:4455
type CommandTriangle C.struct_nk_command_triangle

// CommandTriangleFilled as declared in nk/nuklear.h:4464
type CommandTriangleFilled C.struct_nk_command_triangle_filled

// ConfigStackButtonBehavior as declared in nk/nuklear.h:5514
type ConfigStackButtonBehavior C.struct_nk_config_stack_button_behavior

// ConfigStackButtonBehaviorElement as declared in nk/nuklear.h:5506
type ConfigStackButtonBehaviorElement C.struct_nk_config_stack_button_behavior_element

// ConfigStackColor as declared in nk/nuklear.h:5512
type ConfigStackColor C.struct_nk_config_stack_color

// ConfigStackColorElement as declared in nk/nuklear.h:5504
type ConfigStackColorElement C.struct_nk_config_stack_color_element

// ConfigStackFlags as declared in nk/nuklear.h:5511
type ConfigStackFlags C.struct_nk_config_stack_flags

// ConfigStackFlagsElement as declared in nk/nuklear.h:5503
type ConfigStackFlagsElement C.struct_nk_config_stack_flags_element

// ConfigStackFloat as declared in nk/nuklear.h:5509
type ConfigStackFloat C.struct_nk_config_stack_float

// ConfigStackFloatElement as declared in nk/nuklear.h:5501
type ConfigStackFloatElement C.struct_nk_config_stack_float_element

// ConfigStackStyleItem as declared in nk/nuklear.h:5508
type ConfigStackStyleItem C.struct_nk_config_stack_style_item

// ConfigStackStyleItemElement as declared in nk/nuklear.h:5500
type ConfigStackStyleItemElement C.struct_nk_config_stack_style_item_element

// ConfigStackUserFont as declared in nk/nuklear.h:5513
type ConfigStackUserFont struct {
	Head           int32
	Elements       [8]ConfigStackUserFontElement
	refa664861d    *C.struct_nk_config_stack_user_font
	allocsa664861d interface{}
}

// ConfigStackUserFontElement as declared in nk/nuklear.h:5505
type ConfigStackUserFontElement struct {
	Address        [][]UserFont
	OldValue       []UserFont
	ref5572630c    *C.struct_nk_config_stack_user_font_element
	allocs5572630c interface{}
}

// ConfigStackVec2 as declared in nk/nuklear.h:5510
type ConfigStackVec2 C.struct_nk_config_stack_vec2

// ConfigStackVec2Element as declared in nk/nuklear.h:5502
type ConfigStackVec2Element C.struct_nk_config_stack_vec2_element

// ConfigurationStacks as declared in nk/nuklear.h:5516
type ConfigurationStacks C.struct_nk_configuration_stacks

// Context as declared in nk/nuklear.h:440
type Context C.struct_nk_context

// ConvertConfig as declared in nk/nuklear.h:434
type ConvertConfig struct {
	GlobalAlpha        float32
	LineAa             AntiAliasing
	ShapeAa            AntiAliasing
	CircleSegmentCount uint32
	ArcSegmentCount    uint32
	CurveSegmentCount  uint32
	Null               DrawNullTexture
	VertexLayout       []DrawVertexLayoutElement
	VertexSize         Size
	VertexAlignment    Size
	ref82bf4c25        *C.struct_nk_convert_config
	allocs82bf4c25     interface{}
}

// Cursor as declared in nk/nuklear.h:466
type Cursor C.struct_nk_cursor

// DrawCommand as declared in nk/nuklear.h:433
type DrawCommand C.struct_nk_draw_command

// DrawList as declared in nk/nuklear.h:437
type DrawList C.struct_nk_draw_list

// DrawNullTexture as declared in nk/nuklear.h:1150
type DrawNullTexture C.struct_nk_draw_null_texture

// DrawVertexLayoutElement as declared in nk/nuklear.h:441
type DrawVertexLayoutElement struct {
	Attribute      DrawVertexLayoutAttribute
	Format         DrawVertexLayoutFormat
	Offset         Size
	refeb0614d6    *C.struct_nk_draw_vertex_layout_element
	allocseb0614d6 interface{}
}

// EditState as declared in nk/nuklear.h:5382
type EditState C.struct_nk_edit_state

// Font as declared in nk/nuklear.h:3935
type Font C.struct_nk_font

// FontAtlas as declared in nk/nuklear.h:4009
type FontAtlas C.struct_nk_font_atlas

// FontConfig as declared in nk/nuklear.h:3949
type FontConfig C.struct_nk_font_config

// FontGlyph as declared in nk/nuklear.h:3985
type FontGlyph C.struct_nk_font_glyph

// Image as declared in nk/nuklear.h:465
type Image C.struct_nk_image

// Input as declared in nk/nuklear.h:4625
type Input C.struct_nk_input

// Key as declared in nk/nuklear.h:4615
type Key C.struct_nk_key

// Keyboard as declared in nk/nuklear.h:4619
type Keyboard C.struct_nk_keyboard

// ListView as declared in nk/nuklear.h:3025
type ListView C.struct_nk_list_view

// Memory as declared in nk/nuklear.h:4117
type Memory C.struct_nk_memory

// MemoryStatus as declared in nk/nuklear.h:4092
type MemoryStatus C.struct_nk_memory_status

// MenuState as declared in nk/nuklear.h:5320
type MenuState C.struct_nk_menu_state

// Mouse as declared in nk/nuklear.h:4604
type Mouse C.struct_nk_mouse

// MouseButton as declared in nk/nuklear.h:4599
type MouseButton C.struct_nk_mouse_button

// Page as declared in nk/nuklear.h:5552
type Page C.struct_nk_page

// PageData as declared in nk/nuklear.h:5540
const sizeofPageData = unsafe.Sizeof(C.union_nk_page_data{})

type PageData [sizeofPageData]byte

// PageElement as declared in nk/nuklear.h:5546
type PageElement C.struct_nk_page_element

// Panel as declared in nk/nuklear.h:439
type Panel C.struct_nk_panel

// Pool as declared in nk/nuklear.h:5558
type Pool C.struct_nk_pool

// PopupBuffer as declared in nk/nuklear.h:5312
type PopupBuffer C.struct_nk_popup_buffer

// PopupState as declared in nk/nuklear.h:5370
type PopupState C.struct_nk_popup_state

// PropertyState as declared in nk/nuklear.h:5395
type PropertyState C.struct_nk_property_state

// Rect as declared in nk/nuklear.h:461
type Rect C.struct_nk_rect

// Recti as declared in nk/nuklear.h:462
type Recti C.struct_nk_recti

// RowLayout as declared in nk/nuklear.h:5296
type RowLayout C.struct_nk_row_layout

// Scroll as declared in nk/nuklear.h:467
type Scroll C.struct_nk_scroll

// Str as declared in nk/nuklear.h:4164
type Str C.struct_nk_str

// Style as declared in nk/nuklear.h:5212
type Style C.struct_nk_style

// StyleButton as declared in nk/nuklear.h:442
type StyleButton C.struct_nk_style_button

// StyleChart as declared in nk/nuklear.h:450
type StyleChart C.struct_nk_style_chart

// StyleCombo as declared in nk/nuklear.h:451
type StyleCombo C.struct_nk_style_combo

// StyleEdit as declared in nk/nuklear.h:448
type StyleEdit C.struct_nk_style_edit

// StyleItem as declared in nk/nuklear.h:435
type StyleItem C.struct_nk_style_item

// StyleItemData as declared in nk/nuklear.h:4810
const sizeofStyleItemData = unsafe.Sizeof(C.union_nk_style_item_data{})

type StyleItemData [sizeofStyleItemData]byte

// StyleProgress as declared in nk/nuklear.h:446
type StyleProgress C.struct_nk_style_progress

// StyleProperty as declared in nk/nuklear.h:449
type StyleProperty C.struct_nk_style_property

// StyleScrollbar as declared in nk/nuklear.h:447
type StyleScrollbar C.struct_nk_style_scrollbar

// StyleSelectable as declared in nk/nuklear.h:444
type StyleSelectable C.struct_nk_style_selectable

// StyleSlide as declared in nk/nuklear.h:445
type StyleSlide C.struct_nk_style_slide

// StyleSlider as declared in nk/nuklear.h:4917
type StyleSlider C.struct_nk_style_slider

// StyleTab as declared in nk/nuklear.h:452
type StyleTab C.struct_nk_style_tab

// StyleText as declared in nk/nuklear.h:4820
type StyleText C.struct_nk_style_text

// StyleToggle as declared in nk/nuklear.h:443
type StyleToggle C.struct_nk_style_toggle

// StyleWindow as declared in nk/nuklear.h:454
type StyleWindow C.struct_nk_style_window

// StyleWindowHeader as declared in nk/nuklear.h:453
type StyleWindowHeader C.struct_nk_style_window_header

// Table as declared in nk/nuklear.h:5351
type Table C.struct_nk_table

// TextEdit as declared in nk/nuklear.h:436
type TextEdit C.struct_nk_text_edit

// TextUndoRecord as declared in nk/nuklear.h:4253
type TextUndoRecord C.struct_nk_text_undo_record

// TextUndoState as declared in nk/nuklear.h:4260
type TextUndoState C.struct_nk_text_undo_state

// UserFont as declared in nk/nuklear.h:438
type UserFont struct {
	Userdata       Handle
	Height         float32
	Width          TextWidthF
	Query          QueryFontGlyphF
	Texture        Handle
	ref738ce62e    *C.struct_nk_user_font
	allocs738ce62e interface{}
}

// UserFontGlyph as declared in nk/nuklear.h:3895
type UserFontGlyph struct {
	Uv             [2]Vec2
	Offset         Vec2
	Width          float32
	Height         float32
	Xadvance       float32
	ref4a84b297    *C.struct_nk_user_font_glyph
	allocs4a84b297 interface{}
}

// Vec2 as declared in nk/nuklear.h:459
type Vec2 C.struct_nk_vec2

// Vec2i as declared in nk/nuklear.h:460
type Vec2i C.struct_nk_vec2i

// Window as declared in nk/nuklear.h:5408
type Window C.struct_nk_window
