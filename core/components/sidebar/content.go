package sidebar

type ContentProps struct {
	// Children - 
	Children any //  // React.ReactNode 
	// HasNoBackground - Removes the background color. Optional.
	HasNoBackground bool
}

// contents of react-core/src/components/Sidebar/SidebarContent.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Sidebar/sidebar';
// 
// export interface SidebarContentProps extends React.HTMLProps<HTMLDivElement> {
//   children: React.ReactNode;
//   /** Removes the background color. */
//   hasNoBackground?: boolean;
// }
// 
// export const SidebarContent: React.FunctionComponent<SidebarContentProps> = ({
//   className,
//   children,
//   hasNoBackground,
//   ...props
// }: SidebarContentProps) => (
//   <div className={css(styles.sidebarContent, hasNoBackground && styles.modifiers.noBackground, className)} {...props}>
//     {children}
//   </div>
// );
// SidebarContent.displayName = 'SidebarContent';
