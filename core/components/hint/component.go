package hint

type Props struct {
	// Children - Content rendered inside the hint. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes applied to the hint. Optional.
	ClassName string
	// Actions - Actions of the hint. Optional.
	Actions any //  // React.ReactNode 
}

// contents of react-core/src/components/Hint/Hint.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Hint/hint';
// import { css } from '@patternfly/react-styles';
// 
// export interface HintProps {
//   /** Content rendered inside the hint. */
//   children?: React.ReactNode;
//   /** Additional classes applied to the hint. */
//   className?: string;
//   /** Actions of the hint. */
//   actions?: React.ReactNode;
// }
// 
// export const Hint: React.FunctionComponent<HintProps> = ({ children, className, actions, ...props }: HintProps) => (
//   <div className={css(styles.hint, className)} {...props}>
//     <div className={css(styles.hintActions)}>{actions}</div>
//     {children}
//   </div>
// );
// Hint.displayName = 'Hint';
