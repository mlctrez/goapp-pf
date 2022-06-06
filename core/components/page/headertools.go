package page

type HeaderToolsProps struct {
	// Children - Content rendered in page header tools.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the page header tools. Optional.
	ClassName string
}

// contents of react-core/src/components/Page/PageHeaderTools.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Page/page';
// import { css } from '@patternfly/react-styles';
// 
// export interface PageHeaderToolsProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered in page header tools */
//   children: React.ReactNode;
//   /** Additional classes added to the page header tools. */
//   className?: string;
// }
// 
// export const PageHeaderTools: React.FunctionComponent<PageHeaderToolsProps> = ({
//   children,
//   className,
//   ...props
// }: PageHeaderToolsProps) => (
//   <div className={css(styles.pageHeaderTools, className)} {...props}>
//     {children}
//   </div>
// );
// PageHeaderTools.displayName = 'PageHeaderTools';
