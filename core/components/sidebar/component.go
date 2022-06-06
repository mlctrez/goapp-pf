package sidebar

type Props struct {
	// Children - Optional.
	Children any //  // React.ReactNode 
	// Orientation - Indicates the direction of the layout. Default orientation is stack on small screens, and
	// split on medium screens and above. Optional.
	Orientation any //  /* "stack" | "split" */
	// IsPanelRight - Indicates that the panel is displayed to the right of the content when the oritentation
	// is split. Optional.
	IsPanelRight bool
	// HasGutter - Adds space between the panel and content. Optional.
	HasGutter bool
	// HasNoBackground - Removes the background color. Optional.
	HasNoBackground bool
}

// contents of react-core/src/components/Sidebar/Sidebar.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Sidebar/sidebar';
// 
// export interface SidebarProps extends React.HTMLProps<HTMLDivElement> {
//   children?: React.ReactNode;
//   /** Indicates the direction of the layout. Default orientation is stack on small screens, and split on medium screens and above. */
//   orientation?: 'stack' | 'split';
//   /** Indicates that the panel is displayed to the right of the content when the oritentation is split. */
//   isPanelRight?: boolean;
//   /** Adds space between the panel and content. */
//   hasGutter?: boolean;
//   /** Removes the background color. */
//   hasNoBackground?: boolean;
// }
// 
// export const Sidebar: React.FunctionComponent<SidebarProps> = ({
//   className,
//   children,
//   orientation,
//   isPanelRight = false,
//   hasGutter,
//   hasNoBackground,
//   ...props
// }: SidebarProps) => (
//   <div
//     className={css(
//       styles.sidebar,
//       hasGutter && styles.modifiers.gutter,
//       hasNoBackground && styles.modifiers.noBackground,
//       isPanelRight && styles.modifiers.panelRight,
//       styles.modifiers[orientation],
//       className
//     )}
//     {...props}
//   >
//     <div className={styles.sidebarMain}>{children}</div>
//   </div>
// );
// Sidebar.displayName = 'Sidebar';
