package panel

type Props struct {
	// Children - Content rendered inside the panel. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
	// Variant - Adds panel variant styles. Optional.
	Variant any //  /* "raised" | "bordered" */
	// IsScrollable - Flag to add scrollable styling to the panel. Optional.
	IsScrollable bool
	// InnerRef - Optional.
	InnerRef any //  // React.Ref 
}

// contents of react-core/src/components/Panel/Panel.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Panel/panel';
// import { css } from '@patternfly/react-styles';
// 
// export interface PanelProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the panel */
//   children?: React.ReactNode;
//   /** Class to add to outer div */
//   className?: string;
//   /** Adds panel variant styles */
//   variant?: 'raised' | 'bordered';
//   /** Flag to add scrollable styling to the panel */
//   isScrollable?: boolean;
//   /** @hide Forwarded ref */
//   innerRef?: React.Ref<any>;
// }
// 
// const PanelBase: React.FunctionComponent<PanelProps> = ({
//   className,
//   children,
//   variant,
//   isScrollable,
//   innerRef,
//   ...props
// }: PanelProps) => (
//   <div
//     className={css(
//       styles.panel,
//       variant === 'raised' && styles.modifiers.raised,
//       variant === 'bordered' && styles.modifiers.bordered,
//       isScrollable && styles.modifiers.scrollable,
//       className
//     )}
//     ref={innerRef}
//     {...props}
//   >
//     {children}
//   </div>
// );
// 
// export const Panel = React.forwardRef((props: PanelProps, ref: React.Ref<any>) => (
//   <PanelBase innerRef={ref} {...props} />
// ));
// Panel.displayName = 'Panel';
