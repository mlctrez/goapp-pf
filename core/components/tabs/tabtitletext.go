package tabs

type TabTitleTextProps struct {
	// Children - Text to be rendered inside the tab button title.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the tab title text. Optional.
	ClassName string
}

// contents of react-core/src/components/Tabs/TabTitleText.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Tabs/tabs';
// 
// export interface TabTitleTextProps extends React.HTMLProps<HTMLSpanElement> {
//   /** Text to be rendered inside the tab button title. */
//   children: React.ReactNode;
//   /** additional classes added to the tab title text */
//   className?: string;
// }
// 
// export const TabTitleText: React.FunctionComponent<TabTitleTextProps> = ({
//   children,
//   className = '',
//   ...props
// }: TabTitleTextProps) => (
//   <span className={css(styles.tabsItemText, className)} {...props}>
//     {children}
//   </span>
// );
// TabTitleText.displayName = 'TabTitleText';
