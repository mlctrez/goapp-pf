package masthead

// from file "react-core/src/components/Masthead/Masthead.tsx"

type Props struct {
	// Children - Content rendered inside of the masthead. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the masthead. Optional.
	ClassName string
	// BackgroundColor - Background theme color of the masthead. Optional.
	BackgroundColor any /* "dark" | "light" | "light200" */
	// Display - Display type at various breakpoints. Optional.
	Display map[string]string /* { default:{ inline, stack }, sm:{ inline, stack }, md:{ inline, stack }, lg:{ inline, stack }, xl:{ inline, stack }, 2xl:{ inline, stack } } */
	// Inset - Insets at various breakpoints. Optional.
	Inset map[string]string /* { default:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, sm:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, md:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, lg:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, xl:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, 2xl:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl } } */
}
