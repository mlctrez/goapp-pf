package radio

type Props struct {
	// ClassName - Additional classes added to the radio. Optional.
	ClassName string
	// Id - Id of the radio.
	Id string
	// IsLabelWrapped - Flag to show if the radio label is wrapped on small screen. Optional.
	IsLabelWrapped bool
	// IsLabelBeforeButton - Flag to show if the radio label is shown before the radio button. Optional.
	IsLabelBeforeButton bool
	// Checked - Flag to show if the radio is checked. Optional.
	Checked bool
	// IsChecked - Flag to show if the radio is checked. Optional.
	IsChecked bool
	// IsDisabled - Flag to show if the radio is disabled. Optional.
	IsDisabled bool
	// IsValid - Flag to show if the radio selection is valid or invalid. Optional.
	IsValid bool
	// Label - Label text of the radio. Optional.
	Label any //  // React.ReactNode 
	// Name - Name for group of radios.
	Name string
	// OnChange - A callback for when the radio selection changes. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// AriaLabel - Aria label for the radio. Optional.
	AriaLabel string
	// Description - Description text of the radio. Optional.
	Description any //  // React.ReactNode 
	// Body - Body of the radio. Optional.
	Body any //  // React.ReactNode 
}

// contents of react-core/src/components/Radio/Radio.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Radio/radio';
// import { css } from '@patternfly/react-styles';
// import { PickOptional } from '../../helpers/typeUtils';
// import { getOUIAProps, OUIAProps, getDefaultOUIAId } from '../../helpers';
// 
// export interface RadioProps
//   extends Omit<React.HTMLProps<HTMLInputElement>, 'disabled' | 'label' | 'onChange' | 'type'>,
//     OUIAProps {
//   /** Additional classes added to the radio. */
//   className?: string;
//   /** Id of the radio. */
//   id: string;
//   /** Flag to show if the radio label is wrapped on small screen. */
//   isLabelWrapped?: boolean;
//   /** Flag to show if the radio label is shown before the radio button. */
//   isLabelBeforeButton?: boolean;
//   /** Flag to show if the radio is checked. */
//   checked?: boolean;
//   /** Flag to show if the radio is checked. */
//   isChecked?: boolean;
//   /** Flag to show if the radio is disabled. */
//   isDisabled?: boolean;
//   /** Flag to show if the radio selection is valid or invalid. */
//   isValid?: boolean;
//   /** Label text of the radio. */
//   label?: React.ReactNode;
//   /** Name for group of radios */
//   name: string;
//   /** A callback for when the radio selection changes. */
//   onChange?: (checked: boolean, event: React.FormEvent<HTMLInputElement>) => void;
//   /** Aria label for the radio. */
//   'aria-label'?: string;
//   /** Description text of the radio. */
//   description?: React.ReactNode;
//   /** Body of the radio. */
//   body?: React.ReactNode;
// }
// 
// export class Radio extends React.Component<RadioProps, { ouiaStateId: string }> {
//   static displayName = 'Radio';
//   static defaultProps: PickOptional<RadioProps> = {
//     className: '',
//     isDisabled: false,
//     isValid: true,
//     onChange: () => {}
//   };
// 
//   constructor(props: RadioProps) {
//     super(props);
//     if (!props.label && !props['aria-label']) {
//       // eslint-disable-next-line no-console
//       console.error('Radio:', 'Radio requires an aria-label to be specified');
//     }
//     this.state = {
//       ouiaStateId: getDefaultOUIAId(Radio.displayName)
//     };
//   }
// 
//   handleChange = (event: React.FormEvent<HTMLInputElement>) => {
//     this.props.onChange(event.currentTarget.checked, event);
//   };
// 
//   render() {
//     const {
//       'aria-label': ariaLabel,
//       checked,
//       className,
//       defaultChecked,
//       isLabelWrapped,
//       isLabelBeforeButton,
//       isChecked,
//       isDisabled,
//       isValid,
//       label,
//       // eslint-disable-next-line @typescript-eslint/no-unused-vars
//       onChange,
//       description,
//       body,
//       ouiaId,
//       ouiaSafe = true,
//       ...props
//     } = this.props;
//     if (!props.id) {
//       // eslint-disable-next-line no-console
//       console.error('Radio:', 'id is required to make input accessible');
//     }
// 
//     const inputRendered = (
//       <input
//         {...props}
//         className={css(styles.radioInput)}
//         type="radio"
//         onChange={this.handleChange}
//         aria-invalid={!isValid}
//         disabled={isDisabled}
//         checked={checked || isChecked}
//         {...(checked === undefined && { defaultChecked })}
//         {...(!label && { 'aria-label': ariaLabel })}
//         {...getOUIAProps(Radio.displayName, ouiaId !== undefined ? ouiaId : this.state.ouiaStateId, ouiaSafe)}
//       />
//     );
// 
//     let labelRendered: React.ReactNode = null;
//     if (label && isLabelWrapped) {
//       labelRendered = <span className={css(styles.radioLabel, isDisabled && styles.modifiers.disabled)}>{label}</span>;
//     } else if (label) {
//       labelRendered = (
//         <label className={css(styles.radioLabel, isDisabled && styles.modifiers.disabled)} htmlFor={props.id}>
//           {label}
//         </label>
//       );
//     }
// 
//     const descRender = description ? <span className={css(styles.radioDescription)}>{description}</span> : null;
//     const bodyRender = body ? <span className={css(styles.radioBody)}>{body}</span> : null;
//     const childrenRendered = isLabelBeforeButton ? (
//       <>
//         {labelRendered}
//         {inputRendered}
//         {descRender}
//         {bodyRender}
//       </>
//     ) : (
//       <>
//         {inputRendered}
//         {labelRendered}
//         {descRender}
//         {bodyRender}
//       </>
//     );
// 
//     return isLabelWrapped ? (
//       <label className={css(styles.radio, className)} htmlFor={props.id}>
//         {childrenRendered}
//       </label>
//     ) : (
//       <div className={css(styles.radio, !label && styles.modifiers.standalone, className)}>{childrenRendered}</div>
//     );
//   }
// }
