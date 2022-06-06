package datalist

type DragButtonProps struct {
	// ClassName - Additional classes added to the drag button. Optional.
	ClassName string
	// Type - Sets button type. Optional.
	Type any //  /* "button" | "submit" | "reset" */
	// IsDisabled - Flag indicating if drag is disabled for the item. Optional.
	IsDisabled bool
}

// contents of react-core/src/components/DataList/DataListDragButton.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/DataList/data-list';
// import GripVerticalIcon from '@patternfly/react-icons/dist/esm/icons/grip-vertical-icon';
// import { DataListContext } from './DataList';
// 
// export interface DataListDragButtonProps extends React.HTMLProps<HTMLButtonElement> {
//   /** Additional classes added to the drag button */
//   className?: string;
//   /** Sets button type */
//   type?: 'button' | 'submit' | 'reset';
//   /** Flag indicating if drag is disabled for the item */
//   isDisabled?: boolean;
// }
// 
// export const DataListDragButton: React.FunctionComponent<DataListDragButtonProps> = ({
//   className = '',
//   isDisabled = false,
//   ...props
// }: DataListDragButtonProps) => (
//   <DataListContext.Consumer>
//     {({ dragKeyHandler }) => (
//       <button
//         className={css(styles.dataListItemDraggableButton, isDisabled && styles.modifiers.disabled, className)}
//         onKeyDown={dragKeyHandler}
//         type="button"
//         disabled={isDisabled}
//         {...props}
//       >
//         <span className={css(styles.dataListItemDraggableIcon)}>
//           <GripVerticalIcon />
//         </span>
//       </button>
//     )}
//   </DataListContext.Consumer>
// );
// DataListDragButton.displayName = 'DataListDragButton';
