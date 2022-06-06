package tabs

type TabTitleIconProps struct {
	// Children - Icon to be rendered inside the tab button title.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the tab title icon. Optional.
	ClassName string
}

// contents of react-core/src/components/Tabs/TabTitleIcon.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Tabs/tabs';
// 
// export interface TabTitleIconProps extends React.HTMLProps<HTMLSpanElement> {
//   /** Icon to be rendered inside the tab button title. */
//   children: React.ReactNode;
//   /** additional classes added to the tab title icon */
//   className?: string;
// }
// 
// export const TabTitleIcon: React.FunctionComponent<TabTitleIconProps> = ({
//   children,
//   className = '',
//   ...props
// }: TabTitleIconProps) => (
//   <span className={css(styles.tabsItemIcon, className)} {...props}>
//     {children}
//   </span>
// );
// TabTitleIcon.displayName = 'TabTitleIcon';
