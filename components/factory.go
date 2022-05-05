package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/aboutmodal"
	"github.com/mlctrez/goapp-pf/components/accordion"
	"github.com/mlctrez/goapp-pf/components/actionlist"
	"github.com/mlctrez/goapp-pf/components/alert"
	"github.com/mlctrez/goapp-pf/components/alertgroup"
	"github.com/mlctrez/goapp-pf/components/applauncher"
	"github.com/mlctrez/goapp-pf/components/avatar"
	"github.com/mlctrez/goapp-pf/components/backdrop"
	"github.com/mlctrez/goapp-pf/components/backgroundimage"
	"github.com/mlctrez/goapp-pf/components/backtotop"
	"github.com/mlctrez/goapp-pf/components/badge"
	"github.com/mlctrez/goapp-pf/components/banner"
	"github.com/mlctrez/goapp-pf/components/brand"
	"github.com/mlctrez/goapp-pf/components/breadcrumb"
	"github.com/mlctrez/goapp-pf/components/button"
	"github.com/mlctrez/goapp-pf/components/calendarmonth"
	"github.com/mlctrez/goapp-pf/components/card"
	"github.com/mlctrez/goapp-pf/components/checkbox"
	"github.com/mlctrez/goapp-pf/components/chip"
	"github.com/mlctrez/goapp-pf/components/chipgroup"
	"github.com/mlctrez/goapp-pf/components/clipboardcopy"
	"github.com/mlctrez/goapp-pf/components/codeblock"
	"github.com/mlctrez/goapp-pf/components/codeeditor"
	"github.com/mlctrez/goapp-pf/components/contextselector"
	"github.com/mlctrez/goapp-pf/components/datalist"
	"github.com/mlctrez/goapp-pf/components/datepicker"
	"github.com/mlctrez/goapp-pf/components/descriptionlist"
	"github.com/mlctrez/goapp-pf/components/divider"
	"github.com/mlctrez/goapp-pf/components/draganddrop"
	"github.com/mlctrez/goapp-pf/components/drawer"
	"github.com/mlctrez/goapp-pf/components/dropdown"
	"github.com/mlctrez/goapp-pf/components/duallistselector"
	"github.com/mlctrez/goapp-pf/components/emptystate"
	"github.com/mlctrez/goapp-pf/components/expandablesection"
	"github.com/mlctrez/goapp-pf/components/fileupload"
	"github.com/mlctrez/goapp-pf/components/fileuploadmultiple"
	"github.com/mlctrez/goapp-pf/components/form"
	"github.com/mlctrez/goapp-pf/components/formcontrol"
	"github.com/mlctrez/goapp-pf/components/formselect"
	"github.com/mlctrez/goapp-pf/components/helpertext"
	"github.com/mlctrez/goapp-pf/components/hint"
	"github.com/mlctrez/goapp-pf/components/inlineedit"
	"github.com/mlctrez/goapp-pf/components/inputgroup"
	"github.com/mlctrez/goapp-pf/components/jumplinks"
	"github.com/mlctrez/goapp-pf/components/label"
	"github.com/mlctrez/goapp-pf/components/labelgroup"
	"github.com/mlctrez/goapp-pf/components/list"
	"github.com/mlctrez/goapp-pf/components/loginpage"
	"github.com/mlctrez/goapp-pf/components/masthead"
	"github.com/mlctrez/goapp-pf/components/menu"
	"github.com/mlctrez/goapp-pf/components/menutoggle"
	"github.com/mlctrez/goapp-pf/components/modal"
	"github.com/mlctrez/goapp-pf/components/navigation"
	"github.com/mlctrez/goapp-pf/components/notificationbadge"
	"github.com/mlctrez/goapp-pf/components/notificationdrawer"
	"github.com/mlctrez/goapp-pf/components/numberinput"
	"github.com/mlctrez/goapp-pf/components/optionsmenu"
	"github.com/mlctrez/goapp-pf/components/overflowmenu"
	"github.com/mlctrez/goapp-pf/components/page"
	"github.com/mlctrez/goapp-pf/components/pagination"
	"github.com/mlctrez/goapp-pf/components/panel"
	"github.com/mlctrez/goapp-pf/components/pfselect"
	"github.com/mlctrez/goapp-pf/components/pfswitch"
	"github.com/mlctrez/goapp-pf/components/popover"
	"github.com/mlctrez/goapp-pf/components/progress"
	"github.com/mlctrez/goapp-pf/components/progressstepper"
	"github.com/mlctrez/goapp-pf/components/radio"
	"github.com/mlctrez/goapp-pf/components/searchinput"
	"github.com/mlctrez/goapp-pf/components/sidebar"
	"github.com/mlctrez/goapp-pf/components/simplelist"
	"github.com/mlctrez/goapp-pf/components/skeleton"
	"github.com/mlctrez/goapp-pf/components/skiptocontent"
	"github.com/mlctrez/goapp-pf/components/slider"
	"github.com/mlctrez/goapp-pf/components/spinner"
	"github.com/mlctrez/goapp-pf/components/tabcontent"
	"github.com/mlctrez/goapp-pf/components/table"
	"github.com/mlctrez/goapp-pf/components/tabs"
	"github.com/mlctrez/goapp-pf/components/text"
	"github.com/mlctrez/goapp-pf/components/textarea"
	"github.com/mlctrez/goapp-pf/components/textinput"
	"github.com/mlctrez/goapp-pf/components/textinputgroup"
	"github.com/mlctrez/goapp-pf/components/tile"
	"github.com/mlctrez/goapp-pf/components/timepicker"
	"github.com/mlctrez/goapp-pf/components/title"
	"github.com/mlctrez/goapp-pf/components/togglegroup"
	"github.com/mlctrez/goapp-pf/components/toolbar"
	"github.com/mlctrez/goapp-pf/components/tooltip"
	"github.com/mlctrez/goapp-pf/components/treeview"
	"github.com/mlctrez/goapp-pf/components/truncate"
	"github.com/mlctrez/goapp-pf/components/wizard"
)

func Text() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &text.Basic{} }
	return result
}

func TimePicker() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic12Hour"] = func() app.UI { return &timepicker.Basic12Hour{} }
	result["Basic24Hour"] = func() app.UI { return &timepicker.Basic24Hour{} }
	result["CustomDelimiter"] = func() app.UI { return &timepicker.CustomDelimiter{} }
	result["MinimummaximumTimes"] = func() app.UI { return &timepicker.MinimummaximumTimes{} }
	result["WithSeconds"] = func() app.UI { return &timepicker.WithSeconds{} }
	result["Basic24HoursWithSeconds"] = func() app.UI { return &timepicker.Basic24HoursWithSeconds{} }
	return result
}

func Divider() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Hr"] = func() app.UI { return &divider.Hr{} }
	result["Li"] = func() app.UI { return &divider.Li{} }
	result["Div"] = func() app.UI { return &divider.Div{} }
	result["InsetMedium"] = func() app.UI { return &divider.InsetMedium{} }
	result["MdInsetNoInsetOnMd3xlInsetOnLgLgInsetOnXl"] = func() app.UI { return &divider.MdInsetNoInsetOnMd3xlInsetOnLgLgInsetOnXl{} }
	result["Vertical"] = func() app.UI { return &divider.Vertical{} }
	result["VerticalInsetMedium"] = func() app.UI { return &divider.VerticalInsetMedium{} }
	result["VerticalMdInsetNoInsetOnMdLgInsetOnLgSmInsetOnXl"] = func() app.UI { return &divider.VerticalMdInsetNoInsetOnMdLgInsetOnLgSmInsetOnXl{} }
	result["VerticalOnLg"] = func() app.UI { return &divider.VerticalOnLg{} }
	result["HorizontalOnLg"] = func() app.UI { return &divider.HorizontalOnLg{} }
	return result
}

func InlineEdit() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["InlineEditToggle"] = func() app.UI { return &inlineedit.InlineEditToggle{} }
	result["InlineEditValue"] = func() app.UI { return &inlineedit.InlineEditValue{} }
	result["InlineEditActionGroup"] = func() app.UI { return &inlineedit.InlineEditActionGroup{} }
	result["InlineEditActionGroupIconButtons"] = func() app.UI { return &inlineedit.InlineEditActionGroupIconButtons{} }
	result["SingleInlineEditDefault"] = func() app.UI { return &inlineedit.SingleInlineEditDefault{} }
	result["SingleInlineEditActive"] = func() app.UI { return &inlineedit.SingleInlineEditActive{} }
	result["FreeFormEdit"] = func() app.UI { return &inlineedit.FreeFormEdit{} }
	result["SingleInlineEditWithLabelDefault"] = func() app.UI { return &inlineedit.SingleInlineEditWithLabelDefault{} }
	result["StateValid"] = func() app.UI { return &inlineedit.StateValid{} }
	result["StateInvalid"] = func() app.UI { return &inlineedit.StateInvalid{} }
	result["InlineEditTableRow"] = func() app.UI { return &inlineedit.InlineEditTableRow{} }
	return result
}

func JumpLinks() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["HorizontalDefault"] = func() app.UI { return &jumplinks.HorizontalDefault{} }
	result["HorizontalWithCenteredList"] = func() app.UI { return &jumplinks.HorizontalWithCenteredList{} }
	result["HorizontalWithLabel"] = func() app.UI { return &jumplinks.HorizontalWithLabel{} }
	result["VerticalDefault"] = func() app.UI { return &jumplinks.VerticalDefault{} }
	result["VerticalWithLabel"] = func() app.UI { return &jumplinks.VerticalWithLabel{} }
	result["VerticalWithInactiveSubsections"] = func() app.UI { return &jumplinks.VerticalWithInactiveSubsections{} }
	result["VerticalWithActiveSubsections"] = func() app.UI { return &jumplinks.VerticalWithActiveSubsections{} }
	result["Expandable"] = func() app.UI { return &jumplinks.Expandable{} }
	result["Expanded"] = func() app.UI { return &jumplinks.Expanded{} }
	result["ExpandableResponsive"] = func() app.UI { return &jumplinks.ExpandableResponsive{} }
	result["ExpandableResponsiveWithNoLabel"] = func() app.UI { return &jumplinks.ExpandableResponsiveWithNoLabel{} }
	return result
}

func TreeView() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &treeview.Default{} }
	result["WithSearch"] = func() app.UI { return &treeview.WithSearch{} }
	result["WithCheckboxes"] = func() app.UI { return &treeview.WithCheckboxes{} }
	result["WithIcons"] = func() app.UI { return &treeview.WithIcons{} }
	result["WithBadges"] = func() app.UI { return &treeview.WithBadges{} }
	result["WithActionItem"] = func() app.UI { return &treeview.WithActionItem{} }
	result["WithNonExpandableTopLevelNodes"] = func() app.UI { return &treeview.WithNonExpandableTopLevelNodes{} }
	result["Guides"] = func() app.UI { return &treeview.Guides{} }
	result["Compact"] = func() app.UI { return &treeview.Compact{} }
	result["CompactNoBackground"] = func() app.UI { return &treeview.CompactNoBackground{} }
	return result
}

