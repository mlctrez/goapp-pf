package page

type SectionProps struct {
	// Children - Content rendered inside the section. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// Variant - Section background color variant. Optional.
	Variant any //  /* "default" | "light" | "dark" | "darker" */
	// Type - Section type variant. Optional.
	Type any //  /* "default" | "nav" | "subnav" | "breadcrumb" | "tabs" | "wizard" */
	// IsFilled - Enables the page section to fill the available vertical space. Optional.
	IsFilled bool
	// IsWidthLimited - Limits the width of the section. Optional.
	IsWidthLimited bool
	// IsCenterAligned - Flag indicating if the section content is center aligned. isWidthLimited must be set
	// for this to work. Optional.
	IsCenterAligned bool
	// Padding - Padding at various breakpoints. Optional.
	Padding map[string]string /* { default:{ padding, noPadding }, sm:{ padding, noPadding }, md:{ padding, noPadding }, lg:{ padding, noPadding }, xl:{ padding, noPadding }, 2xl:{ padding, noPadding } } */
	// Sticky - Modifier indicating if PageSection is sticky to the top or bottom. Optional.
	Sticky any //  /* "top" | "bottom" */
	// HasShadowTop - Modifier indicating if PageSection should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Modifier indicating if PageSection should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageSection has a scrolling overflow. Optional.
	HasOverflowScroll bool
}

// contents of react-core/src/components/Page/PageSection.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Page/page';
// import { css } from '@patternfly/react-styles';
// import { formatBreakpointMods } from '../../helpers/util';
// 
// export enum PageSectionVariants {
//   default = 'default',
//   light = 'light',
//   dark = 'dark',
//   darker = 'darker'
// }
// 
// export enum PageSectionTypes {
//   default = 'default',
//   nav = 'nav',
//   subNav = 'subnav',
//   breadcrumb = 'breadcrumb',
//   tabs = 'tabs',
//   wizard = 'wizard'
// }
// 
// export interface PageSectionProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the section */
//   children?: React.ReactNode;
//   /** Additional classes added to the section */
//   className?: string;
//   /** Section background color variant */
//   variant?: 'default' | 'light' | 'dark' | 'darker';
//   /** Section type variant */
//   type?: 'default' | 'nav' | 'subnav' | 'breadcrumb' | 'tabs' | 'wizard';
//   /** Enables the page section to fill the available vertical space */
//   isFilled?: boolean;
//   /** Limits the width of the section */
//   isWidthLimited?: boolean;
//   /** Flag indicating if the section content is center aligned. isWidthLimited must be set for this to work  */
//   isCenterAligned?: boolean;
//   /** Padding at various breakpoints. */
//   padding?: {
//     default?: 'padding' | 'noPadding';
//     sm?: 'padding' | 'noPadding';
//     md?: 'padding' | 'noPadding';
//     lg?: 'padding' | 'noPadding';
//     xl?: 'padding' | 'noPadding';
//     '2xl'?: 'padding' | 'noPadding';
//   };
//   /** Modifier indicating if PageSection is sticky to the top or bottom */
//   sticky?: 'top' | 'bottom';
//   /** Modifier indicating if PageSection should have a shadow at the top */
//   hasShadowTop?: boolean;
//   /** Modifier indicating if PageSection should have a shadow at the bottom */
//   hasShadowBottom?: boolean;
//   /** Flag indicating if the PageSection has a scrolling overflow */
//   hasOverflowScroll?: boolean;
// }
// 
// const variantType = {
//   [PageSectionTypes.default]: styles.pageMainSection,
//   [PageSectionTypes.nav]: styles.pageMainNav,
//   [PageSectionTypes.subNav]: styles.pageMainSubnav,
//   [PageSectionTypes.breadcrumb]: styles.pageMainBreadcrumb,
//   [PageSectionTypes.tabs]: styles.pageMainTabs,
//   [PageSectionTypes.wizard]: styles.pageMainWizard
// };
// 
// const variantStyle = {
//   [PageSectionVariants.default]: '',
//   [PageSectionVariants.light]: styles.modifiers.light,
//   [PageSectionVariants.dark]: styles.modifiers.dark_200,
//   [PageSectionVariants.darker]: styles.modifiers.dark_100
// };
// 
// export const PageSection: React.FunctionComponent<PageSectionProps> = ({
//   className = '',
//   children,
//   variant = 'default',
//   type = 'default',
//   padding,
//   isFilled,
//   isWidthLimited = false,
//   isCenterAligned = false,
//   sticky,
//   hasShadowTop = false,
//   hasShadowBottom = false,
//   hasOverflowScroll = false,
//   ...props
// }: PageSectionProps) => (
//   <section
//     {...props}
//     className={css(
//       variantType[type],
//       formatBreakpointMods(padding, styles),
//       variantStyle[variant],
//       isFilled === false && styles.modifiers.noFill,
//       isFilled === true && styles.modifiers.fill,
//       isWidthLimited && styles.modifiers.limitWidth,
//       isWidthLimited && isCenterAligned && type !== PageSectionTypes.subNav && styles.modifiers.alignCenter,
//       sticky === 'top' && styles.modifiers.stickyTop,
//       sticky === 'bottom' && styles.modifiers.stickyBottom,
//       hasShadowTop && styles.modifiers.shadowTop,
//       hasShadowBottom && styles.modifiers.shadowBottom,
//       hasOverflowScroll && styles.modifiers.overflowScroll,
//       className
//     )}
//     {...(hasOverflowScroll && { tabIndex: 0 })}
//   >
//     {isWidthLimited && <div className={css(styles.pageMainBody)}>{children}</div>}
//     {!isWidthLimited && children}
//   </section>
// );
// PageSection.displayName = 'PageSection';
