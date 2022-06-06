package emptystate

type Props struct {
	// ClassName - Additional classes added to the EmptyState. Optional.
	ClassName string
	// Children - Content rendered inside the EmptyState.
	Children any //  // React.ReactNode 
	// Variant - Modifies EmptyState max-width. Optional.
	Variant any //  /* "xs" | "small" | "large" | "xl" | "full" */
	// IsFullHeight - Cause component to consume the available height of its container. Optional.
	IsFullHeight bool
}

// contents of react-core/src/components/EmptyState/EmptyState.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/EmptyState/empty-state';
// 
// export enum EmptyStateVariant {
//   'xs' = 'xs',
//   small = 'small',
//   large = 'large',
//   'xl' = 'xl',
//   full = 'full'
// }
// 
// export interface EmptyStateProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the EmptyState */
//   className?: string;
//   /** Content rendered inside the EmptyState */
//   children: React.ReactNode;
//   /** Modifies EmptyState max-width */
//   variant?: 'xs' | 'small' | 'large' | 'xl' | 'full';
//   /** Cause component to consume the available height of its container */
//   isFullHeight?: boolean;
// }
// 
// export const EmptyState: React.FunctionComponent<EmptyStateProps> = ({
//   children,
//   className = '',
//   variant = EmptyStateVariant.full,
//   isFullHeight,
//   ...props
// }: EmptyStateProps) => (
//   <div
//     className={css(
//       styles.emptyState,
//       variant === 'xs' && styles.modifiers.xs,
//       variant === 'small' && styles.modifiers.sm,
//       variant === 'large' && styles.modifiers.lg,
//       variant === 'xl' && styles.modifiers.xl,
//       isFullHeight && styles.modifiers.fullHeight,
//       className
//     )}
//     {...props}
//   >
//     <div className={css(styles.emptyStateContent)}>{children}</div>
//   </div>
// );
// EmptyState.displayName = 'EmptyState';