func Backdrop() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &backdrop.Basic{} }
	return result
}

func ContextSelector() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &contextselector.Basic{} }
	result["PlainWithText"] = func() app.UI { return &contextselector.PlainWithText{} }
	result["WithFooter"] = func() app.UI { return &contextselector.WithFooter{} }
	return result
}

func InputGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Variations"] = func() app.UI { return &inputgroup.Variations{} }
	return result
}

func NumberInput() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &numberinput.Default{} }
	result["WithUnit"] = func() app.UI { return &numberinput.WithUnit{} }
	result["WithUnitAndLowerThresholdReached"] = func() app.UI { return &numberinput.WithUnitAndLowerThresholdReached{} }
	result["WithUnitAndUpperThresholdReached"] = func() app.UI { return &numberinput.WithUnitAndUpperThresholdReached{} }
	result["Disabled"] = func() app.UI { return &numberinput.Disabled{} }
	result["VaryingSizes"] = func() app.UI { return &numberinput.VaryingSizes{} }
	return result
}

func Table() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["BasicTableExample"] = func() app.UI { return &table.BasicTableExample{} }
	result["SortableExample"] = func() app.UI { return &table.SortableExample{} }
	result["OverflowMenuUsageDesktop"] = func() app.UI { return &table.OverflowMenuUsageDesktop{} }
	result["OverflowMenuUsageOverflowMenuCollapsed"] = func() app.UI { return &table.OverflowMenuUsageOverflowMenuCollapsed{} }
	result["OverflowMenuUsageMobile"] = func() app.UI { return &table.OverflowMenuUsageMobile{} }
	result["CheckboxesAndActionsExample"] = func() app.UI { return &table.CheckboxesAndActionsExample{} }
	result["SingleSelectRadioExample"] = func() app.UI { return &table.SingleSelectRadioExample{} }
	result["ExpandableExample"] = func() app.UI { return &table.ExpandableExample{} }
	result["ExpandableWithNestedTableExample"] = func() app.UI { return &table.ExpandableWithNestedTableExample{} }
	result["CompoundExpansionExample"] = func() app.UI { return &table.CompoundExpansionExample{} }
	result["CompactExample"] = func() app.UI { return &table.CompactExample{} }
	result["CompactExpandableExample"] = func() app.UI { return &table.CompactExpandableExample{} }
	result["HoverableAndSelectedExample"] = func() app.UI { return &table.HoverableAndSelectedExample{} }
	result["ExpandableHoverableAndSelectedExample"] = func() app.UI { return &table.ExpandableHoverableAndSelectedExample{} }
	result["TreeTableBasic"] = func() app.UI { return &table.TreeTableBasic{} }
	result["TreeTableWithCheckboxes"] = func() app.UI { return &table.TreeTableWithCheckboxes{} }
	result["TreeTableWithCheckboxesAndIcons"] = func() app.UI { return &table.TreeTableWithCheckboxesAndIcons{} }
	result["BorderlessExample"] = func() app.UI { return &table.BorderlessExample{} }
	result["BorderlessCompactExample"] = func() app.UI { return &table.BorderlessCompactExample{} }
	result["BorderlessExpandableExample"] = func() app.UI { return &table.BorderlessExpandableExample{} }
	result["BorderlessWithCompoundExpansionExample"] = func() app.UI { return &table.BorderlessWithCompoundExpansionExample{} }
	result["WidthModifiersExamples"] = func() app.UI { return &table.WidthModifiersExamples{} }
	result["HiddenvisibleBreakpointModifiersExample"] = func() app.UI { return &table.HiddenvisibleBreakpointModifiersExample{} }
	result["TextControlExample"] = func() app.UI { return &table.TextControlExample{} }
	result["TextControlUsingTheTableTextElementExample"] = func() app.UI { return &table.TextControlUsingTheTableTextElementExample{} }
	result["TableWithLongStringsInTheContent"] = func() app.UI { return &table.TableWithLongStringsInTheContent{} }
	result["WidthConstrained"] = func() app.UI { return &table.WidthConstrained{} }
	result["StickyHeader"] = func() app.UI { return &table.StickyHeader{} }
	result["StickyColumn"] = func() app.UI { return &table.StickyColumn{} }
	result["MultipleStickyColumns"] = func() app.UI { return &table.MultipleStickyColumns{} }
	result["StickyColumnsAndHeader"] = func() app.UI { return &table.StickyColumnsAndHeader{} }
	result["NestedColumnHeadersAndExpandableRows"] = func() app.UI { return &table.NestedColumnHeadersAndExpandableRows{} }
	result["NestedColumnHeaders"] = func() app.UI { return &table.NestedColumnHeaders{} }
	result["FavoritesExamples"] = func() app.UI { return &table.FavoritesExamples{} }
	result["FavoritesSortableExample"] = func() app.UI { return &table.FavoritesSortableExample{} }
	result["DraggableRowsExample"] = func() app.UI { return &table.DraggableRowsExample{} }
	result["StripedTableExample"] = func() app.UI { return &table.StripedTableExample{} }
	result["StripedExpandableTableExample"] = func() app.UI { return &table.StripedExpandableTableExample{} }
	result["StripedMultipleTbodyExample"] = func() app.UI { return &table.StripedMultipleTbodyExample{} }
	result["StripedTrExample"] = func() app.UI { return &table.StripedTrExample{} }
	return result
}

func TextArea() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &textarea.Basic{} }
	result["Invalid"] = func() app.UI { return &textarea.Invalid{} }
	result["Validated"] = func() app.UI { return &textarea.Validated{} }
	result["VerticallyResizableTextArea"] = func() app.UI { return &textarea.VerticallyResizableTextArea{} }
	result["HorizontallyResizableTextArea"] = func() app.UI { return &textarea.HorizontallyResizableTextArea{} }
	result["Uncontrolled"] = func() app.UI { return &textarea.Uncontrolled{} }
	result["Disabled"] = func() app.UI { return &textarea.Disabled{} }
	result["AutoResizing"] = func() app.UI { return &textarea.AutoResizing{} }
	result["IconSpriteVariants"] = func() app.UI { return &textarea.IconSpriteVariants{} }
	return result
}

func Tile() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["BasicTiles"] = func() app.UI { return &tile.BasicTiles{} }
	result["StackedTiles"] = func() app.UI { return &tile.StackedTiles{} }
	result["StackedTilesLarge"] = func() app.UI { return &tile.StackedTilesLarge{} }
	result["ExtraContent"] = func() app.UI { return &tile.ExtraContent{} }
	return result
}

func Button() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Variations"] = func() app.UI { return &button.Variations{} }
	result["Disabled"] = func() app.UI { return &button.Disabled{} }
	result["AriaDisabled"] = func() app.UI { return &button.AriaDisabled{} }
	result["LinksAsButtons"] = func() app.UI { return &button.LinksAsButtons{} }
	result["InlineLinkAsSpan"] = func() app.UI { return &button.InlineLinkAsSpan{} }
	result["BlockLevel"] = func() app.UI { return &button.BlockLevel{} }
	result["Types"] = func() app.UI { return &button.Types{} }
	result["CallToAction"] = func() app.UI { return &button.CallToAction{} }
	result["Progress"] = func() app.UI { return &button.Progress{} }
	return result
}

func Chip() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &chip.Basic{} }
	return result
}

func CodeBlock() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &codeblock.Basic{} }
	result["Expandable"] = func() app.UI { return &codeblock.Expandable{} }
	result["ExpandableExpanded"] = func() app.UI { return &codeblock.ExpandableExpanded{} }
	return result
}

func ClipboardCopy() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &clipboardcopy.Basic{} }
	result["Expandable"] = func() app.UI { return &clipboardcopy.Expandable{} }
	result["InlineCompact"] = func() app.UI { return &clipboardcopy.InlineCompact{} }
	result["InlineCompactCode"] = func() app.UI { return &clipboardcopy.InlineCompactCode{} }
	result["InlineCompactWithAdditionalAction"] = func() app.UI { return &clipboardcopy.InlineCompactWithAdditionalAction{} }
	result["InlineCompactInSentence"] = func() app.UI { return &clipboardcopy.InlineCompactInSentence{} }
	return result
}

func Drawer() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["ClosedPanelOnRightDefault"] = func() app.UI { return &drawer.ClosedPanelOnRightDefault{} }
	result["ExpandedPanelOnRight"] = func() app.UI { return &drawer.ExpandedPanelOnRight{} }
	result["ClosedPanelOnLeft"] = func() app.UI { return &drawer.ClosedPanelOnLeft{} }
	result["ExpandedPanelOnLeft"] = func() app.UI { return &drawer.ExpandedPanelOnLeft{} }
	result["ClosedPanelOnBottom"] = func() app.UI { return &drawer.ClosedPanelOnBottom{} }
	result["ExpandedPanelOnBottom"] = func() app.UI { return &drawer.ExpandedPanelOnBottom{} }
	result["ExpandedInlinePanel"] = func() app.UI { return &drawer.ExpandedInlinePanel{} }
	result["ExpandedInlinePanelOnLeft"] = func() app.UI { return &drawer.ExpandedInlinePanelOnLeft{} }
	result["Static"] = func() app.UI { return &drawer.Static{} }
	result["StackedContentBodyElements"] = func() app.UI { return &drawer.StackedContentBodyElements{} }
	result["ModifiedContentPadding"] = func() app.UI { return &drawer.ModifiedContentPadding{} }
	result["ModifiedPanelPadding"] = func() app.UI { return &drawer.ModifiedPanelPadding{} }
	result["ModifiedPanelWidth"] = func() app.UI { return &drawer.ModifiedPanelWidth{} }
	result["AdditionalSectionAboveMain"] = func() app.UI { return &drawer.AdditionalSectionAboveMain{} }
	result["ResizablePanel"] = func() app.UI { return &drawer.ResizablePanel{} }
	result["ResizableLeftPanel"] = func() app.UI { return &drawer.ResizableLeftPanel{} }
	result["ResizableBottomPanel"] = func() app.UI { return &drawer.ResizableBottomPanel{} }
	result["ResizableInlinePanel"] = func() app.UI { return &drawer.ResizableInlinePanel{} }
	result["PanelWithLight200Background"] = func() app.UI { return &drawer.PanelWithLight200Background{} }
	return result
}

func LabelGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &labelgroup.Basic{} }
	result["Overflow"] = func() app.UI { return &labelgroup.Overflow{} }
	result["OverflowExpanded"] = func() app.UI { return &labelgroup.OverflowExpanded{} }
	result["Category"] = func() app.UI { return &labelgroup.Category{} }
	result["CategoryRemovable"] = func() app.UI { return &labelgroup.CategoryRemovable{} }
	result["Vertical"] = func() app.UI { return &labelgroup.Vertical{} }
	result["VerticalOverflow"] = func() app.UI { return &labelgroup.VerticalOverflow{} }
	result["VerticalOverflowExpanded"] = func() app.UI { return &labelgroup.VerticalOverflowExpanded{} }
	result["VerticalCategory"] = func() app.UI { return &labelgroup.VerticalCategory{} }
	result["VerticalCategoryRemovable"] = func() app.UI { return &labelgroup.VerticalCategoryRemovable{} }
	result["EditableLabelsDynamicLabelGroup"] = func() app.UI { return &labelgroup.EditableLabelsDynamicLabelGroup{} }
	result["EditableLabelsLabelActiveDynamicLabelGroup"] = func() app.UI { return &labelgroup.EditableLabelsLabelActiveDynamicLabelGroup{} }
	result["StaticLabelsDynamicLabelGroup"] = func() app.UI { return &labelgroup.StaticLabelsDynamicLabelGroup{} }
	result["MixedLabelsStaticEditableDynamicLabelGroup"] = func() app.UI { return &labelgroup.MixedLabelsStaticEditableDynamicLabelGroup{} }
	result["CompactLabels"] = func() app.UI { return &labelgroup.CompactLabels{} }
	result["CompactLabelsOverflow"] = func() app.UI { return &labelgroup.CompactLabelsOverflow{} }
	result["CompactLabelsVertical"] = func() app.UI { return &labelgroup.CompactLabelsVertical{} }
	return result
}

func List() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Unordered"] = func() app.UI { return &list.Unordered{} }
	result["Ordered"] = func() app.UI { return &list.Ordered{} }
	result["Inline"] = func() app.UI { return &list.Inline{} }
	result["Plain"] = func() app.UI { return &list.Plain{} }
	result["WithHorizontalRules"] = func() app.UI { return &list.WithHorizontalRules{} }
	result["WithSmallIcons"] = func() app.UI { return &list.WithSmallIcons{} }
	result["WithLargeIcons"] = func() app.UI { return &list.WithLargeIcons{} }
	return result
}

func LoginPage() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &loginpage.Basic{} }
	result["Invalid"] = func() app.UI { return &loginpage.Invalid{} }
	result["ShowPassword"] = func() app.UI { return &loginpage.ShowPassword{} }
	result["HidePassword"] = func() app.UI { return &loginpage.HidePassword{} }
	result["WithLanguageSelector"] = func() app.UI { return &loginpage.WithLanguageSelector{} }
	return result
}

func Accordion() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Fluid"] = func() app.UI { return &accordion.Fluid{} }
	result["Fixed"] = func() app.UI { return &accordion.Fixed{} }
	result["DefinitionList"] = func() app.UI { return &accordion.DefinitionList{} }
	result["Bordered"] = func() app.UI { return &accordion.Bordered{} }
	result["LargeBordered"] = func() app.UI { return &accordion.LargeBordered{} }
	return result
}

func Alert() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Types"] = func() app.UI { return &alert.Types{} }
	result["Variations"] = func() app.UI { return &alert.Variations{} }
	result["InlineTypes"] = func() app.UI { return &alert.InlineTypes{} }
	result["InlineVariations"] = func() app.UI { return &alert.InlineVariations{} }
	result["CustomIcon"] = func() app.UI { return &alert.CustomIcon{} }
	result["InlinePlain"] = func() app.UI { return &alert.InlinePlain{} }
	result["Expandable"] = func() app.UI { return &alert.Expandable{} }
	return result
}

func Brand() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &brand.Basic{} }
	result["Responsive"] = func() app.UI { return &brand.Responsive{} }
	return result
}

func SimpleList() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SimpleList"] = func() app.UI { return &simplelist.SimpleList{} }
	result["SimpleListWithLinks"] = func() app.UI { return &simplelist.SimpleListWithLinks{} }
	result["GroupedList"] = func() app.UI { return &simplelist.GroupedList{} }
	return result
}

func Wizard() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &wizard.Basic{} }
	result["NavExpandedMobile"] = func() app.UI { return &wizard.NavExpandedMobile{} }
	result["WithDrawer"] = func() app.UI { return &wizard.WithDrawer{} }
	result["ExpandableCollapsed"] = func() app.UI { return &wizard.ExpandableCollapsed{} }
	result["ExpandableExpanded"] = func() app.UI { return &wizard.ExpandableExpanded{} }
	result["Finished"] = func() app.UI { return &wizard.Finished{} }
	return result
}

func Masthead() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &masthead.Basic{} }
	result["BasicWithMixedContent"] = func() app.UI { return &masthead.BasicWithMixedContent{} }
	result["DisplayInline"] = func() app.UI { return &masthead.DisplayInline{} }
	result["DisplayStack"] = func() app.UI { return &masthead.DisplayStack{} }
	result["DisplayStackDisplayInlineResponsive"] = func() app.UI { return &masthead.DisplayStackDisplayInlineResponsive{} }
	result["LightVariant"] = func() app.UI { return &masthead.LightVariant{} }
	result["Light200Variant"] = func() app.UI { return &masthead.Light200Variant{} }
	result["Insets"] = func() app.UI { return &masthead.Insets{} }
	return result
}

func Page() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["VerticalNav"] = func() app.UI { return &page.VerticalNav{} }
	result["HorizontalNav"] = func() app.UI { return &page.HorizontalNav{} }
	result["WithOrWithoutFill"] = func() app.UI { return &page.WithOrWithoutFill{} }
	result["MainSectionPadding"] = func() app.UI { return &page.MainSectionPadding{} }
	result["MainSectionVariations"] = func() app.UI { return &page.MainSectionVariations{} }
	result["CenteredSection"] = func() app.UI { return &page.CenteredSection{} }
	return result
}

func Select() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SingleSelect"] = func() app.UI { return &pfselect.SingleSelect{} }
	result["SingleExpanded"] = func() app.UI { return &pfselect.SingleExpanded{} }
	result["SingleWithTopExpanded"] = func() app.UI { return &pfselect.SingleWithTopExpanded{} }
	result["SingleExpandedAndSelected"] = func() app.UI { return &pfselect.SingleExpandedAndSelected{} }
	result["Disabled"] = func() app.UI { return &pfselect.Disabled{} }
	result["Success"] = func() app.UI { return &pfselect.Success{} }
	result["Warning"] = func() app.UI { return &pfselect.Warning{} }
	result["Invalid"] = func() app.UI { return &pfselect.Invalid{} }
	result["SingleWithTypeahead"] = func() app.UI { return &pfselect.SingleWithTypeahead{} }
	result["SingleWithTypeaheadExpanded"] = func() app.UI { return &pfselect.SingleWithTypeaheadExpanded{} }
	result["SingleWithTypeaheadExpandedAndSelected"] = func() app.UI { return &pfselect.SingleWithTypeaheadExpandedAndSelected{} }
	result["DisabledWithTypeahead"] = func() app.UI { return &pfselect.DisabledWithTypeahead{} }
	result["InvalidWithTypeahead"] = func() app.UI { return &pfselect.InvalidWithTypeahead{} }
	result["SelectMultiWithTypeahead"] = func() app.UI { return &pfselect.SelectMultiWithTypeahead{} }
	result["MultiWithTypeaheadChipGroupExpanded"] = func() app.UI { return &pfselect.MultiWithTypeaheadChipGroupExpanded{} }
	result["MultiWithTypeaheadChipGroupCollapsed"] = func() app.UI { return &pfselect.MultiWithTypeaheadChipGroupCollapsed{} }
	result["MultiWithTypeaheadInvalid"] = func() app.UI { return &pfselect.MultiWithTypeaheadInvalid{} }
	result["CheckboxSelect"] = func() app.UI { return &pfselect.CheckboxSelect{} }
	result["CheckboxExpanded"] = func() app.UI { return &pfselect.CheckboxExpanded{} }
	result["CheckboxExpandedAndSelectedWithGroups"] = func() app.UI { return &pfselect.CheckboxExpandedAndSelectedWithGroups{} }
	result["CheckboxExpandedAndSelectedWithGroupsAndFilter"] = func() app.UI { return &pfselect.CheckboxExpandedAndSelectedWithGroupsAndFilter{} }
	result["CheckboxExpandedWithoutBadge"] = func() app.UI { return &pfselect.CheckboxExpandedWithoutBadge{} }
	result["CheckboxWithCounts"] = func() app.UI { return &pfselect.CheckboxWithCounts{} }
	result["PlainToggle"] = func() app.UI { return &pfselect.PlainToggle{} }
	result["PlainToggleExpanded"] = func() app.UI { return &pfselect.PlainToggleExpanded{} }
	result["ToggleIcon"] = func() app.UI { return &pfselect.ToggleIcon{} }
	result["PanelMenu"] = func() app.UI { return &pfselect.PanelMenu{} }
	result["ItemDescription"] = func() app.UI { return &pfselect.ItemDescription{} }
	result["CheckboxItemDescription"] = func() app.UI { return &pfselect.CheckboxItemDescription{} }
	result["MenuItemFavorites"] = func() app.UI { return &pfselect.MenuItemFavorites{} }
	result["ViewMoreMenuItem"] = func() app.UI { return &pfselect.ViewMoreMenuItem{} }
	result["LoadingMenuItem"] = func() app.UI { return &pfselect.LoadingMenuItem{} }
	result["MenuFooter"] = func() app.UI { return &pfselect.MenuFooter{} }
	result["PlaceholderCollapsed"] = func() app.UI { return &pfselect.PlaceholderCollapsed{} }
	result["PlaceholderExpanded"] = func() app.UI { return &pfselect.PlaceholderExpanded{} }
	result["PlaceholderItemDisabled"] = func() app.UI { return &pfselect.PlaceholderItemDisabled{} }
	result["PlaceholderItemEnabled"] = func() app.UI { return &pfselect.PlaceholderItemEnabled{} }
	return result
}

func OverflowMenu() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SimpleCollapsed"] = func() app.UI { return &overflowmenu.SimpleCollapsed{} }
	result["SimpleExpanded"] = func() app.UI { return &overflowmenu.SimpleExpanded{} }
	result["GroupTypes"] = func() app.UI { return &overflowmenu.GroupTypes{} }
	result["AdditionalOptionsInDropdownHidden"] = func() app.UI { return &overflowmenu.AdditionalOptionsInDropdownHidden{} }
	result["AdditionalOptionsInDropdownVisible"] = func() app.UI { return &overflowmenu.AdditionalOptionsInDropdownVisible{} }
	result["PersistentAdditionalOptionsHidden"] = func() app.UI { return &overflowmenu.PersistentAdditionalOptionsHidden{} }
	result["PersistentAdditionalOptionsVisible"] = func() app.UI { return &overflowmenu.PersistentAdditionalOptionsVisible{} }
	return result
}

func ActionList() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["ActionListSingleGroup"] = func() app.UI { return &actionlist.ActionListSingleGroup{} }
	result["ActionListWithIcons"] = func() app.UI { return &actionlist.ActionListWithIcons{} }
	result["ActionListMultipleGroups"] = func() app.UI { return &actionlist.ActionListMultipleGroups{} }
	result["ActionListWithCancelButton"] = func() app.UI { return &actionlist.ActionListWithCancelButton{} }
	return result
}

