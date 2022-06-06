package aboutmodal

type BoxProps struct {
	// Children - content rendered inside the AboutModelBox.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the AboutModalBox. Optional.
	ClassName string
}

// contents of react-core/src/components/AboutModal/AboutModalBox.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/AboutModalBox/about-modal-box';
// 
// export interface AboutModalBoxProps extends React.HTMLProps<HTMLDivElement> {
//   /** content rendered inside the AboutModelBox.  */
//   children: React.ReactNode;
//   /** additional classes added to the AboutModalBox  */
//   className?: string;
// }
// 
// export const AboutModalBox: React.FunctionComponent<AboutModalBoxProps> = ({
//   children,
//   className = '',
//   ...props
// }: AboutModalBoxProps) => (
//   <div role="dialog" aria-modal="true" className={css(styles.aboutModalBox, className)} {...props}>
//     {children}
//   </div>
// );
// AboutModalBox.displayName = 'AboutModalBox';
