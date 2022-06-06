package progress

type AriaProps struct {
	// AriaLabelledby - Optional.
	AriaLabelledby string
	// AriaLabel - Optional.
	AriaLabel string
	// AriaValuemin - Optional.
	AriaValuemin int
	// AriaValuenow - Optional.
	AriaValuenow int
	// AriaValuemax - Optional.
	AriaValuemax int
	// AriaValuetext - Optional.
	AriaValuetext string
}

type BarProps struct {
	// Children - What should be rendered inside progress bar. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes for Progres bar. Optional.
	ClassName string
	// Value - Actual progress value.
	Value int
	// ProgressBarAriaProps - Minimal value of progress. Optional.
	ProgressBarAriaProps any //  // AriaProps 
}

// contents of react-core/src/components/Progress/ProgressBar.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Progress/progress';
// import { css } from '@patternfly/react-styles';
// 
// export interface AriaProps {
//   'aria-labelledby'?: string;
//   'aria-label'?: string;
//   'aria-valuemin'?: number;
//   'aria-valuenow'?: number;
//   'aria-valuemax'?: number;
//   'aria-valuetext'?: string;
// }
// 
// export interface ProgressBarProps extends React.HTMLProps<HTMLDivElement> {
//   /** What should be rendered inside progress bar. */
//   children?: React.ReactNode;
//   /** Additional classes for Progres bar. */
//   className?: string;
//   /** Actual progress value. */
//   value: number;
//   /** Minimal value of progress. */
//   progressBarAriaProps?: AriaProps;
// }
// 
// export const ProgressBar: React.FunctionComponent<ProgressBarProps> = ({
//   progressBarAriaProps,
//   className = '',
//   children = null,
//   value,
//   ...props
// }: ProgressBarProps) => (
//   <div {...props} className={css(styles.progressBar, className)} {...progressBarAriaProps}>
//     <div className={css(styles.progressIndicator)} style={{ width: `${value}%!`(MISSING) }}>
//       <span className={css(styles.progressMeasure)}>{children}</span>
//     </div>
//   </div>
// );
// ProgressBar.displayName = 'ProgressBar';