func Banner() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &banner.Basic{} }
	result["BannerWithLinks"] = func() app.UI { return &banner.BannerWithLinks{} }
	return result
}

func DescriptionList() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &descriptionlist.Default{} }
	result["TermHelpText"] = func() app.UI { return &descriptionlist.TermHelpText{} }
	result["Default2Col"] = func() app.UI { return &descriptionlist.Default2Col{} }
	result["Default3ColOnLg"] = func() app.UI { return &descriptionlist.Default3ColOnLg{} }
	result["Horizontal"] = func() app.UI { return &descriptionlist.Horizontal{} }
	result["Horizontal2Col"] = func() app.UI { return &descriptionlist.Horizontal2Col{} }
	result["Horizontal3ColOnLg"] = func() app.UI { return &descriptionlist.Horizontal3ColOnLg{} }
	result["Compact"] = func() app.UI { return &descriptionlist.Compact{} }
	result["CompactHorizontal"] = func() app.UI { return &descriptionlist.CompactHorizontal{} }
	result["FluidHorizontal"] = func() app.UI { return &descriptionlist.FluidHorizontal{} }
	result["ColumnFill"] = func() app.UI { return &descriptionlist.ColumnFill{} }
	result["AutoFitBasic"] = func() app.UI { return &descriptionlist.AutoFitBasic{} }
	result["AutoFitMinWidthModifiedGridTemplateColumns"] = func() app.UI { return &descriptionlist.AutoFitMinWidthModifiedGridTemplateColumns{} }
	result["AutoFitMinWidthModifiedResponsiveGridTemplateColumns"] = func() app.UI { return &descriptionlist.AutoFitMinWidthModifiedResponsiveGridTemplateColumns{} }
	result["DefaultResponsiveColumns"] = func() app.UI { return &descriptionlist.DefaultResponsiveColumns{} }
	result["HorizontalResponsiveColumns"] = func() app.UI { return &descriptionlist.HorizontalResponsiveColumns{} }
	result["ResponsiveHorizontalVerticalGroupLayout"] = func() app.UI { return &descriptionlist.ResponsiveHorizontalVerticalGroupLayout{} }
	result["DefaultAutoColumnsWidth"] = func() app.UI { return &descriptionlist.DefaultAutoColumnsWidth{} }
	result["HorizontalAutoColumnWidth"] = func() app.UI { return &descriptionlist.HorizontalAutoColumnWidth{} }
	result["DefaultInlineGrid"] = func() app.UI { return &descriptionlist.DefaultInlineGrid{} }
	result["IconsOnTerms"] = func() app.UI { return &descriptionlist.IconsOnTerms{} }
	return result
}

func TextInputGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &textinputgroup.Basic{} }
	result["Disabled"] = func() app.UI { return &textinputgroup.Disabled{} }
	result["UtilitiesAndIconWithPlaceholderText"] = func() app.UI { return &textinputgroup.UtilitiesAndIconWithPlaceholderText{} }
	result["Filters"] = func() app.UI { return &textinputgroup.Filters{} }
	result["FiltersExpanded"] = func() app.UI { return &textinputgroup.FiltersExpanded{} }
	result["AutocompleteLastOptionHint"] = func() app.UI { return &textinputgroup.AutocompleteLastOptionHint{} }
	result["SearchInputGroup"] = func() app.UI { return &textinputgroup.SearchInputGroup{} }
	result["SearchInputGroupNoMatch"] = func() app.UI { return &textinputgroup.SearchInputGroupNoMatch{} }
	result["SearchInputGroupMatchWithResultCount"] = func() app.UI { return &textinputgroup.SearchInputGroupMatchWithResultCount{} }
	result["SearchInputGroupMatchWithNavigableOptions"] = func() app.UI { return &textinputgroup.SearchInputGroupMatchWithNavigableOptions{} }
	result["SearchInputGroupWithSubmitButton"] = func() app.UI { return &textinputgroup.SearchInputGroupWithSubmitButton{} }
	result["SearchInputGroupAdvancedSearch"] = func() app.UI { return &textinputgroup.SearchInputGroupAdvancedSearch{} }
	result["SearchInputGroupAdvancedSearchExpanded"] = func() app.UI { return &textinputgroup.SearchInputGroupAdvancedSearchExpanded{} }
	result["SearchInputGroupAutocomplete"] = func() app.UI { return &textinputgroup.SearchInputGroupAutocomplete{} }
	result["SearchInputGroupAutocompleteLastOptionHint"] = func() app.UI { return &textinputgroup.SearchInputGroupAutocompleteLastOptionHint{} }
	result["SearchInputGroupAdvancedSearchExpandedWithAutocomplete"] = func() app.UI { return &textinputgroup.SearchInputGroupAdvancedSearchExpandedWithAutocomplete{} }
	return result
}

func DataList() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &datalist.Basic{} }
	result["WithHeadings"] = func() app.UI { return &datalist.WithHeadings{} }
	result["CheckboxesActionsAndAdditionalCells"] = func() app.UI { return &datalist.CheckboxesActionsAndAdditionalCells{} }
	result["Expandable"] = func() app.UI { return &datalist.Expandable{} }
	result["ExpandableCompact"] = func() app.UI { return &datalist.ExpandableCompact{} }
	result["ExpandableNested"] = func() app.UI { return &datalist.ExpandableNested{} }
	result["Compact"] = func() app.UI { return &datalist.Compact{} }
	result["Modifiers"] = func() app.UI { return &datalist.Modifiers{} }
	result["SelectableRows"] = func() app.UI { return &datalist.SelectableRows{} }
	result["SelectableExpandableRows"] = func() app.UI { return &datalist.SelectableExpandableRows{} }
	result["Draggable"] = func() app.UI { return &datalist.Draggable{} }
	result["TextModifiers"] = func() app.UI { return &datalist.TextModifiers{} }
	result["TextModifiersDataListText"] = func() app.UI { return &datalist.TextModifiersDataListText{} }
	result["Grid"] = func() app.UI { return &datalist.Grid{} }
	result["GridSmallBreakpoint"] = func() app.UI { return &datalist.GridSmallBreakpoint{} }
	result["GridNone"] = func() app.UI { return &datalist.GridNone{} }
	return result
}

func Form() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["VerticallyAlignedLabels"] = func() app.UI { return &form.VerticallyAlignedLabels{} }
	result["HorizontallyAlignedLabels"] = func() app.UI { return &form.HorizontallyAlignedLabels{} }
	result["HorizontalLayoutAtACustomBreakpoint"] = func() app.UI { return &form.HorizontalLayoutAtACustomBreakpoint{} }
	result["FormSections"] = func() app.UI { return &form.FormSections{} }
	result["HelpText"] = func() app.UI { return &form.HelpText{} }
	result["LabelWithAdditionalInfo"] = func() app.UI { return &form.LabelWithAdditionalInfo{} }
	result["ActionGroup"] = func() app.UI { return &form.ActionGroup{} }
	result["FieldGroups"] = func() app.UI { return &form.FieldGroups{} }
	return result
}

func SkipToContent() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &skiptocontent.Basic{} }
	return result
}

func Popover() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Top"] = func() app.UI { return &popover.Top{} }
	result["Right"] = func() app.UI { return &popover.Right{} }
	result["Bottom"] = func() app.UI { return &popover.Bottom{} }
	result["Left"] = func() app.UI { return &popover.Left{} }
	result["LeftWithTopAndBottomPositions"] = func() app.UI { return &popover.LeftWithTopAndBottomPositions{} }
	result["BottomWithLeftAndRightPositions"] = func() app.UI { return &popover.BottomWithLeftAndRightPositions{} }
	result["WithoutHeaderfooter"] = func() app.UI { return &popover.WithoutHeaderfooter{} }
	result["NoPadding"] = func() app.UI { return &popover.NoPadding{} }
	result["WidthAuto"] = func() app.UI { return &popover.WidthAuto{} }
	result["PopoverWithIconInTheTitle"] = func() app.UI { return &popover.PopoverWithIconInTheTitle{} }
	result["DefaultAlertPopover"] = func() app.UI { return &popover.DefaultAlertPopover{} }
	result["InfoAlertPopover"] = func() app.UI { return &popover.InfoAlertPopover{} }
	result["SuccessAlertPopover"] = func() app.UI { return &popover.SuccessAlertPopover{} }
	result["WarningAlertPopover"] = func() app.UI { return &popover.WarningAlertPopover{} }
	result["DangerAlertPopover"] = func() app.UI { return &popover.DangerAlertPopover{} }
	return result
}

func Radio() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &radio.Basic{} }
	result["Checked"] = func() app.UI { return &radio.Checked{} }
	result["LabelWrappingInput"] = func() app.UI { return &radio.LabelWrappingInput{} }
	result["Reversed"] = func() app.UI { return &radio.Reversed{} }
	result["Disabled"] = func() app.UI { return &radio.Disabled{} }
	result["WithDescription"] = func() app.UI { return &radio.WithDescription{} }
	result["WithBody"] = func() app.UI { return &radio.WithBody{} }
	result["WithDescriptionAndBody"] = func() app.UI { return &radio.WithDescriptionAndBody{} }
	result["StandaloneInput"] = func() app.UI { return &radio.StandaloneInput{} }
	return result
}

func Truncate() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &truncate.Default{} }
	result["Middle"] = func() app.UI { return &truncate.Middle{} }
	result["Start"] = func() app.UI { return &truncate.Start{} }
	return result
}

func FormControl() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Input"] = func() app.UI { return &formcontrol.Input{} }
	result["Select"] = func() app.UI { return &formcontrol.Select{} }
	result["Textarea"] = func() app.UI { return &formcontrol.Textarea{} }
	result["IconSprite"] = func() app.UI { return &formcontrol.IconSprite{} }
	return result
}

func Progress() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Simple"] = func() app.UI { return &progress.Simple{} }
	result["Small"] = func() app.UI { return &progress.Small{} }
	result["Large"] = func() app.UI { return &progress.Large{} }
	result["Outside"] = func() app.UI { return &progress.Outside{} }
	result["Inside"] = func() app.UI { return &progress.Inside{} }
	result["Success"] = func() app.UI { return &progress.Success{} }
	result["Warning"] = func() app.UI { return &progress.Warning{} }
	result["Failure"] = func() app.UI { return &progress.Failure{} }
	result["InsideSuccess"] = func() app.UI { return &progress.InsideSuccess{} }
	result["InsideWarning"] = func() app.UI { return &progress.InsideWarning{} }
	result["OutsideFailure"] = func() app.UI { return &progress.OutsideFailure{} }
	result["OutsideStaticWidthMeasure"] = func() app.UI { return &progress.OutsideStaticWidthMeasure{} }
	result["OnSingleLine"] = func() app.UI { return &progress.OnSingleLine{} }
	result["WithoutMeasure"] = func() app.UI { return &progress.WithoutMeasure{} }
	result["FailureWithoutMeasure"] = func() app.UI { return &progress.FailureWithoutMeasure{} }
	result["FiniteStep"] = func() app.UI { return &progress.FiniteStep{} }
	result["TruncateDescription"] = func() app.UI { return &progress.TruncateDescription{} }
	result["ProgressStepInstruction"] = func() app.UI { return &progress.ProgressStepInstruction{} }
	return result
}

