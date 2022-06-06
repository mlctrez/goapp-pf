package duallistselector

type TreeItemData struct {
	// Children - Content rendered inside the dual list selector. Optional.
	Children []any // DualListSelectorTreeItemData 
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// DefaultExpanded - Flag indicating this option is expanded by default. Optional.
	DefaultExpanded bool
	// HasBadge - Flag indicating this option has a badge. Optional.
	HasBadge bool
	// OnOptionCheck - Callback fired when an option is checked. Optional.
	OnOptionCheck func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, isChosen bool, itemData any // DualListSelectorTreeItemData )
	// Id - ID of the option.
	Id string
	// Text - Text of the option.
	Text string
	// ParentId - Parent id of an option. Optional.
	ParentId string
	// IsChecked - Checked state of the option.
	IsChecked bool
	// CheckProps - Additional properties to pass to the option checkbox. Optional.
	CheckProps any // 
	// BadgeProps - Additional properties to pass to the option badge. Optional.
	BadgeProps any // 
	// IsDisabled - Flag indicating whether the component is disabled. Optional.
	IsDisabled bool
}

type TreeProps struct {
	// Data - Data of the tree view.
	Data any //  /* []any // DualListSelectorTreeItemData  | func() []any // DualListSelectorTreeItemData  */
	// Id - ID of the tree view. Optional.
	Id string
	// IsNested - Optional.
	IsNested bool
	// HasBadges - Flag indicating if all options should have badges. Optional.
	HasBadges bool
	// DefaultAllExpanded - Sets the default expanded behavior. Optional.
	DefaultAllExpanded bool
	// IsDisabled - Callback fired when an option is checked. Optional.
	IsDisabled bool
	// OnOptionCheck - Optional.
	OnOptionCheck func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, itemData any // DualListSelectorTreeItemData )
}

// contents of react-core/src/components/DualListSelector/DualListSelectorTree.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/DualListSelector/dual-list-selector';
// import { DualListSelectorTreeItem } from './DualListSelectorTreeItem';
// 
// export interface DualListSelectorTreeItemData {
//   /** Content rendered inside the dual list selector. */
//   children?: DualListSelectorTreeItemData[];
//   /** Additional classes applied to the dual list selector. */
//   className?: string;
//   /** Flag indicating this option is expanded by default. */
//   defaultExpanded?: boolean;
//   /** Flag indicating this option has a badge */
//   hasBadge?: boolean;
//   /** Callback fired when an option is checked */
//   onOptionCheck?: (
//     event: React.MouseEvent | React.ChangeEvent<HTMLInputElement> | React.KeyboardEvent,
//     isChecked: boolean,
//     isChosen: boolean,
//     itemData: DualListSelectorTreeItemData
//   ) => void;
//   /** ID of the option */
//   id: string;
//   /** Text of the option */
//   text: string;
//   /** Parent id of an option */
//   parentId?: string;
//   /** Checked state of the option */
//   isChecked: boolean;
//   /** Additional properties to pass to the option checkbox */
//   checkProps?: any;
//   /** Additional properties to pass to the option badge */
//   badgeProps?: any;
//   /** Flag indicating whether the component is disabled. */
//   isDisabled?: boolean;
// }
// 
// export interface DualListSelectorTreeProps {
//   /** Data of the tree view */
//   data: DualListSelectorTreeItemData[] | (() => DualListSelectorTreeItemData[]);
//   /** ID of the tree view */
//   id?: string;
//   /** @hide Flag indicating if the list is nested */
//   isNested?: boolean;
//   /** Flag indicating if all options should have badges */
//   hasBadges?: boolean;
//   /** Sets the default expanded behavior */
//   defaultAllExpanded?: boolean;
//   /** Callback fired when an option is checked */
//   isDisabled?: boolean;
//   onOptionCheck?: (
//     event: React.MouseEvent | React.ChangeEvent<HTMLInputElement> | React.KeyboardEvent,
//     isChecked: boolean,
//     itemData: DualListSelectorTreeItemData
//   ) => void;
// }
// 
// export const DualListSelectorTree: React.FunctionComponent<DualListSelectorTreeProps> = ({
//   data,
//   hasBadges = false,
//   isNested = false,
//   defaultAllExpanded = false,
//   onOptionCheck,
//   isDisabled = false,
//   ...props
// }: DualListSelectorTreeProps) => {
//   const dataToRender = typeof data === 'function' ? data() : data;
//   const tree = dataToRender.map(item => (
//     <DualListSelectorTreeItem
//       key={item.id}
//       text={item.text}
//       id={item.id}
//       defaultExpanded={item.defaultExpanded !== undefined ? item.defaultExpanded : defaultAllExpanded}
//       onOptionCheck={onOptionCheck}
//       isChecked={item.isChecked}
//       checkProps={item.checkProps}
//       hasBadge={item.hasBadge !== undefined ? item.hasBadge : hasBadges}
//       badgeProps={item.badgeProps}
//       itemData={item}
//       isDisabled={isDisabled}
//       useMemo={true}
//       {...(item.children && {
//         children: (
//           <DualListSelectorTree
//             isNested
//             data={item.children}
//             hasBadges={hasBadges}
//             defaultAllExpanded={defaultAllExpanded}
//             onOptionCheck={onOptionCheck}
//             isDisabled={isDisabled}
//           />
//         )
//       })}
//     />
//   ));
//   return isNested ? (
//     <ul className={css(styles.dualListSelectorList)} role="group" {...props}>
//       {tree}
//     </ul>
//   ) : (
//     <>{tree}</>
//   );
// };
// 
// DualListSelectorTree.displayName = 'DualListSelectorTree';
