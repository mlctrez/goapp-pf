package split

type Props struct {
	// HasGutter - Adds space between children. Optional.
	HasGutter bool
	// IsWrappable - Allows children to wrap. Optional.
	IsWrappable bool
	// Children - content rendered inside the Split layout. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the Split layout. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any //  // React.ReactNode 
}

// contents of react-core/src/layouts/Split/Split.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/layouts/Split/split';
// import { css } from '@patternfly/react-styles';
// 
// export interface SplitProps extends React.HTMLProps<HTMLDivElement> {
//   /** Adds space between children. */
//   hasGutter?: boolean;
//   /** Allows children to wrap */
//   isWrappable?: boolean;
//   /** content rendered inside the Split layout */
//   children?: React.ReactNode;
//   /** additional classes added to the Split layout */
//   className?: string;
//   /** Sets the base component to render. defaults to div */
//   component?: React.ReactNode;
// }
// 
// export const Split: React.FunctionComponent<SplitProps> = ({
//   hasGutter = false,
//   isWrappable = false,
//   className = '',
//   children = null,
//   component = 'div',
//   ...props
// }: SplitProps) => {
//   const Component = component as any;
//   return (
//     <Component
//       {...props}
//       className={css(
//         styles.split,
//         hasGutter && styles.modifiers.gutter,
//         isWrappable && styles.modifiers.wrap,
//         className
//       )}
//     >
//       {children}
//     </Component>
//   );
// };
// Split.displayName = 'Split';