func Avatar() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &avatar.Basic{} }
	result["BorderedLight"] = func() app.UI { return &avatar.BorderedLight{} }
	result["BorderedDark"] = func() app.UI { return &avatar.BorderedDark{} }
	result["Small"] = func() app.UI { return &avatar.Small{} }
	result["Medium"] = func() app.UI { return &avatar.Medium{} }
	result["Large"] = func() app.UI { return &avatar.Large{} }
	result["ExtraLarge"] = func() app.UI { return &avatar.ExtraLarge{} }
	return result
}

func Card() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &card.Basic{} }
	result["WithImageAndAction"] = func() app.UI { return &card.WithImageAndAction{} }
	result["WithTitleInHead"] = func() app.UI { return &card.WithTitleInHead{} }
	result["WithOnlyActionsInHeadNoTitlefooter"] = func() app.UI { return &card.WithOnlyActionsInHeadNoTitlefooter{} }
	result["ActionsWithNoOffset"] = func() app.UI { return &card.ActionsWithNoOffset{} }
	result["WithOnlyImageInHead"] = func() app.UI { return &card.WithOnlyImageInHead{} }
	result["WithNoFooter"] = func() app.UI { return &card.WithNoFooter{} }
	result["WithNoTitle"] = func() app.UI { return &card.WithNoTitle{} }
	result["WithOnlyAContentSection"] = func() app.UI { return &card.WithOnlyAContentSection{} }
	result["WithMultipleBodySections"] = func() app.UI { return &card.WithMultipleBodySections{} }
	result["WithOnlyOneBodyThatFills"] = func() app.UI { return &card.WithOnlyOneBodyThatFills{} }
	result["Compact"] = func() app.UI { return &card.Compact{} }
	result["Large"] = func() app.UI { return &card.Large{} }
	result["Hoverable"] = func() app.UI { return &card.Hoverable{} }
	result["Selectable"] = func() app.UI { return &card.Selectable{} }
	result["Selected"] = func() app.UI { return &card.Selected{} }
	result["NonSelectable"] = func() app.UI { return &card.NonSelectable{} }
	result["HoverableLegacy"] = func() app.UI { return &card.HoverableLegacy{} }
	result["SelectableLegacy"] = func() app.UI { return &card.SelectableLegacy{} }
	result["SelectedLegacy"] = func() app.UI { return &card.SelectedLegacy{} }
	result["Flat"] = func() app.UI { return &card.Flat{} }
	result["Rounded"] = func() app.UI { return &card.Rounded{} }
	result["Plain"] = func() app.UI { return &card.Plain{} }
	result["Expandable"] = func() app.UI { return &card.Expandable{} }
	result["ExpandableWithImage"] = func() app.UI { return &card.ExpandableWithImage{} }
	result["Expanded"] = func() app.UI { return &card.Expanded{} }
	result["FullHeightCard"] = func() app.UI { return &card.FullHeightCard{} }
	result["ExpandableToggleOnRight"] = func() app.UI { return &card.ExpandableToggleOnRight{} }
	result["CardWithDividersBetweenSections"] = func() app.UI { return &card.CardWithDividersBetweenSections{} }
	return result
}

func ChipGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SimpleInlineChipGroupOverflow"] = func() app.UI { return &chipgroup.SimpleInlineChipGroupOverflow{} }
	result["SimpleInlineChipGroupExpanded"] = func() app.UI { return &chipgroup.SimpleInlineChipGroupExpanded{} }
	result["ChipGroupWithCategories"] = func() app.UI { return &chipgroup.ChipGroupWithCategories{} }
	result["ChipGroupWithCategoriesOverflow"] = func() app.UI { return &chipgroup.ChipGroupWithCategoriesOverflow{} }
	result["ChipGroupWithCategoriesOverflowExpanded"] = func() app.UI { return &chipgroup.ChipGroupWithCategoriesOverflowExpanded{} }
	result["ChipGroupWithCategoriesRemovable"] = func() app.UI { return &chipgroup.ChipGroupWithCategoriesRemovable{} }
	result["LegacyChipGroupExamplesWithoutMainElement"] = func() app.UI { return &chipgroup.LegacyChipGroupExamplesWithoutMainElement{} }
	return result
}

func ToggleGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &togglegroup.Default{} }
	result["WithIcon"] = func() app.UI { return &togglegroup.WithIcon{} }
	result["IconAndText"] = func() app.UI { return &togglegroup.IconAndText{} }
	result["Compact"] = func() app.UI { return &togglegroup.Compact{} }
	return result
}

func Menu() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &menu.Basic{} }
	result["WithIcons"] = func() app.UI { return &menu.WithIcons{} }
	result["WithCheckbox"] = func() app.UI { return &menu.WithCheckbox{} }
	result["Scrollable"] = func() app.UI { return &menu.Scrollable{} }
	result["ScrollableWithCustomMenuHeight"] = func() app.UI { return &menu.ScrollableWithCustomMenuHeight{} }
	result["WithFlyout"] = func() app.UI { return &menu.WithFlyout{} }
	result["WithFlyoutMenuTop"] = func() app.UI { return &menu.WithFlyoutMenuTop{} }
	result["WithFlyoutMenuLeft"] = func() app.UI { return &menu.WithFlyoutMenuLeft{} }
	result["WithFlyoutMenuLeftTop"] = func() app.UI { return &menu.WithFlyoutMenuLeftTop{} }
	result["StandardMenuFlyoutChild"] = func() app.UI { return &menu.StandardMenuFlyoutChild{} }
	result["Drilldown"] = func() app.UI { return &menu.Drilldown{} }
	result["DrilldownLevelTwo"] = func() app.UI { return &menu.DrilldownLevelTwo{} }
	result["DrilldownLevelThree"] = func() app.UI { return &menu.DrilldownLevelThree{} }
	result["DrilldownLevelFour"] = func() app.UI { return &menu.DrilldownLevelFour{} }
	result["ScrollableDrilldown"] = func() app.UI { return &menu.ScrollableDrilldown{} }
	result["WidthModifiedDrilldown"] = func() app.UI { return &menu.WidthModifiedDrilldown{} }
	result["DrilldownWithBreadcrumbsLevel1"] = func() app.UI { return &menu.DrilldownWithBreadcrumbsLevel1{} }
	result["DrilldownWithBreadcrumbsLevel2"] = func() app.UI { return &menu.DrilldownWithBreadcrumbsLevel2{} }
	result["DrilldownWithBreadcrumbsLevel3"] = func() app.UI { return &menu.DrilldownWithBreadcrumbsLevel3{} }
	result["DrilldownWithBreadcrumbsLevel4"] = func() app.UI { return &menu.DrilldownWithBreadcrumbsLevel4{} }
	result["ScrollableMenuWithHeaderAndFooter"] = func() app.UI { return &menu.ScrollableMenuWithHeaderAndFooter{} }
	result["ScrollableMenuWithSearchAndFooter"] = func() app.UI { return &menu.ScrollableMenuWithSearchAndFooter{} }
	result["WithFiltering"] = func() app.UI { return &menu.WithFiltering{} }
	result["WithLinks"] = func() app.UI { return &menu.WithLinks{} }
	result["WithSeparators"] = func() app.UI { return &menu.WithSeparators{} }
	result["WithTitledGroups"] = func() app.UI { return &menu.WithTitledGroups{} }
	result["WithDescription"] = func() app.UI { return &menu.WithDescription{} }
	result["WithActions"] = func() app.UI { return &menu.WithActions{} }
	result["WithFavorites"] = func() app.UI { return &menu.WithFavorites{} }
	result["OptionSingleSelect"] = func() app.UI { return &menu.OptionSingleSelect{} }
	result["OptionMultiSelect"] = func() app.UI { return &menu.OptionMultiSelect{} }
	result["ViewMore"] = func() app.UI { return &menu.ViewMore{} }
	result["Loading"] = func() app.UI { return &menu.Loading{} }
	result["Footer"] = func() app.UI { return &menu.Footer{} }
	result["Plain"] = func() app.UI { return &menu.Plain{} }
	result["PlainWithSearchAndFooter"] = func() app.UI { return &menu.PlainWithSearchAndFooter{} }
	result["PlainScrollableWithSearchAndFooter"] = func() app.UI { return &menu.PlainScrollableWithSearchAndFooter{} }
	return result
}

func Pagination() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Top"] = func() app.UI { return &pagination.Top{} }
	result["TopExpanded"] = func() app.UI { return &pagination.TopExpanded{} }
	result["TopSticky"] = func() app.UI { return &pagination.TopSticky{} }
	result["IndeterminateItemCountIsNotKnown"] = func() app.UI { return &pagination.IndeterminateItemCountIsNotKnown{} }
	result["Bottom"] = func() app.UI { return &pagination.Bottom{} }
	result["BottomSticky"] = func() app.UI { return &pagination.BottomSticky{} }
	result["TopDisabled"] = func() app.UI { return &pagination.TopDisabled{} }
	result["Compact"] = func() app.UI { return &pagination.Compact{} }
	result["TopWithDisplaySummaryModifier"] = func() app.UI { return &pagination.TopWithDisplaySummaryModifier{} }
	result["TopWithDisplayFullModifier"] = func() app.UI { return &pagination.TopWithDisplayFullModifier{} }
	result["TopWithResponsiveDisplaySummaryAndDisplayFullModifiers"] = func() app.UI { return &pagination.TopWithResponsiveDisplaySummaryAndDisplayFullModifiers{} }
	result["CompactDisplayFullModifier"] = func() app.UI { return &pagination.CompactDisplayFullModifier{} }
	return result
}

func ProgressStepper() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &progressstepper.Basic{} }
	result["BasicWithDescriptions"] = func() app.UI { return &progressstepper.BasicWithDescriptions{} }
	result["CenterAlignedWithDescriptions"] = func() app.UI { return &progressstepper.CenterAlignedWithDescriptions{} }
	result["VerticalWithDescriptions"] = func() app.UI { return &progressstepper.VerticalWithDescriptions{} }
	result["Compact"] = func() app.UI { return &progressstepper.Compact{} }
	result["BasicWithAnIssue"] = func() app.UI { return &progressstepper.BasicWithAnIssue{} }
	result["BasicWithAFailure"] = func() app.UI { return &progressstepper.BasicWithAFailure{} }
	result["BasicWithPatternflyIcons"] = func() app.UI { return &progressstepper.BasicWithPatternflyIcons{} }
	result["WithHelpText"] = func() app.UI { return &progressstepper.WithHelpText{} }
	return result
}

