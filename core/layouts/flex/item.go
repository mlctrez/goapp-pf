package flex

// from file "react-core/src/layouts/Flex/FlexItem.tsx"

type ItemProps struct {
	// Children - content rendered inside the Flex layout. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Flex layout. Optional.
	ClassName string
	// Spacer - Spacers at various breakpoints. Optional.
	Spacer map[string]string /* { default:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl }, sm:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl }, md:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl }, lg:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl }, xl:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl }, 2xl:{ spacerNone, spacerXs, spacerSm, spacerMd, spacerLg, spacerXl, spacer2xl, spacer3xl, spacer4xl } } */
	// Grow - Whether to add flex: grow at various breakpoints. Optional.
	Grow map[string]string /* { 'grow', 'grow', 'grow', 'grow', 'grow', 'grow' } */
	// Shrink - Whether to add flex: shrink at various breakpoints. Optional.
	Shrink map[string]string /* { 'shrink', 'shrink', 'shrink', 'shrink', 'shrink', 'shrink' } */
	// Flex - Value to add for flex property at various breakpoints. Optional.
	Flex map[string]string /* { default:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 }, sm:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 }, md:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 }, lg:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 }, xl:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 }, 2xl:{ flexDefault, flexNone, flex_1, flex_2, flex_3, flex_4 } } */
	// AlignSelf - Value to add for align-self property at various breakpoints. Optional.
	AlignSelf map[string]string /* { default:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline }, sm:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline }, md:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline }, lg:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline }, xl:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline }, 2xl:{ alignSelfFlexStart, alignSelfFlexEnd, alignSelfCenter, alignSelfStretch, alignSelfBaseline } } */
	// Align - Value to use for margin: auto at various breakpoints. Optional.
	Align map[string]string /* { default:{ alignLeft, alignRight }, sm:{ alignLeft, alignRight }, md:{ alignLeft, alignRight }, lg:{ alignLeft, alignRight }, xl:{ alignLeft, alignRight }, 2xl:{ alignLeft, alignRight } } */
	// FullWidth - Whether to set width: 100%!a(MISSING)t various breakpoints. Optional.
	FullWidth map[string]string /* { 'fullWidth', 'fullWidth', 'fullWidth', 'fullWidth', 'fullWidth', 'fullWidth' } */
	// Order - Modifies the flex layout element order property. Optional.
	Order map[string]string /* { default:string, md:string, lg:string, xl:string, 2xl:string } */
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
}
