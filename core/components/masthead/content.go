package masthead

type ContentProps struct {
	// Children - Content rendered inside of the masthead content block. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the masthead content. Optional.
	ClassName string
}

// contents of react-core/src/components/Masthead/MastheadContent.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Masthead/masthead';
// import { css } from '@patternfly/react-styles';
// 
// export interface MastheadContentProps extends React.DetailedHTMLProps<React.HTMLProps<HTMLDivElement>, HTMLDivElement> {
//   /** Content rendered inside of the masthead content block. */
//   children?: React.ReactNode;
//   /** Additional classes added to the masthead content. */
//   className?: string;
// }
// 
// export const MastheadContent: React.FunctionComponent<MastheadContentProps> = ({
//   children,
//   className,
//   ...props
// }: MastheadContentProps) => (
//   <div className={css(styles.mastheadContent, className)} {...props}>
//     {children}
//   </div>
// );
// MastheadContent.displayName = 'MastheadContent';