func SearchInput() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &searchinput.Basic{} }
	result["NoMatch"] = func() app.UI { return &searchinput.NoMatch{} }
	result["MatchWithResultCount"] = func() app.UI { return &searchinput.MatchWithResultCount{} }
	result["MatchWithNavigableOptions"] = func() app.UI { return &searchinput.MatchWithNavigableOptions{} }
	result["WithSubmitButton"] = func() app.UI { return &searchinput.WithSubmitButton{} }
	result["AdvancedSearch"] = func() app.UI { return &searchinput.AdvancedSearch{} }
	result["AdvancedSearchExpanded"] = func() app.UI { return &searchinput.AdvancedSearchExpanded{} }
	result["Autocomplete"] = func() app.UI { return &searchinput.Autocomplete{} }
	result["AutocompleteLastOptionHint"] = func() app.UI { return &searchinput.AutocompleteLastOptionHint{} }
	result["AdvancedSearchExpandedWithAutocomplete"] = func() app.UI { return &searchinput.AdvancedSearchExpandedWithAutocomplete{} }
	return result
}

func Skeleton() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &skeleton.Default{} }
	result["PercentageWidthModifiers"] = func() app.UI { return &skeleton.PercentageWidthModifiers{} }
	result["PercentageHeightModifiers"] = func() app.UI { return &skeleton.PercentageHeightModifiers{} }
	result["TextModifiers"] = func() app.UI { return &skeleton.TextModifiers{} }
	result["StaticHeightWidthAndShapeModifiers"] = func() app.UI { return &skeleton.StaticHeightWidthAndShapeModifiers{} }
	return result
}

func Tooltip() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Top"] = func() app.UI { return &tooltip.Top{} }
	result["Right"] = func() app.UI { return &tooltip.Right{} }
	result["Bottom"] = func() app.UI { return &tooltip.Bottom{} }
	result["Left"] = func() app.UI { return &tooltip.Left{} }
	result["LeftWithTopAndBottomPositions"] = func() app.UI { return &tooltip.LeftWithTopAndBottomPositions{} }
	result["BottomWithLeftAndRightPositions"] = func() app.UI { return &tooltip.BottomWithLeftAndRightPositions{} }
	result["LeftAlignedText"] = func() app.UI { return &tooltip.LeftAlignedText{} }
	return result
}

func CalendarMonth() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["DateSelected"] = func() app.UI { return &calendarmonth.DateSelected{} }
	result["RangeStartDateSelectedEndDateHovered"] = func() app.UI { return &calendarmonth.RangeStartDateSelectedEndDateHovered{} }
	result["RangeEndDateSelectedStartDateFocused"] = func() app.UI { return &calendarmonth.RangeEndDateSelectedStartDateFocused{} }
	result["RangeStartAndEndDatesSelected"] = func() app.UI { return &calendarmonth.RangeStartAndEndDatesSelected{} }
	return result
}

func FormSelect() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &formselect.Basic{} }
	result["Invalid"] = func() app.UI { return &formselect.Invalid{} }
	result["Validated"] = func() app.UI { return &formselect.Validated{} }
	result["Disabled"] = func() app.UI { return &formselect.Disabled{} }
	result["Grouped"] = func() app.UI { return &formselect.Grouped{} }
	result["IconSpriteVariant"] = func() app.UI { return &formselect.IconSpriteVariant{} }
	return result
}

func Modal() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &modal.Basic{} }
	result["WithHelpButton"] = func() app.UI { return &modal.WithHelpButton{} }
	result["Small"] = func() app.UI { return &modal.Small{} }
	result["Medium"] = func() app.UI { return &modal.Medium{} }
	result["Large"] = func() app.UI { return &modal.Large{} }
	result["WithoutTitle"] = func() app.UI { return &modal.WithoutTitle{} }
	result["WithDescription"] = func() app.UI { return &modal.WithDescription{} }
	result["CustomTitle"] = func() app.UI { return &modal.CustomTitle{} }
	result["ModalBoxAsGenericContainer"] = func() app.UI { return &modal.ModalBoxAsGenericContainer{} }
	result["Icon"] = func() app.UI { return &modal.Icon{} }
	result["DefaultAlert"] = func() app.UI { return &modal.DefaultAlert{} }
	result["InfoAlert"] = func() app.UI { return &modal.InfoAlert{} }
	result["SuccessAlert"] = func() app.UI { return &modal.SuccessAlert{} }
	result["WarningAlert"] = func() app.UI { return &modal.WarningAlert{} }
	result["DangerAlert"] = func() app.UI { return &modal.DangerAlert{} }
	return result
}

func Spinner() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &spinner.Basic{} }
	result["Sizes"] = func() app.UI { return &spinner.Sizes{} }
	result["CustomSize"] = func() app.UI { return &spinner.CustomSize{} }
	result["BasicLegacy"] = func() app.UI { return &spinner.BasicLegacy{} }
	result["SizesLegacy"] = func() app.UI { return &spinner.SizesLegacy{} }
	return result
}

func Badge() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Read"] = func() app.UI { return &badge.Read{} }
	result["Unread"] = func() app.UI { return &badge.Unread{} }
	return result
}

func Breadcrumb() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &breadcrumb.Basic{} }
	result["WithoutHomeLink"] = func() app.UI { return &breadcrumb.WithoutHomeLink{} }
	result["WithHeading"] = func() app.UI { return &breadcrumb.WithHeading{} }
	result["WithDropdown"] = func() app.UI { return &breadcrumb.WithDropdown{} }
	result["WithButtons"] = func() app.UI { return &breadcrumb.WithButtons{} }
	return result
}

func OptionsMenu() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SingleOption"] = func() app.UI { return &optionsmenu.SingleOption{} }
	result["Disabled"] = func() app.UI { return &optionsmenu.Disabled{} }
	result["MultipleOptions"] = func() app.UI { return &optionsmenu.MultipleOptions{} }
	result["Plain"] = func() app.UI { return &optionsmenu.Plain{} }
	result["AlignTop"] = func() app.UI { return &optionsmenu.AlignTop{} }
	result["AlignRight"] = func() app.UI { return &optionsmenu.AlignRight{} }
	result["PlainWithText"] = func() app.UI { return &optionsmenu.PlainWithText{} }
	result["WithGroups"] = func() app.UI { return &optionsmenu.WithGroups{} }
	result["WithGroupsAndDividersBetweenGroups"] = func() app.UI { return &optionsmenu.WithGroupsAndDividersBetweenGroups{} }
	result["WithGroupsAndDividersBetweenItems"] = func() app.UI { return &optionsmenu.WithGroupsAndDividersBetweenItems{} }
	return result
}

func NotificationBadge() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &notificationbadge.Basic{} }
	result["WithCount"] = func() app.UI { return &notificationbadge.WithCount{} }
	return result
}

func NotificationDrawer() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &notificationdrawer.Basic{} }
	result["Groups"] = func() app.UI { return &notificationdrawer.Groups{} }
	return result
}

func Sidebar() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &sidebar.Basic{} }
	result["Gutter"] = func() app.UI { return &sidebar.Gutter{} }
	result["Stack"] = func() app.UI { return &sidebar.Stack{} }
	result["Split"] = func() app.UI { return &sidebar.Split{} }
	result["PanelRightHtml"] = func() app.UI { return &sidebar.PanelRightHtml{} }
	result["PanelRightModifier"] = func() app.UI { return &sidebar.PanelRightModifier{} }
	result["StickyPanel"] = func() app.UI { return &sidebar.StickyPanel{} }
	result["StaticPanel"] = func() app.UI { return &sidebar.StaticPanel{} }
	result["ResponsivePanelWidth"] = func() app.UI { return &sidebar.ResponsivePanelWidth{} }
	return result
}

func Switch() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &pfswitch.Basic{} }
	result["ReverseToggleOnRight"] = func() app.UI { return &pfswitch.ReverseToggleOnRight{} }
	result["LabelAndCheck"] = func() app.UI { return &pfswitch.LabelAndCheck{} }
	result["WithoutLabel"] = func() app.UI { return &pfswitch.WithoutLabel{} }
	result["Disabled"] = func() app.UI { return &pfswitch.Disabled{} }
	return result
}

func Title() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["SizeModifiers"] = func() app.UI { return &title.SizeModifiers{} }
	return result
}

func AppLauncher() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Collapsed"] = func() app.UI { return &applauncher.Collapsed{} }
	result["Disabled"] = func() app.UI { return &applauncher.Disabled{} }
	result["Expanded"] = func() app.UI { return &applauncher.Expanded{} }
	result["AlignedRight"] = func() app.UI { return &applauncher.AlignedRight{} }
	result["AlignedTop"] = func() app.UI { return &applauncher.AlignedTop{} }
	result["WithSectionsAndDividersBetweenSections"] = func() app.UI { return &applauncher.WithSectionsAndDividersBetweenSections{} }
	result["WithSectionsAndDividersBetweenItems"] = func() app.UI { return &applauncher.WithSectionsAndDividersBetweenItems{} }
	result["WithSectionsDividersIconsAndExternalLinks"] = func() app.UI { return &applauncher.WithSectionsDividersIconsAndExternalLinks{} }
	result["Favorites"] = func() app.UI { return &applauncher.Favorites{} }
	return result
}

func Checkbox() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &checkbox.Basic{} }
	result["Checked"] = func() app.UI { return &checkbox.Checked{} }
	result["LabelWrappingInput"] = func() app.UI { return &checkbox.LabelWrappingInput{} }
	result["Reversed"] = func() app.UI { return &checkbox.Reversed{} }
	result["Disabled"] = func() app.UI { return &checkbox.Disabled{} }
	result["WithDescription"] = func() app.UI { return &checkbox.WithDescription{} }
	result["WithBody"] = func() app.UI { return &checkbox.WithBody{} }
	result["WithDescriptionAndBody"] = func() app.UI { return &checkbox.WithDescriptionAndBody{} }
	result["StandaloneInput"] = func() app.UI { return &checkbox.StandaloneInput{} }
	return result
}

