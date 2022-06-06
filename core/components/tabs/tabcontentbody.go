package tabs

type TabContentBodyProps struct {
	// Children - Content rendered inside the tab content body.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the tab content body. Optional.
	ClassName string
	// HasPadding - Indicates if there should be padding around the tab content body. Optional.
	HasPadding bool
}

// contents of react-core/src/components/Tabs/TabContentBody.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/TabContent/tab-content';
// 
// export interface TabContentBodyProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the tab content body. */
//   children: React.ReactNode;
//   /** Additional classes added to the tab content body. */
//   className?: string;
//   /** Indicates if there should be padding around the tab content body */
//   hasPadding?: boolean;
// }
// 
// export const TabContentBody: React.FunctionComponent<TabContentBodyProps> = ({
//   children,
//   className,
//   hasPadding,
//   ...props
// }: TabContentBodyProps) => (
//   <div className={css(styles.tabContentBody, hasPadding && styles.modifiers.padding, className)} {...props}>
//     {children}
//   </div>
// );
// TabContentBody.displayName = 'TabContentBody';
