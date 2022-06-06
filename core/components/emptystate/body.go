package emptystate

type BodyProps struct {
	// Children - Content rendered inside the EmptyState. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the EmptyState. Optional.
	ClassName string
}

// contents of react-core/src/components/EmptyState/EmptyStateBody.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/EmptyState/empty-state';
// 
// export interface EmptyStateBodyProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the EmptyState */
//   children?: React.ReactNode;
//   /** Additional classes added to the EmptyState */
//   className?: string;
// }
// 
// export const EmptyStateBody: React.FunctionComponent<EmptyStateBodyProps> = ({
//   children,
//   className = '',
//   ...props
// }: EmptyStateBodyProps) => (
//   <div className={css(styles.emptyStateBody, className)} {...props}>
//     {children}
//   </div>
// );
// EmptyStateBody.displayName = 'EmptyStateBody';