func Dropdown() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Expanded"] = func() app.UI { return &dropdown.Expanded{} }
	result["Collapsed"] = func() app.UI { return &dropdown.Collapsed{} }
	result["Disabled"] = func() app.UI { return &dropdown.Disabled{} }
	result["AriaDisabledItems"] = func() app.UI { return &dropdown.AriaDisabledItems{} }
	result["Kebab"] = func() app.UI { return &dropdown.Kebab{} }
	result["KebabAlignRight"] = func() app.UI { return &dropdown.KebabAlignRight{} }
	result["AlignRight"] = func() app.UI { return &dropdown.AlignRight{} }
	result["AlignOnDifferentBreakpoint"] = func() app.UI { return &dropdown.AlignOnDifferentBreakpoint{} }
	result["AlignTop"] = func() app.UI { return &dropdown.AlignTop{} }
	result["PlainWithText"] = func() app.UI { return &dropdown.PlainWithText{} }
	result["BadgeToggle"] = func() app.UI { return &dropdown.BadgeToggle{} }
	result["MenuItemIcons"] = func() app.UI { return &dropdown.MenuItemIcons{} }
	result["SplitButtonCheckbox"] = func() app.UI { return &dropdown.SplitButtonCheckbox{} }
	result["SplitButtonCheckboxWithToggleText"] = func() app.UI { return &dropdown.SplitButtonCheckboxWithToggleText{} }
	result["SplitButtonAction"] = func() app.UI { return &dropdown.SplitButtonAction{} }
	result["SplitButtonPrimaryAction"] = func() app.UI { return &dropdown.SplitButtonPrimaryAction{} }
	result["WithGroups"] = func() app.UI { return &dropdown.WithGroups{} }
	result["WithGroupsAndDividersBetweenGroups"] = func() app.UI { return &dropdown.WithGroupsAndDividersBetweenGroups{} }
	result["WithGroupsAndDividersBetweenItems"] = func() app.UI { return &dropdown.WithGroupsAndDividersBetweenItems{} }
	result["Panel"] = func() app.UI { return &dropdown.Panel{} }
	result["PrimaryToggle"] = func() app.UI { return &dropdown.PrimaryToggle{} }
	result["SecondaryToggle"] = func() app.UI { return &dropdown.SecondaryToggle{} }
	result["DropdownWithImageAndText"] = func() app.UI { return &dropdown.DropdownWithImageAndText{} }
	result["DropdownWithDescription"] = func() app.UI { return &dropdown.DropdownWithDescription{} }
	return result
}

func Toolbar() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Simple"] = func() app.UI { return &toolbar.Simple{} }
	result["AdjustedSpacers"] = func() app.UI { return &toolbar.AdjustedSpacers{} }
	result["AdjustedGroupSpacers"] = func() app.UI { return &toolbar.AdjustedGroupSpacers{} }
	result["Insets"] = func() app.UI { return &toolbar.Insets{} }
	result["PageInsets"] = func() app.UI { return &toolbar.PageInsets{} }
	result["WidthControl"] = func() app.UI { return &toolbar.WidthControl{} }
	result["GroupTypes"] = func() app.UI { return &toolbar.GroupTypes{} }
	result["ToggleGroup"] = func() app.UI { return &toolbar.ToggleGroup{} }
	result["ToggleGroupOnMobileFiltersCollapsedExpandableContentExpanded"] = func() app.UI { return &toolbar.ToggleGroupOnMobileFiltersCollapsedExpandableContentExpanded{} }
	result["SelectedFiltersOnMobileFiltersCollapsedSelectedFiltersSummaryVisible"] = func() app.UI { return &toolbar.SelectedFiltersOnMobileFiltersCollapsedSelectedFiltersSummaryVisible{} }
	result["SelectedFiltersOnMobileFiltersCollapsedExpandableContentExpanded"] = func() app.UI { return &toolbar.SelectedFiltersOnMobileFiltersCollapsedExpandableContentExpanded{} }
	result["SelectedFiltersOnDesktopNotResponsive"] = func() app.UI { return &toolbar.SelectedFiltersOnDesktopNotResponsive{} }
	result["StackedOnDesktop"] = func() app.UI { return &toolbar.StackedOnDesktop{} }
	result["StackedOnMobileFiltersCollapsedExpandableContentExpanded"] = func() app.UI { return &toolbar.StackedOnMobileFiltersCollapsedExpandableContentExpanded{} }
	result["ExpandedElements"] = func() app.UI { return &toolbar.ExpandedElements{} }
	return result
}

func BackToTop() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &backtotop.Basic{} }
	return result
}

func FileUploadMultiple() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &fileuploadmultiple.Basic{} }
	result["DragOver"] = func() app.UI { return &fileuploadmultiple.DragOver{} }
	result["Horizontal"] = func() app.UI { return &fileuploadmultiple.Horizontal{} }
	result["HorizontalDragOver"] = func() app.UI { return &fileuploadmultiple.HorizontalDragOver{} }
	result["FileUploadStatus"] = func() app.UI { return &fileuploadmultiple.FileUploadStatus{} }
	result["FileUploadStatusExpanded"] = func() app.UI { return &fileuploadmultiple.FileUploadStatusExpanded{} }
	result["HorizontalFileUploadStatusExpanded"] = func() app.UI { return &fileuploadmultiple.HorizontalFileUploadStatusExpanded{} }
	return result
}

func TabContent() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &tabcontent.Basic{} }
	result["Padding"] = func() app.UI { return &tabcontent.Padding{} }
	result["Light300Background"] = func() app.UI { return &tabcontent.Light300Background{} }
	return result
}

func DualListSelector() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &duallistselector.Basic{} }
	result["AvailableItemSelected"] = func() app.UI { return &duallistselector.AvailableItemSelected{} }
	result["MultipleAvailableItemsSelected"] = func() app.UI { return &duallistselector.MultipleAvailableItemsSelected{} }
	result["ChosenItem"] = func() app.UI { return &duallistselector.ChosenItem{} }
	result["ChosenItemSelected"] = func() app.UI { return &duallistselector.ChosenItemSelected{} }
	result["TreeView"] = func() app.UI { return &duallistselector.TreeView{} }
	result["TreeViewWithChosenAndDisabledOptions"] = func() app.UI { return &duallistselector.TreeViewWithChosenAndDisabledOptions{} }
	result["Draggable"] = func() app.UI { return &duallistselector.Draggable{} }
	return result
}

func EmptyState() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &emptystate.Basic{} }
	result["ExtraSmall"] = func() app.UI { return &emptystate.ExtraSmall{} }
	result["Small"] = func() app.UI { return &emptystate.Small{} }
	result["Large"] = func() app.UI { return &emptystate.Large{} }
	result["ExtraLarge"] = func() app.UI { return &emptystate.ExtraLarge{} }
	result["WithPrimaryElement"] = func() app.UI { return &emptystate.WithPrimaryElement{} }
	return result
}

func ExpandableSection() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Hidden"] = func() app.UI { return &expandablesection.Hidden{} }
	result["Expanded"] = func() app.UI { return &expandablesection.Expanded{} }
	result["DisclosureVariationHidden"] = func() app.UI { return &expandablesection.DisclosureVariationHidden{} }
	result["DisclosureVariationExpanded"] = func() app.UI { return &expandablesection.DisclosureVariationExpanded{} }
	result["DetachedToggle"] = func() app.UI { return &expandablesection.DetachedToggle{} }
	result["Indented"] = func() app.UI { return &expandablesection.Indented{} }
	return result
}

func Hint() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["HintWithTitle"] = func() app.UI { return &hint.HintWithTitle{} }
	result["DefaultWithNoTitle"] = func() app.UI { return &hint.DefaultWithNoTitle{} }
	return result
}

func TextInput() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &textinput.Basic{} }
	result["Disabled"] = func() app.UI { return &textinput.Disabled{} }
	result["TruncatedOnLeft"] = func() app.UI { return &textinput.TruncatedOnLeft{} }
	result["ReadOnly"] = func() app.UI { return &textinput.ReadOnly{} }
	result["Invalid"] = func() app.UI { return &textinput.Invalid{} }
	result["SelectTextUsingRef"] = func() app.UI { return &textinput.SelectTextUsingRef{} }
	result["IconVariants"] = func() app.UI { return &textinput.IconVariants{} }
	result["IconSpriteVariants"] = func() app.UI { return &textinput.IconSpriteVariants{} }
	return result
}

func AlertGroup() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["StaticAlertGroup"] = func() app.UI { return &alertgroup.StaticAlertGroup{} }
	result["ToastAlertGroup"] = func() app.UI { return &alertgroup.ToastAlertGroup{} }
	return result
}

func BackgroundImage() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &backgroundimage.Basic{} }
	return result
}

func DragAndDrop() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &draganddrop.Basic{} }
	result["Dragging"] = func() app.UI { return &draganddrop.Dragging{} }
	result["DragOutside"] = func() app.UI { return &draganddrop.DragOutside{} }
	return result
}

func FileUpload() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["BasicFileUpload"] = func() app.UI { return &fileupload.BasicFileUpload{} }
	result["UploadCompleteNonEditable"] = func() app.UI { return &fileupload.UploadCompleteNonEditable{} }
	result["UploadCompleteEditable"] = func() app.UI { return &fileupload.UploadCompleteEditable{} }
	result["DragFileHoverComponent"] = func() app.UI { return &fileupload.DragFileHoverComponent{} }
	result["FileUploadInFormWithError"] = func() app.UI { return &fileupload.FileUploadInFormWithError{} }
	result["FileUploadLoading"] = func() app.UI { return &fileupload.FileUploadLoading{} }
	return result
}

func Label() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Filled"] = func() app.UI { return &label.Filled{} }
	result["Outline"] = func() app.UI { return &label.Outline{} }
	result["Compact"] = func() app.UI { return &label.Compact{} }
	result["Overflow"] = func() app.UI { return &label.Overflow{} }
	result["Editable"] = func() app.UI { return &label.Editable{} }
	return result
}

func Panel() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &panel.Basic{} }
	result["Header"] = func() app.UI { return &panel.Header{} }
	result["Footer"] = func() app.UI { return &panel.Footer{} }
	result["HeaderAndFooter"] = func() app.UI { return &panel.HeaderAndFooter{} }
	result["NoBody"] = func() app.UI { return &panel.NoBody{} }
	result["Raised"] = func() app.UI { return &panel.Raised{} }
	result["Bordered"] = func() app.UI { return &panel.Bordered{} }
	result["Scrollable"] = func() app.UI { return &panel.Scrollable{} }
	result["ScrollableWithHeaderAndFooter"] = func() app.UI { return &panel.ScrollableWithHeaderAndFooter{} }
	return result
}

func HelperText() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Static"] = func() app.UI { return &helpertext.Static{} }
	result["Icon"] = func() app.UI { return &helpertext.Icon{} }
	result["MultipleStatic"] = func() app.UI { return &helpertext.MultipleStatic{} }
	result["Dynamic"] = func() app.UI { return &helpertext.Dynamic{} }
	result["DynamicList"] = func() app.UI { return &helpertext.DynamicList{} }
	return result
}

