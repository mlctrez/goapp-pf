package divider

// from file "react-core/src/components/Divider/Divider.tsx"

type Props struct {
	// ClassName - Additional classes added to the divider. Optional.
	ClassName string
	// Component - The component type to use. Optional.
	Component any /* "hr" | "li" | "div" */
	// IsVertical - Optional.
	IsVertical bool
	// Inset - Insets at various breakpoints. Optional.
	Inset map[string]string /* { default:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, sm:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, md:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, lg:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, xl:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl }, 2xl:{ insetNone, insetXs, insetSm, insetMd, insetLg, insetXl, inset2xl, inset3xl } } */
	// Orientation - Indicates how the divider will display at various breakpoints. Vertical divider must be
	// in a flex layout. Optional.
	Orientation map[string]string /* { default:{ vertical, horizontal }, sm:{ vertical, horizontal }, md:{ vertical, horizontal }, lg:{ vertical, horizontal }, xl:{ vertical, horizontal }, 2xl:{ vertical, horizontal } } */
}
