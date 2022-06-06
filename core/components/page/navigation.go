package page

type NavigationProps struct {
	// ClassName - Additional classes to apply to the PageNavigation. Optional.
	ClassName string
	// Children - Content rendered inside of the PageNavigation. Optional.
	Children any //  // React.ReactNode 
	// IsWidthLimited - Limits the width of the PageNavigation. Optional.
	IsWidthLimited bool
	// Sticky - Modifier indicating if the PageNavigation is sticky to the top or bottom. Optional.
	Sticky any //  /* "top" | "bottom" */
	// HasShadowTop - Flag indicating if PageNavigation should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Flag indicating if PageNavigation should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageNavigation has a scrolling overflow. Optional.
	HasOverflowScroll bool
}

// contents of react-core/src/components/Page/PageNavigation.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Page/page';
// 
// export interface PageNavigationProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes to apply to the PageNavigation */
//   className?: string;
//   /** Content rendered inside of the PageNavigation */
//   children?: React.ReactNode;
//   /** Limits the width of the PageNavigation */
//   isWidthLimited?: boolean;
//   /** Modifier indicating if the PageNavigation is sticky to the top or bottom */
//   sticky?: 'top' | 'bottom';
//   /** Flag indicating if PageNavigation should have a shadow at the top */
//   hasShadowTop?: boolean;
//   /** Flag indicating if PageNavigation should have a shadow at the bottom */
//   hasShadowBottom?: boolean;
//   /** Flag indicating if the PageNavigation has a scrolling overflow */
//   hasOverflowScroll?: boolean;
// }
// 
// export const PageNavigation = ({
//   className = '',
//   children,
//   isWidthLimited,
//   sticky,
//   hasShadowTop = false,
//   hasShadowBottom = false,
//   hasOverflowScroll = false,
//   ...props
// }: PageNavigationProps) => (
//   <div
//     className={css(
//       styles.pageMainNav,
//       isWidthLimited && styles.modifiers.limitWidth,
//       sticky === 'top' && styles.modifiers.stickyTop,
//       sticky === 'bottom' && styles.modifiers.stickyBottom,
//       hasShadowTop && styles.modifiers.shadowTop,
//       hasShadowBottom && styles.modifiers.shadowBottom,
//       hasOverflowScroll && styles.modifiers.overflowScroll,
//       className
//     )}
//     {...(hasOverflowScroll && { tabIndex: 0 })}
//     {...props}
//   >
//     {isWidthLimited && <div className={css(styles.pageMainBody)}>{children}</div>}
//     {!isWidthLimited && children}
//   </div>
// );
// PageNavigation.displayName = 'PageNavigation';
