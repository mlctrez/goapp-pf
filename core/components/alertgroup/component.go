package alertgroup

type Props struct {
	// ClassName - Additional classes added to the AlertGroup. Optional.
	ClassName string
	// Children - Alerts to be rendered in the AlertGroup. Optional.
	Children any //  // React.ReactNode 
	// IsToast - Toast notifications are positioned at the top right corner of the viewport. Optional.
	IsToast bool
	// IsLiveRegion - Turns the container into a live region so that changes to content within the AlertGroup,
	// such as appending an Alert, are reliably announced to assistive technology. Optional.
	IsLiveRegion bool
	// AppendTo - Determine where the alert is appended to. Optional.
	AppendTo any //  /* any // HTMLElement  | func() any // HTMLElement  */
	// OnOverflowClick - Function to call if user clicks on overflow message. Optional.
	OnOverflowClick func()
	// OverflowMessage - Custom text to show for the overflow message. Optional.
	OverflowMessage string
}

type State struct {
	// Container - 
	Container any //  // HTMLElement 
}

// contents of react-core/src/components/AlertGroup/AlertGroup.tsx
// import * as React from 'react';
// import * as ReactDOM from 'react-dom';
// import { canUseDOM } from '../../helpers';
// import { AlertGroupInline } from './AlertGroupInline';
// 
// export interface AlertGroupProps extends Omit<React.HTMLProps<HTMLUListElement>, 'className'> {
//   /** Additional classes added to the AlertGroup */
//   className?: string;
//   /** Alerts to be rendered in the AlertGroup */
//   children?: React.ReactNode;
//   /** Toast notifications are positioned at the top right corner of the viewport */
//   isToast?: boolean;
//   /** Turns the container into a live region so that changes to content within the AlertGroup, such as appending an Alert, are reliably announced to assistive technology. */
//   isLiveRegion?: boolean;
//   /** Determine where the alert is appended to */
//   appendTo?: HTMLElement | (() => HTMLElement);
//   /** Function to call if user clicks on overflow message */
//   onOverflowClick?: () => void;
//   /** Custom text to show for the overflow message */
//   overflowMessage?: string;
// }
// 
// interface AlertGroupState {
//   container: HTMLElement;
// }
// 
// export class AlertGroup extends React.Component<AlertGroupProps, AlertGroupState> {
//   static displayName = 'AlertGroup';
//   state = {
//     container: undefined
//   } as AlertGroupState;
// 
//   componentDidMount() {
//     const container = document.createElement('div');
//     const target: HTMLElement = this.getTargetElement();
//     this.setState({ container });
//     target.appendChild(container);
//   }
// 
//   componentWillUnmount() {
//     const target: HTMLElement = this.getTargetElement();
//     if (this.state.container) {
//       target.removeChild(this.state.container);
//     }
//   }
// 
//   getTargetElement() {
//     const appendTo = this.props.appendTo;
//     if (typeof appendTo === 'function') {
//       return appendTo();
//     }
//     return appendTo || document.body;
//   }
// 
//   render() {
//     const { className, children, isToast, isLiveRegion, onOverflowClick, overflowMessage, ...props } = this.props;
//     const alertGroup = (
//       <AlertGroupInline
//         onOverflowClick={onOverflowClick}
//         className={className}
//         isToast={isToast}
//         isLiveRegion={isLiveRegion}
//         overflowMessage={overflowMessage}
//         {...props}
//       >
//         {children}
//       </AlertGroupInline>
//     );
//     if (!this.props.isToast) {
//       return alertGroup;
//     }
// 
//     const container = this.state.container;
// 
//     if (!canUseDOM || !container) {
//       return null;
//     }
// 
//     return ReactDOM.createPortal(alertGroup, container);
//   }
// }
