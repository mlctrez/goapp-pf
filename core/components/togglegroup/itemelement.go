package togglegroup

type ItemElementProps struct {
	// Children - Content rendered inside the toggle group item. Optional.
	Children any //  // React.ReactNode 
	// Variant - Adds toggle group item variant styles. Optional.
	Variant any //  /* any // ToggleGroupItemVariant  | "icon" | "text" */
}

// contents of react-core/src/components/ToggleGroup/ToggleGroupItemElement.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/ToggleGroup/toggle-group';
// 
// export enum ToggleGroupItemVariant {
//   icon = 'icon',
//   text = 'text'
// }
// 
// export interface ToggleGroupItemElementProps {
//   /** Content rendered inside the toggle group item */
//   children?: React.ReactNode;
//   /** Adds toggle group item variant styles */
//   variant?: ToggleGroupItemVariant | 'icon' | 'text';
// }
// 
// export const ToggleGroupItemElement: React.FunctionComponent<ToggleGroupItemElementProps> = ({ variant, children }) => (
//   <span className={css(variant === 'icon' && styles.toggleGroupIcon, variant === 'text' && styles.toggleGroupText)}>
//     {children}
//   </span>
// );
// ToggleGroupItemElement.displayName = 'ToggleGroupItemElement';
