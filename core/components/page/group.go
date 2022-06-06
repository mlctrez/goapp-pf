package page

type GroupProps struct {
	// ClassName - Additional classes to apply to the PageGroup. Optional.
	ClassName string
	// Children - Content rendered inside of the PageGroup. Optional.
	Children any //  // React.ReactNode 
	// Sticky - Modifier indicating if PageGroup is sticky to the top or bottom. Optional.
	Sticky any //  /* "top" | "bottom" */
	// HasShadowTop - Modifier indicating if PageGroup should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Modifier indicating if PageGroup should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageGroup has a scrolling overflow. Optional.
	HasOverflowScroll bool
}

// contents of react-core/src/components/Page/PageGroup.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Page/page';
// 
// export interface PageGroupProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes to apply to the PageGroup */
//   className?: string;
//   /** Content rendered inside of the PageGroup */
//   children?: React.ReactNode;
//   /** Modifier indicating if PageGroup is sticky to the top or bottom */
//   sticky?: 'top' | 'bottom';
//   /** Modifier indicating if PageGroup should have a shadow at the top */
//   hasShadowTop?: boolean;
//   /** Modifier indicating if PageGroup should have a shadow at the bottom */
//   hasShadowBottom?: boolean;
//   /** Flag indicating if the PageGroup has a scrolling overflow */
//   hasOverflowScroll?: boolean;
// }
// 
// export const PageGroup = ({
//   className = '',
//   children,
//   sticky,
//   hasShadowTop = false,
//   hasShadowBottom = false,
//   hasOverflowScroll = false,
//   ...props
// }: PageGroupProps) => (
//   <div
//     {...props}
//     className={css(
//       styles.pageMainGroup,
//       sticky === 'top' && styles.modifiers.stickyTop,
//       sticky === 'bottom' && styles.modifiers.stickyBottom,
//       hasShadowTop && styles.modifiers.shadowTop,
//       hasShadowBottom && styles.modifiers.shadowBottom,
//       hasOverflowScroll && styles.modifiers.overflowScroll,
//       className
//     )}
//     {...(hasOverflowScroll && { tabIndex: 0 })}
//   >
//     {children}
//   </div>
// );
// PageGroup.displayName = 'PageGroup';
