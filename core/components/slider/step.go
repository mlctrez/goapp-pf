package slider

type StepProps struct {
	// ClassName - Additional classes added to the slider steps. Optional.
	ClassName string
	// Value - Step value *. Optional.
	Value int
	// Label - Step label *. Optional.
	Label string
	// IsTickHidden - Flag indicating that the tick should be hidden. Optional.
	IsTickHidden bool
	// IsLabelHidden - Flag indicating that the label should be hidden. Optional.
	IsLabelHidden bool
	// IsActive - Flag indicating the step is active. Optional.
	IsActive bool
}

// contents of react-core/src/components/Slider/SliderStep.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Slider/slider';
// import { css } from '@patternfly/react-styles';
// 
// export interface SliderStepProps extends Omit<React.HTMLProps<HTMLDivElement>, 'label'> {
//   /** Additional classes added to the slider steps. */
//   className?: string;
//   /** Step value **/
//   value?: number;
//   /** Step label **/
//   label?: string;
//   /** Flag indicating that the tick should be hidden */
//   isTickHidden?: boolean;
//   /** Flag indicating that the label should be hidden */
//   isLabelHidden?: boolean;
//   /** Flag indicating the step is active */
//   isActive?: boolean;
// }
// 
// export const SliderStep: React.FunctionComponent<SliderStepProps> = ({
//   className,
//   label,
//   value,
//   isTickHidden = false,
//   isLabelHidden = false,
//   isActive = false,
//   ...props
// }: SliderStepProps) => {
//   const style = { '--pf-c-slider__step--Left': `${value}%!`(MISSING) } as React.CSSProperties;
//   return (
//     <div className={css(styles.sliderStep, isActive && styles.modifiers.active, className)} style={style} {...props}>
//       {!isTickHidden && <div className={css(styles.sliderStepTick)} />}
//       {!isLabelHidden && label && <div className={css(styles.sliderStepLabel)}>{label}</div>}
//     </div>
//   );
// };
// SliderStep.displayName = 'SliderStep';
