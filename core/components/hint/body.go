package hint

type BodyProps struct {
	// Children - Content rendered inside the hint body. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes applied to the hint body. Optional.
	ClassName string
}

// contents of react-core/src/components/Hint/HintBody.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Hint/hint';
// import { css } from '@patternfly/react-styles';
// 
// export interface HintBodyProps {
//   /** Content rendered inside the hint body. */
//   children?: React.ReactNode;
//   /** Additional classes applied to the hint body. */
//   className?: string;
// }
// 
// export const HintBody: React.FunctionComponent<HintBodyProps> = ({ children, className, ...props }: HintBodyProps) => (
//   <div className={css(styles.hintBody, className)} {...props}>
//     {children}
//   </div>
// );
// HintBody.displayName = 'HintBody';