func MenuToggle() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Collapsed"] = func() app.UI { return &menutoggle.Collapsed{} }
	result["Expanded"] = func() app.UI { return &menutoggle.Expanded{} }
	result["Disabled"] = func() app.UI { return &menutoggle.Disabled{} }
	result["Count"] = func() app.UI { return &menutoggle.Count{} }
	result["Primary"] = func() app.UI { return &menutoggle.Primary{} }
	result["Secondary"] = func() app.UI { return &menutoggle.Secondary{} }
	result["Plain"] = func() app.UI { return &menutoggle.Plain{} }
	result["PlainWithText"] = func() app.UI { return &menutoggle.PlainWithText{} }
	result["SplitButtonCheckbox"] = func() app.UI { return &menutoggle.SplitButtonCheckbox{} }
	result["SplitButtonCheckboxWithToggleText"] = func() app.UI { return &menutoggle.SplitButtonCheckboxWithToggleText{} }
	result["SplitButtonPrimary"] = func() app.UI { return &menutoggle.SplitButtonPrimary{} }
	result["SplitButtonSecondary"] = func() app.UI { return &menutoggle.SplitButtonSecondary{} }
	result["SplitButtonAction"] = func() app.UI { return &menutoggle.SplitButtonAction{} }
	result["SplitButtonPrimaryAction"] = func() app.UI { return &menutoggle.SplitButtonPrimaryAction{} }
	result["SplitButtonSecondaryAction"] = func() app.UI { return &menutoggle.SplitButtonSecondaryAction{} }
	result["WithIconimageAndText"] = func() app.UI { return &menutoggle.WithIconimageAndText{} }
	result["WithAvatarAndText"] = func() app.UI { return &menutoggle.WithAvatarAndText{} }
	result["FullHeight"] = func() app.UI { return &menutoggle.FullHeight{} }
	result["Typeahead"] = func() app.UI { return &menutoggle.Typeahead{} }
	return result
}

func Navigation() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &navigation.Default{} }
	result["GroupedNav"] = func() app.UI { return &navigation.GroupedNav{} }
	result["GroupedNavNoTitles"] = func() app.UI { return &navigation.GroupedNavNoTitles{} }
	result["GroupedNavNoTitlesNoMarginTop"] = func() app.UI { return &navigation.GroupedNavNoTitlesNoMarginTop{} }
	result["Expanded"] = func() app.UI { return &navigation.Expanded{} }
	result["ExpandedWithSubnavTitles"] = func() app.UI { return &navigation.ExpandedWithSubnavTitles{} }
	result["Mixed"] = func() app.UI { return &navigation.Mixed{} }
	result["ExpandableThirdLevel"] = func() app.UI { return &navigation.ExpandableThirdLevel{} }
	result["Horizontal"] = func() app.UI { return &navigation.Horizontal{} }
	result["HorizontalOverflow"] = func() app.UI { return &navigation.HorizontalOverflow{} }
	result["HorizontalSubnav"] = func() app.UI { return &navigation.HorizontalSubnav{} }
	result["HorizontalSubnavOverflow"] = func() app.UI { return &navigation.HorizontalSubnavOverflow{} }
	result["LegacyTertiary"] = func() app.UI { return &navigation.LegacyTertiary{} }
	result["LegacyTertiaryOverflow"] = func() app.UI { return &navigation.LegacyTertiaryOverflow{} }
	result["DefaultInLightMode"] = func() app.UI { return &navigation.DefaultInLightMode{} }
	result["ExpandedInLightMode"] = func() app.UI { return &navigation.ExpandedInLightMode{} }
	result["NavWithFlyout"] = func() app.UI { return &navigation.NavWithFlyout{} }
	result["NavWithDrilldownMenu"] = func() app.UI { return &navigation.NavWithDrilldownMenu{} }
	result["NavWithDrilldownMenuLevelTwo"] = func() app.UI { return &navigation.NavWithDrilldownMenuLevelTwo{} }
	result["NavWithDrilldownMenuLevelThree"] = func() app.UI { return &navigation.NavWithDrilldownMenuLevelThree{} }
	return result
}

func Slider() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Discrete"] = func() app.UI { return &slider.Discrete{} }
	result["Continuous"] = func() app.UI { return &slider.Continuous{} }
	result["ValueInput"] = func() app.UI { return &slider.ValueInput{} }
	result["ThumbValueInput"] = func() app.UI { return &slider.ThumbValueInput{} }
	result["Actions"] = func() app.UI { return &slider.Actions{} }
	result["Disabled"] = func() app.UI { return &slider.Disabled{} }
	return result
}

func Tabs() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &tabs.Default{} }
	result["DefaultOverflowBeginningOfList"] = func() app.UI { return &tabs.DefaultOverflowBeginningOfList{} }
	result["Vertical"] = func() app.UI { return &tabs.Vertical{} }
	result["Box"] = func() app.UI { return &tabs.Box{} }
	result["BoxOverflow"] = func() app.UI { return &tabs.BoxOverflow{} }
	result["BoxVertical"] = func() app.UI { return &tabs.BoxVertical{} }
	result["BoxTabsColorSchemeLight300"] = func() app.UI { return &tabs.BoxTabsColorSchemeLight300{} }
	result["Inset"] = func() app.UI { return &tabs.Inset{} }
	result["InsetBox"] = func() app.UI { return &tabs.InsetBox{} }
	result["PageInsets"] = func() app.UI { return &tabs.PageInsets{} }
	result["IconsAndText"] = func() app.UI { return &tabs.IconsAndText{} }
	result["TabsWithSubTabs"] = func() app.UI { return &tabs.TabsWithSubTabs{} }
	result["BoxTabsWithSubTabs"] = func() app.UI { return &tabs.BoxTabsWithSubTabs{} }
	result["Filled"] = func() app.UI { return &tabs.Filled{} }
	result["FilledWithIcons"] = func() app.UI { return &tabs.FilledWithIcons{} }
	result["FilledBox"] = func() app.UI { return &tabs.FilledBox{} }
	result["FilledBoxWithIcons"] = func() app.UI { return &tabs.FilledBoxWithIcons{} }
	result["UsingTheNavElement"] = func() app.UI { return &tabs.UsingTheNavElement{} }
	result["SubNavUsingTheNavElement"] = func() app.UI { return &tabs.SubNavUsingTheNavElement{} }
	result["VerticalExpandable"] = func() app.UI { return &tabs.VerticalExpandable{} }
	result["VerticalExpanded"] = func() app.UI { return &tabs.VerticalExpanded{} }
	result["VerticalExpandableResponsive"] = func() app.UI { return &tabs.VerticalExpandableResponsive{} }
	result["VerticalExpandableLegacy"] = func() app.UI { return &tabs.VerticalExpandableLegacy{} }
	result["VerticalExpandedLegacy"] = func() app.UI { return &tabs.VerticalExpandedLegacy{} }
	result["VerticalExpandableResponsiveLegacy"] = func() app.UI { return &tabs.VerticalExpandableResponsiveLegacy{} }
	result["CloseButton"] = func() app.UI { return &tabs.CloseButton{} }
	result["AddTabButton"] = func() app.UI { return &tabs.AddTabButton{} }
	return result
}

func AboutModal() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &aboutmodal.Basic{} }
	return result
}

func CodeEditor() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Default"] = func() app.UI { return &codeeditor.Default{} }
	result["ReadOnly"] = func() app.UI { return &codeeditor.ReadOnly{} }
	result["WithoutActions"] = func() app.UI { return &codeeditor.WithoutActions{} }
	result["DragFileAndHoverOverComponent"] = func() app.UI { return &codeeditor.DragFileAndHoverOverComponent{} }
	result["WithOptionalHeaderContentAndKeyboardShortcuts"] = func() app.UI { return &codeeditor.WithOptionalHeaderContentAndKeyboardShortcuts{} }
	return result
}

func DatePicker() map[string]func() app.UI {
	result := make(map[string]func() app.UI)
	result["Basic"] = func() app.UI { return &datepicker.Basic{} }
	result["HelperText"] = func() app.UI { return &datepicker.HelperText{} }
	result["Invalid"] = func() app.UI { return &datepicker.Invalid{} }
	result["Expanded"] = func() app.UI { return &datepicker.Expanded{} }
	result["CustomWidthInput"] = func() app.UI { return &datepicker.CustomWidthInput{} }
	result["CustomWidthInputBasedOnNumberOfCharacters"] = func() app.UI { return &datepicker.CustomWidthInputBasedOnNumberOfCharacters{} }
	return result
}

func AllComponents() []func() map[string]func() app.UI {
	var result []func() map[string]func() app.UI
	result = append(result, AboutModal)
	result = append(result, Accordion)
	result = append(result, ActionList)
	result = append(result, Alert)
	result = append(result, AlertGroup)
	result = append(result, AppLauncher)
	result = append(result, Avatar)
	result = append(result, BackToTop)
	result = append(result, Backdrop)
	result = append(result, BackgroundImage)
	result = append(result, Badge)
	result = append(result, Banner)
	result = append(result, Brand)
	result = append(result, Breadcrumb)
	result = append(result, Button)
	result = append(result, CalendarMonth)
	result = append(result, Card)
	result = append(result, Checkbox)
	result = append(result, Chip)
	result = append(result, ChipGroup)
	result = append(result, ClipboardCopy)
	result = append(result, CodeBlock)
	result = append(result, CodeEditor)
	result = append(result, ContextSelector)
	result = append(result, DataList)
	result = append(result, DatePicker)
	result = append(result, DescriptionList)
	result = append(result, Divider)
	result = append(result, DragAndDrop)
	result = append(result, Drawer)
	result = append(result, Dropdown)
	result = append(result, DualListSelector)
	result = append(result, EmptyState)
	result = append(result, ExpandableSection)
	result = append(result, FileUpload)
	result = append(result, FileUploadMultiple)
	result = append(result, Form)
	result = append(result, FormControl)
	result = append(result, FormSelect)
	result = append(result, HelperText)
	result = append(result, Hint)
	result = append(result, InlineEdit)
	result = append(result, InputGroup)
	result = append(result, JumpLinks)
	result = append(result, Label)
	result = append(result, LabelGroup)
	result = append(result, List)
	result = append(result, LoginPage)
	result = append(result, Masthead)
	result = append(result, Menu)
	result = append(result, MenuToggle)
	result = append(result, Modal)
	result = append(result, Navigation)
	result = append(result, NotificationBadge)
	result = append(result, NotificationDrawer)
	result = append(result, NumberInput)
	result = append(result, OptionsMenu)
	result = append(result, OverflowMenu)
	result = append(result, Page)
	result = append(result, Pagination)
	result = append(result, Panel)
	result = append(result, Popover)
	result = append(result, Progress)
	result = append(result, ProgressStepper)
	result = append(result, Radio)
	result = append(result, SearchInput)
	result = append(result, Select)
	result = append(result, Sidebar)
	result = append(result, SimpleList)
	result = append(result, Skeleton)
	result = append(result, SkipToContent)
	result = append(result, Slider)
	result = append(result, Spinner)
	result = append(result, Switch)
	result = append(result, TabContent)
	result = append(result, Table)
	result = append(result, Tabs)
	result = append(result, Text)
	result = append(result, TextArea)
	result = append(result, TextInput)
	result = append(result, TextInputGroup)
	result = append(result, Tile)
	result = append(result, TimePicker)
	result = append(result, Title)
	result = append(result, ToggleGroup)
	result = append(result, Toolbar)
	result = append(result, Tooltip)
	result = append(result, TreeView)
	result = append(result, Truncate)
	result = append(result, Wizard)
	return result
}
