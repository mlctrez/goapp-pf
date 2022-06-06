package duallistselector

type Props struct {
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// Id - Id of the dual list selector. Optional.
	Id string
	// IsTree - Flag indicating if the dual list selector uses trees instead of simple lists. Optional.
	IsTree bool
	// IsDisabled - Flag indicating if the dual list selector is in a disabled state. Optional.
	IsDisabled bool
	// Children - Content to be rendered in the dual list selector. Panes & controls will not be built dynamically
	// when children are provided. Optional.
	Children any //  // React.ReactNode 
	// AvailableOptionsTitle - Title applied to the dynamically built available options pane. Optional.
	AvailableOptionsTitle string
	// AvailableOptions - Options to display in the dynamically built available options pane. When using trees,
	// the options should be in the DualListSelectorTreeItemData[] format. Optional.
	AvailableOptions any //  /* []any // React.ReactNode  | []any // DualListSelectorTreeItemData  */
	// AvailableOptionsStatus - Status message to display above the dynamically built available options pane.
	// Optional.
	AvailableOptionsStatus string
	// AvailableOptionsActions - Actions to be displayed above the dynamically built available options pane.
	// Optional.
	AvailableOptionsActions []any // React.ReactNode 
	// ChosenOptionsTitle - Title applied to the dynamically built chosen options pane. Optional.
	ChosenOptionsTitle string
	// ChosenOptions - Options to display in the dynamically built chosen options pane. When using trees, the
	// options should be in the DualListSelectorTreeItemData[] format. Optional.
	ChosenOptions any //  /* []any // React.ReactNode  | []any // DualListSelectorTreeItemData  */
	// ChosenOptionsStatus - Status message to display above the dynamically built chosen options pane. Optional.
	ChosenOptionsStatus string
	// ChosenOptionsActions - Actions to be displayed above the dynamically built chosen options pane. Optional.
	ChosenOptionsActions []any // React.ReactNode 
	// ControlsAriaLabel - Accessible label for the dynamically built controls between the two panes. Optional.
	ControlsAriaLabel string
	// AddSelected - Optional callback for the dynamically built add selected button. Optional.
	AddSelected func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddSelectedAriaLabel - Accessible label for the dynamically built add selected button. Optional.
	AddSelectedAriaLabel string
	// AddSelectedTooltip - Tooltip content for the dynamically built add selected button. Optional.
	AddSelectedTooltip any //  // React.ReactNode 
	// AddSelectedTooltipProps - Additonal tooltip properties for the dynamically built add selected tooltip.
	// Optional.
	AddSelectedTooltipProps any // 
	// OnListChange - Callback fired every time dynamically built options are chosen or removed. Optional.
	OnListChange func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddAll - Optional callback for the dynamically built add all button. Optional.
	AddAll func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddAllAriaLabel - Accessible label for the dynamically built add all button. Optional.
	AddAllAriaLabel string
	// AddAllTooltip - Tooltip content for the dynamically built add all button. Optional.
	AddAllTooltip any //  // React.ReactNode 
	// AddAllTooltipProps - Additonal tooltip properties for the dynamically built add all tooltip. Optional.
	AddAllTooltipProps any // 
	// RemoveSelected - Optional callback for the dynamically built remove selected button. Optional.
	RemoveSelected func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// RemoveSelectedAriaLabel - Accessible label for the dynamically built remove selected button. Optional.
	RemoveSelectedAriaLabel string
	// RemoveSelectedTooltip - Tooltip content for the dynamically built remove selected button. Optional.
	RemoveSelectedTooltip any //  // React.ReactNode 
	// RemoveSelectedTooltipProps - Additonal tooltip properties for the dynamically built remove selected tooltip.
	// Optional.
	RemoveSelectedTooltipProps any // 
	// RemoveAll - Optional callback for the dynamically built remove all button. Optional.
	RemoveAll func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// RemoveAllAriaLabel - Accessible label for the dynamically built remove all button. Optional.
	RemoveAllAriaLabel string
	// RemoveAllTooltip - Tooltip content for the dynamically built remove all button. Optional.
	RemoveAllTooltip any //  // React.ReactNode 
	// RemoveAllTooltipProps - Additonal tooltip properties for the dynamically built remove all tooltip. Optional.
	RemoveAllTooltipProps any // 
	// OnOptionSelect - Optional callback fired when a dynamically built option is selected. Optional.
	OnOptionSelect func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, index int, isChosen bool, id string, itemData any, parentData any)
	// OnOptionCheck - Optional callback fired when a dynamically built option is checked. Optional.
	OnOptionCheck func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, checked bool, checkedId string, newCheckedItems []string)
	// IsSearchable - Flag indicating a search bar should be included above both the dynamically built available
	// and chosen panes. Optional.
	IsSearchable bool
	// AvailableOptionsSearchAriaLabel - Accessible label for the search input on the dynamically built available
	// options pane. Optional.
	AvailableOptionsSearchAriaLabel string
	// OnAvailableOptionsSearchInputChanged - A callback for when the search input value for the dynamically
	// built available options changes. Optional.
	OnAvailableOptionsSearchInputChanged func(value string, event any // React.FormEvent )
	// ChosenOptionsSearchAriaLabel - Accessible label for the search input on the dynamically built chosen options
	// pane. Optional.
	ChosenOptionsSearchAriaLabel string
	// OnChosenOptionsSearchInputChanged - A callback for when the search input value for the dynamically built
	// chosen options changes. Optional.
	OnChosenOptionsSearchInputChanged func(value string, event any // React.FormEvent )
	// FilterOption - Optional filter function for custom filtering based on search string. Used with a dynamically
	// built search input. Optional.
	FilterOption func(option any // React.ReactNode , input string) bool
}

type State struct {
	// AvailableOptions - 
	AvailableOptions []any // React.ReactNode 
	// AvailableOptionsSelected - 
	AvailableOptionsSelected []int
	// AvailableFilteredOptions - 
	AvailableFilteredOptions []any // React.ReactNode 
	// ChosenOptions - 
	ChosenOptions []any // React.ReactNode 
	// ChosenOptionsSelected - 
	ChosenOptionsSelected []int
	// ChosenFilteredOptions - 
	ChosenFilteredOptions []any // React.ReactNode 
	// AvailableTreeFilteredOptions - 
	AvailableTreeFilteredOptions []string
	// AvailableTreeOptionsChecked - 
	AvailableTreeOptionsChecked []string
	// ChosenTreeOptionsChecked - 
	ChosenTreeOptionsChecked []string
	// ChosenTreeFilteredOptions - 
	ChosenTreeFilteredOptions []string
}

// contents of react-core/src/components/DualListSelector/DualListSelector.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/DualListSelector/dual-list-selector';
// import { css } from '@patternfly/react-styles';
// import AngleDoubleLeftIcon from '@patternfly/react-icons/dist/esm/icons/angle-double-left-icon';
// import AngleLeftIcon from '@patternfly/react-icons/dist/esm/icons/angle-left-icon';
// import AngleDoubleRightIcon from '@patternfly/react-icons/dist/esm/icons/angle-double-right-icon';
// import AngleRightIcon from '@patternfly/react-icons/dist/esm/icons/angle-right-icon';
// import { DualListSelectorPane } from './DualListSelectorPane';
// import { getUniqueId, PickOptional } from '../../helpers';
// import { DualListSelectorTreeItemData } from './DualListSelectorTree';
// import {
//   flattenTree,
//   flattenTreeWithFolders,
//   filterFolders,
//   filterTreeItems,
//   filterTreeItemsWithoutFolders,
//   filterRestTreeItems
// } from './treeUtils';
// import { DualListSelectorControlsWrapper } from './DualListSelectorControlsWrapper';
// import { DualListSelectorControl } from './DualListSelectorControl';
// import { DualListSelectorContext } from './DualListSelectorContext';
// 
// export interface DualListSelectorProps {
//   /** Additional classes applied to the dual list selector. */
//   className?: string;
//   /** Id of the dual list selector. */
//   id?: string;
//   /** Flag indicating if the dual list selector uses trees instead of simple lists */
//   isTree?: boolean;
//   /** Flag indicating if the dual list selector is in a disabled state */
//   isDisabled?: boolean;
//   /** Content to be rendered in the dual list selector. Panes & controls will not be built dynamically when children are provided. */
//   children?: React.ReactNode;
//   /** Title applied to the dynamically built available options pane. */
//   availableOptionsTitle?: string;
//   /** Options to display in the dynamically built available options pane. When using trees, the options should be in the DualListSelectorTreeItemData[] format. */
//   availableOptions?: React.ReactNode[] | DualListSelectorTreeItemData[];
//   /** Status message to display above the dynamically built available options pane. */
//   availableOptionsStatus?: string;
//   /** Actions to be displayed above the dynamically built available options pane. */
//   availableOptionsActions?: React.ReactNode[];
//   /** Title applied to the dynamically built chosen options pane. */
//   chosenOptionsTitle?: string;
//   /** Options to display in the dynamically built chosen options pane. When using trees, the options should be in the DualListSelectorTreeItemData[] format. */
//   chosenOptions?: React.ReactNode[] | DualListSelectorTreeItemData[];
//   /** Status message to display above the dynamically built chosen options pane.*/
//   chosenOptionsStatus?: string;
//   /** Actions to be displayed above the dynamically built chosen options pane. */
//   chosenOptionsActions?: React.ReactNode[];
//   /** Accessible label for the dynamically built controls between the two panes. */
//   controlsAriaLabel?: string;
//   /** Optional callback for the dynamically built add selected button */
//   addSelected?: (newAvailableOptions: React.ReactNode[], newChosenOptions: React.ReactNode[]) => void;
//   /** Accessible label for the dynamically built add selected button */
//   addSelectedAriaLabel?: string;
//   /** Tooltip content for the dynamically built add selected button */
//   addSelectedTooltip?: React.ReactNode;
//   /** Additonal tooltip properties for the dynamically built add selected tooltip */
//   addSelectedTooltipProps?: any;
//   /** Callback fired every time dynamically built options are chosen or removed */
//   onListChange?: (newAvailableOptions: React.ReactNode[], newChosenOptions: React.ReactNode[]) => void;
//   /** Optional callback for the dynamically built add all button */
//   addAll?: (newAvailableOptions: React.ReactNode[], newChosenOptions: React.ReactNode[]) => void;
//   /** Accessible label for the dynamically built add all button */
//   addAllAriaLabel?: string;
//   /** Tooltip content for the dynamically built add all button */
//   addAllTooltip?: React.ReactNode;
//   /** Additonal tooltip properties for the dynamically built add all tooltip */
//   addAllTooltipProps?: any;
//   /** Optional callback for the dynamically built remove selected button */
//   removeSelected?: (newAvailableOptions: React.ReactNode[], newChosenOptions: React.ReactNode[]) => void;
//   /** Accessible label for the dynamically built remove selected button */
//   removeSelectedAriaLabel?: string;
//   /** Tooltip content for the dynamically built remove selected button */
//   removeSelectedTooltip?: React.ReactNode;
//   /** Additonal tooltip properties for the dynamically built remove selected tooltip  */
//   removeSelectedTooltipProps?: any;
//   /** Optional callback for the dynamically built remove all button */
//   removeAll?: (newAvailableOptions: React.ReactNode[], newChosenOptions: React.ReactNode[]) => void;
//   /** Accessible label for the dynamically built remove all button */
//   removeAllAriaLabel?: string;
//   /** Tooltip content for the dynamically built remove all button */
//   removeAllTooltip?: React.ReactNode;
//   /** Additonal tooltip properties for the dynamically built remove all tooltip */
//   removeAllTooltipProps?: any;
//   /** Optional callback fired when a dynamically built option is selected */
//   onOptionSelect?: (
//     e: React.MouseEvent | React.ChangeEvent | React.KeyboardEvent,
//     index: number,
//     isChosen: boolean,
//     id: string,
//     itemData: any,
//     parentData: any
//   ) => void;
//   /** Optional callback fired when a dynamically built option is checked */
//   onOptionCheck?: (
//     e: React.MouseEvent | React.ChangeEvent<HTMLInputElement> | React.KeyboardEvent,
//     checked: boolean,
//     checkedId: string,
//     newCheckedItems: string[]
//   ) => void;
//   /** Flag indicating a search bar should be included above both the dynamically built available and chosen panes. */
//   isSearchable?: boolean;
//   /** Accessible label for the search input on the dynamically built available options pane. */
//   availableOptionsSearchAriaLabel?: string;
//   /** A callback for when the search input value for the dynamically built available options changes. */
//   onAvailableOptionsSearchInputChanged?: (value: string, event: React.FormEvent<HTMLInputElement>) => void;
//   /** Accessible label for the search input on the dynamically built chosen options pane. */
//   chosenOptionsSearchAriaLabel?: string;
//   /** A callback for when the search input value for the dynamically built chosen options changes. */
//   onChosenOptionsSearchInputChanged?: (value: string, event: React.FormEvent<HTMLInputElement>) => void;
//   /** Optional filter function for custom filtering based on search string. Used with a dynamically built search input. */
//   filterOption?: (option: React.ReactNode, input: string) => boolean;
// }
// 
// interface DualListSelectorState {
//   availableOptions: React.ReactNode[];
//   availableOptionsSelected: number[];
//   availableFilteredOptions: React.ReactNode[];
//   chosenOptions: React.ReactNode[];
//   chosenOptionsSelected: number[];
//   chosenFilteredOptions: React.ReactNode[];
//   availableTreeFilteredOptions: string[];
//   availableTreeOptionsChecked: string[];
//   chosenTreeOptionsChecked: string[];
//   chosenTreeFilteredOptions: string[];
// }
// 
// export class DualListSelector extends React.Component<DualListSelectorProps, DualListSelectorState> {
//   static displayName = 'DualListSelector';
//   private addAllButtonRef = React.createRef<HTMLButtonElement>();
//   private addSelectedButtonRef = React.createRef<HTMLButtonElement>();
//   private removeSelectedButtonRef = React.createRef<HTMLButtonElement>();
//   private removeAllButtonRef = React.createRef<HTMLButtonElement>();
//   static defaultProps: PickOptional<DualListSelectorProps> = {
//     children: '',
//     availableOptions: [] as React.ReactNode[],
//     availableOptionsTitle: 'Available options',
//     availableOptionsSearchAriaLabel: 'Available search input',
//     chosenOptions: [] as React.ReactNode[],
//     chosenOptionsTitle: 'Chosen options',
//     chosenOptionsSearchAriaLabel: 'Chosen search input',
//     id: getUniqueId('dual-list-selector'),
//     controlsAriaLabel: 'Selector controls',
//     addAllAriaLabel: 'Add all',
//     addSelectedAriaLabel: 'Add selected',
//     removeSelectedAriaLabel: 'Remove selected',
//     removeAllAriaLabel: 'Remove all',
//     isTree: false,
//     isDisabled: false
//   };
// 
//   // If the DualListSelector uses trees, concat the two initial arrays and merge duplicate folder IDs
//   private createMergedCopy() {
//     const copyOfAvailable = JSON.parse(JSON.stringify(this.props.availableOptions));
//     const copyOfChosen = JSON.parse(JSON.stringify(this.props.chosenOptions));
// 
//     return this.props.isTree
//       ? Object.values(
//           (copyOfAvailable as DualListSelectorTreeItemData[])
//             .concat(copyOfChosen as DualListSelectorTreeItemData[])
//             .reduce((mapObj: any, item: DualListSelectorTreeItemData) => {
//               const key = item.id;
//               if (mapObj[key]) {
//                 // If map already has an item ID, add the dupe ID's children to the existing map
//                 mapObj[key].children.push(...item.children);
//               } else {
//                 // Else clone the item data
//                 mapObj[key] = { ...item };
//               }
//               return mapObj;
//             }, {})
//         )
//       : null;
//   }
// 
//   constructor(props: DualListSelectorProps) {
//     super(props);
//     this.state = {
//       availableOptions: [...this.props.availableOptions],
//       availableOptionsSelected: [],
//       availableFilteredOptions: null,
//       availableTreeFilteredOptions: null,
//       chosenOptions: [...this.props.chosenOptions],
//       chosenOptionsSelected: [],
//       chosenFilteredOptions: null,
//       chosenTreeFilteredOptions: null,
//       availableTreeOptionsChecked: [],
//       chosenTreeOptionsChecked: []
//     };
//   }
// 
//   componentDidUpdate() {
//     if (
//       JSON.stringify(this.props.availableOptions) !== JSON.stringify(this.state.availableOptions) ||
//       JSON.stringify(this.props.chosenOptions) !== JSON.stringify(this.state.chosenOptions)
//     ) {
//       this.setState({
//         availableOptions: [...this.props.availableOptions],
//         chosenOptions: [...this.props.chosenOptions]
//       });
//     }
//   }
// 
//   onFilterUpdate = (newFilteredOptions: React.ReactNode[], paneType: string, isSearchReset: boolean) => {
//     const { isTree } = this.props;
//     if (paneType === 'available') {
//       if (isSearchReset) {
//         this.setState({
//           availableFilteredOptions: null,
//           availableTreeFilteredOptions: null
//         });
//         return;
//       }
//       if (isTree) {
//         this.setState({
//           availableTreeFilteredOptions: flattenTreeWithFolders(newFilteredOptions as DualListSelectorTreeItemData[])
//         });
//       } else {
//         this.setState({
//           availableFilteredOptions: newFilteredOptions as React.ReactNode[]
//         });
//       }
//     } else if (paneType === 'chosen') {
//       if (isSearchReset) {
//         this.setState({
//           chosenFilteredOptions: null,
//           chosenTreeFilteredOptions: null
//         });
//         return;
//       }
//       if (isTree) {
//         this.setState({
//           chosenTreeFilteredOptions: flattenTreeWithFolders(newFilteredOptions as DualListSelectorTreeItemData[])
//         });
//       } else {
//         this.setState({
//           chosenFilteredOptions: newFilteredOptions as React.ReactNode[]
//         });
//       }
//     }
//   };
// 
//   addAllVisible = () => {
//     this.setState(prevState => {
//       const itemsToRemove = [] as React.ReactNode[];
//       const newAvailable = [] as React.ReactNode[];
//       const movedOptions = prevState.availableFilteredOptions || prevState.availableOptions;
//       prevState.availableOptions.forEach(value => {
//         if (movedOptions.indexOf(value) !== -1) {
//           itemsToRemove.push(value);
//         } else {
//           newAvailable.push(value);
//         }
//       });
// 
//       const newChosen = [...prevState.chosenOptions, ...itemsToRemove];
//       this.props.addAll && this.props.addAll(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptions: newChosen,
//         availableOptions: newAvailable,
//         chosenOptionsSelected: [],
//         availableOptionsSelected: []
//       };
//     });
//   };
// 
//   addAllTreeVisible = () => {
//     this.setState(prevState => {
//       const movedOptions =
//         prevState.availableTreeFilteredOptions ||
//         flattenTreeWithFolders(prevState.availableOptions as DualListSelectorTreeItemData[]);
//       const newAvailable = prevState.availableOptions
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterRestTreeItems(item as DualListSelectorTreeItemData, movedOptions));
// 
//       const currChosen = flattenTree(prevState.chosenOptions as DualListSelectorTreeItemData[]);
//       const nextChosenOptions = currChosen.concat(movedOptions);
//       const newChosen = this.createMergedCopy()
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, nextChosenOptions));
// 
//       this.props.addAll && this.props.addAll(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptions: newChosen,
//         chosenFilteredOptions: newChosen,
//         availableOptions: newAvailable,
//         availableFilteredOptions: newAvailable,
//         availableTreeOptionsChecked: [],
//         chosenTreeOptionsChecked: []
//       };
//     });
//   };
// 
//   addSelected = () => {
//     this.setState(prevState => {
//       const itemsToRemove = [] as React.ReactNode[];
//       const newAvailable = [] as React.ReactNode[];
//       prevState.availableOptions.forEach((value, index) => {
//         if (prevState.availableOptionsSelected.indexOf(index) !== -1) {
//           itemsToRemove.push(value);
//         } else {
//           newAvailable.push(value);
//         }
//       });
// 
//       const newChosen = [...prevState.chosenOptions, ...itemsToRemove];
//       this.props.addSelected && this.props.addSelected(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptionsSelected: [],
//         availableOptionsSelected: [],
//         chosenOptions: newChosen,
//         availableOptions: newAvailable
//       };
//     });
//   };
// 
//   addTreeSelected = () => {
//     this.setState(prevState => {
//       // Remove selected available nodes from current available nodes
//       const newAvailable = prevState.availableOptions
//         .map(opt => Object.assign({}, opt))
//         .filter(item =>
//           filterRestTreeItems(item as DualListSelectorTreeItemData, prevState.availableTreeOptionsChecked)
//         );
// 
//       // Get next chosen options from current + new nodes and remap from base
//       const currChosen = flattenTree(prevState.chosenOptions as DualListSelectorTreeItemData[]);
//       const nextChosenOptions = currChosen.concat(prevState.availableTreeOptionsChecked);
//       const newChosen = this.createMergedCopy()
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, nextChosenOptions));
// 
//       this.props.addSelected && this.props.addSelected(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         availableTreeOptionsChecked: [],
//         chosenTreeOptionsChecked: [],
//         availableOptions: newAvailable,
//         chosenOptions: newChosen
//       };
//     });
//   };
// 
//   removeAllVisible = () => {
//     this.setState(prevState => {
//       const itemsToRemove = [] as React.ReactNode[];
//       const newChosen = [] as React.ReactNode[];
//       const movedOptions = prevState.chosenFilteredOptions || prevState.chosenOptions;
//       prevState.chosenOptions.forEach(value => {
//         if (movedOptions.indexOf(value) !== -1) {
//           itemsToRemove.push(value);
//         } else {
//           newChosen.push(value);
//         }
//       });
// 
//       const newAvailable = [...prevState.availableOptions, ...itemsToRemove];
//       this.props.removeAll && this.props.removeAll(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptions: newChosen,
//         availableOptions: newAvailable,
//         chosenOptionsSelected: [],
//         availableOptionsSelected: []
//       };
//     });
//   };
// 
//   removeAllTreeVisible = () => {
//     this.setState(prevState => {
//       const movedOptions =
//         prevState.chosenTreeFilteredOptions ||
//         flattenTreeWithFolders(prevState.chosenOptions as DualListSelectorTreeItemData[]);
// 
//       const newChosen = prevState.chosenOptions
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterRestTreeItems(item as DualListSelectorTreeItemData, movedOptions));
//       const currAvailable = flattenTree(prevState.availableOptions as DualListSelectorTreeItemData[]);
//       const nextAvailableOptions = currAvailable.concat(movedOptions);
//       const newAvailable = this.createMergedCopy()
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, nextAvailableOptions));
// 
//       this.props.removeAll && this.props.removeAll(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptions: newChosen,
//         availableOptions: newAvailable,
//         availableTreeOptionsChecked: [],
//         chosenTreeOptionsChecked: []
//       };
//     });
//   };
// 
//   removeSelected = () => {
//     this.setState(prevState => {
//       const itemsToRemove = [] as React.ReactNode[];
//       const newChosen = [] as React.ReactNode[];
//       prevState.chosenOptions.forEach((value, index) => {
//         if (prevState.chosenOptionsSelected.indexOf(index) !== -1) {
//           itemsToRemove.push(value);
//         } else {
//           newChosen.push(value);
//         }
//       });
// 
//       const newAvailable = [...prevState.availableOptions, ...itemsToRemove];
//       this.props.removeSelected && this.props.removeSelected(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         chosenOptionsSelected: [],
//         availableOptionsSelected: [],
//         chosenOptions: newChosen,
//         availableOptions: newAvailable
//       };
//     });
//   };
// 
//   removeTreeSelected = () => {
//     this.setState(prevState => {
//       // Remove selected chosen nodes from current chosen nodes
//       const newChosen = prevState.chosenOptions
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterRestTreeItems(item as DualListSelectorTreeItemData, prevState.chosenTreeOptionsChecked));
// 
//       // Get next chosen options from current and remap from base
//       const currAvailable = flattenTree(prevState.availableOptions as DualListSelectorTreeItemData[]);
//       const nextAvailableOptions = currAvailable.concat(prevState.chosenTreeOptionsChecked);
//       const newAvailable = this.createMergedCopy()
//         .map(opt => Object.assign({}, opt))
//         .filter(item => filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, nextAvailableOptions));
// 
//       this.props.removeSelected && this.props.removeSelected(newAvailable, newChosen);
//       this.props.onListChange && this.props.onListChange(newAvailable, newChosen);
// 
//       return {
//         availableTreeOptionsChecked: [],
//         chosenTreeOptionsChecked: [],
//         availableOptions: newAvailable,
//         chosenOptions: newChosen
//       };
//     });
//   };
// 
//   onOptionSelect = (
//     e: React.MouseEvent | React.ChangeEvent | React.KeyboardEvent,
//     index: number,
//     isChosen: boolean,
//     /* eslint-disable @typescript-eslint/no-unused-vars */
//     id?: string,
//     itemData?: any,
//     parentData?: any
//     /* eslint-enable @typescript-eslint/no-unused-vars */
//   ) => {
//     this.setState(prevState => {
//       const originalArray = isChosen ? prevState.chosenOptionsSelected : prevState.availableOptionsSelected;
// 
//       let updatedArray = null;
//       if (originalArray.indexOf(index) !== -1) {
//         updatedArray = originalArray.filter(value => value !== index);
//       } else {
//         updatedArray = [...originalArray, index];
//       }
// 
//       return {
//         chosenOptionsSelected: isChosen ? updatedArray : prevState.chosenOptionsSelected,
//         availableOptionsSelected: isChosen ? prevState.availableOptionsSelected : updatedArray
//       };
//     });
// 
//     this.props.onOptionSelect && this.props.onOptionSelect(e, index, isChosen, id, itemData, parentData);
//   };
// 
//   isChecked = (treeItem: DualListSelectorTreeItemData, isChosen: boolean) =>
//     isChosen
//       ? this.state.chosenTreeOptionsChecked.includes(treeItem.id)
//       : this.state.availableTreeOptionsChecked.includes(treeItem.id);
//   areAllDescendantsChecked = (treeItem: DualListSelectorTreeItemData, isChosen: boolean): boolean =>
//     treeItem.children
//       ? treeItem.children.every(child => this.areAllDescendantsChecked(child, isChosen))
//       : this.isChecked(treeItem, isChosen);
//   areSomeDescendantsChecked = (treeItem: DualListSelectorTreeItemData, isChosen: boolean): boolean =>
//     treeItem.children
//       ? treeItem.children.some(child => this.areSomeDescendantsChecked(child, isChosen))
//       : this.isChecked(treeItem, isChosen);
// 
//   mapChecked = (item: DualListSelectorTreeItemData, isChosen: boolean): DualListSelectorTreeItemData => {
//     const hasCheck = this.areAllDescendantsChecked(item, isChosen);
//     item.isChecked = false;
// 
//     if (hasCheck) {
//       item.isChecked = true;
//     } else {
//       const hasPartialCheck = this.areSomeDescendantsChecked(item, isChosen);
//       if (hasPartialCheck) {
//         item.isChecked = null;
//       }
//     }
// 
//     if (item.children) {
//       return {
//         ...item,
//         children: item.children.map(child => this.mapChecked(child, isChosen))
//       };
//     }
//     return item;
//   };
// 
//   onTreeOptionCheck = (
//     evt: React.MouseEvent | React.ChangeEvent<HTMLInputElement> | React.KeyboardEvent,
//     isChecked: boolean,
//     itemData: DualListSelectorTreeItemData,
//     isChosen: boolean
//   ) => {
//     const { availableOptions, availableTreeFilteredOptions, chosenOptions, chosenTreeFilteredOptions } = this.state;
//     let panelOptions;
//     if (isChosen) {
//       if (chosenTreeFilteredOptions) {
//         panelOptions = chosenOptions
//           .map(opt => Object.assign({}, opt))
//           .filter(item =>
//             filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, chosenTreeFilteredOptions)
//           );
//       } else {
//         panelOptions = chosenOptions;
//       }
//     } else {
//       if (availableTreeFilteredOptions) {
//         panelOptions = availableOptions
//           .map(opt => Object.assign({}, opt))
//           .filter(item =>
//             filterTreeItemsWithoutFolders(item as DualListSelectorTreeItemData, availableTreeFilteredOptions)
//           );
//       } else {
//         panelOptions = availableOptions;
//       }
//     }
//     const checkedOptionTree = panelOptions
//       .map(opt => Object.assign({}, opt))
//       .filter(item => filterTreeItems(item as DualListSelectorTreeItemData, [itemData.id]));
//     const flatTree = flattenTreeWithFolders(checkedOptionTree as DualListSelectorTreeItemData[]);
// 
//     const prevChecked = isChosen ? this.state.chosenTreeOptionsChecked : this.state.availableTreeOptionsChecked;
//     let updatedChecked = [] as string[];
//     if (isChecked) {
//       updatedChecked = prevChecked.concat(flatTree.filter(id => !prevChecked.includes(id)));
//     } else {
//       updatedChecked = prevChecked.filter(id => !flatTree.includes(id));
//     }
// 
//     this.setState(
//       prevState => ({
//         availableTreeOptionsChecked: isChosen ? prevState.availableTreeOptionsChecked : updatedChecked,
//         chosenTreeOptionsChecked: isChosen ? updatedChecked : prevState.chosenTreeOptionsChecked
//       }),
//       () => {
//         this.props.onOptionCheck && this.props.onOptionCheck(evt, isChecked, itemData.id, updatedChecked);
//       }
//     );
//   };
// 
//   render() {
//     const {
//       availableOptionsTitle,
//       availableOptionsActions,
//       availableOptionsSearchAriaLabel,
//       className,
//       children,
//       chosenOptionsTitle,
//       chosenOptionsActions,
//       chosenOptionsSearchAriaLabel,
//       filterOption,
//       isSearchable,
//       chosenOptionsStatus,
//       availableOptionsStatus,
//       controlsAriaLabel,
//       addAllAriaLabel,
//       addSelectedAriaLabel,
//       removeSelectedAriaLabel,
//       removeAllAriaLabel,
//       /* eslint-disable @typescript-eslint/no-unused-vars */
//       availableOptions: consumerPassedAvailableOptions,
//       chosenOptions: consumerPassedChosenOptions,
//       removeSelected,
//       addAll,
//       removeAll,
//       addSelected,
//       onListChange,
//       onAvailableOptionsSearchInputChanged,
//       onChosenOptionsSearchInputChanged,
//       onOptionSelect,
//       onOptionCheck,
//       id,
//       isTree,
//       isDisabled,
//       addAllTooltip,
//       addAllTooltipProps,
//       addSelectedTooltip,
//       addSelectedTooltipProps,
//       removeAllTooltip,
//       removeAllTooltipProps,
//       removeSelectedTooltip,
//       removeSelectedTooltipProps,
//       ...props
//     } = this.props;
//     const {
//       availableOptions,
//       chosenOptions,
//       chosenOptionsSelected,
//       availableOptionsSelected,
//       chosenTreeOptionsChecked,
//       availableTreeOptionsChecked
//     } = this.state;
//     const availableOptionsStatusToDisplay =
//       availableOptionsStatus ||
//       (isTree
//         ? `${
//             filterFolders(availableOptions as DualListSelectorTreeItemData[], availableTreeOptionsChecked).length
//           } of ${flattenTree(availableOptions as DualListSelectorTreeItemData[]).length} items selected`
//         : `${availableOptionsSelected.length} of ${availableOptions.length} items selected`);
//     const chosenOptionsStatusToDisplay =
//       chosenOptionsStatus ||
//       (isTree
//         ? `${filterFolders(chosenOptions as DualListSelectorTreeItemData[], chosenTreeOptionsChecked).length} of ${
//             flattenTree(chosenOptions as DualListSelectorTreeItemData[]).length
//           } items selected`
//         : `${chosenOptionsSelected.length} of ${chosenOptions.length} items selected`);
// 
//     const available = isTree
//       ? availableOptions.map(item => this.mapChecked(item as DualListSelectorTreeItemData, false))
//       : availableOptions;
//     const chosen = isTree
//       ? chosenOptions.map(item => this.mapChecked(item as DualListSelectorTreeItemData, true))
//       : chosenOptions;
// 
//     return (
//       <DualListSelectorContext.Provider value={{ isTree }}>
//         <div className={css(styles.dualListSelector, className)} id={id} {...props}>
//           {children === '' ? (
//             <>
//               <DualListSelectorPane
//                 isSearchable={isSearchable}
//                 onFilterUpdate={this.onFilterUpdate}
//                 searchInputAriaLabel={availableOptionsSearchAriaLabel}
//                 filterOption={filterOption}
//                 onSearchInputChanged={onAvailableOptionsSearchInputChanged}
//                 status={availableOptionsStatusToDisplay}
//                 title={availableOptionsTitle}
//                 options={available}
//                 selectedOptions={isTree ? availableTreeOptionsChecked : availableOptionsSelected}
//                 onOptionSelect={this.onOptionSelect}
//                 onOptionCheck={(e, isChecked, itemData) => this.onTreeOptionCheck(e, isChecked, itemData, false)}
//                 actions={availableOptionsActions}
//                 id={`${id}-available-pane`}
//                 isDisabled={isDisabled}
//               />
//               <DualListSelectorControlsWrapper aria-label={controlsAriaLabel}>
//                 <DualListSelectorControl
//                   isDisabled={
//                     (isTree ? availableTreeOptionsChecked.length === 0 : availableOptionsSelected.length === 0) ||
//                     isDisabled
//                   }
//                   onClick={isTree ? this.addTreeSelected : this.addSelected}
//                   ref={this.addSelectedButtonRef}
//                   aria-label={addSelectedAriaLabel}
//                   tooltipContent={addSelectedTooltip}
//                   tooltipProps={addSelectedTooltipProps}
//                 >
//                   <AngleRightIcon />
//                 </DualListSelectorControl>
//                 <DualListSelectorControl
//                   isDisabled={availableOptions.length === 0 || isDisabled}
//                   onClick={isTree ? this.addAllTreeVisible : this.addAllVisible}
//                   ref={this.addAllButtonRef}
//                   aria-label={addAllAriaLabel}
//                   tooltipContent={addAllTooltip}
//                   tooltipProps={addAllTooltipProps}
//                 >
//                   <AngleDoubleRightIcon />
//                 </DualListSelectorControl>
//                 <DualListSelectorControl
//                   isDisabled={chosenOptions.length === 0 || isDisabled}
//                   onClick={isTree ? this.removeAllTreeVisible : this.removeAllVisible}
//                   aria-label={removeAllAriaLabel}
//                   ref={this.removeAllButtonRef}
//                   tooltipContent={removeAllTooltip}
//                   tooltipProps={removeAllTooltipProps}
//                 >
//                   <AngleDoubleLeftIcon />
//                 </DualListSelectorControl>
//                 <DualListSelectorControl
//                   onClick={isTree ? this.removeTreeSelected : this.removeSelected}
//                   isDisabled={
//                     (isTree ? chosenTreeOptionsChecked.length === 0 : chosenOptionsSelected.length === 0) || isDisabled
//                   }
//                   ref={this.removeSelectedButtonRef}
//                   aria-label={removeSelectedAriaLabel}
//                   tooltipContent={removeSelectedTooltip}
//                   tooltipProps={removeSelectedTooltipProps}
//                 >
//                   <AngleLeftIcon />
//                 </DualListSelectorControl>
//               </DualListSelectorControlsWrapper>
//               <DualListSelectorPane
//                 isChosen
//                 isSearchable={isSearchable}
//                 onFilterUpdate={this.onFilterUpdate}
//                 searchInputAriaLabel={chosenOptionsSearchAriaLabel}
//                 filterOption={filterOption}
//                 onSearchInputChanged={onChosenOptionsSearchInputChanged}
//                 title={chosenOptionsTitle}
//                 status={chosenOptionsStatusToDisplay}
//                 options={chosen}
//                 selectedOptions={isTree ? chosenTreeOptionsChecked : chosenOptionsSelected}
//                 onOptionSelect={this.onOptionSelect}
//                 onOptionCheck={(e, isChecked, itemData) => this.onTreeOptionCheck(e, isChecked, itemData, true)}
//                 actions={chosenOptionsActions}
//                 id={`${id}-chosen-pane`}
//                 isDisabled={isDisabled}
//               />
//             </>
//           ) : (
//             children
//           )}
//         </div>
//       </DualListSelectorContext.Provider>
//     );
//   }
// }
