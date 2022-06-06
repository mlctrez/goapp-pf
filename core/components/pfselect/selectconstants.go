package pfselect

type SelectContextInterface struct {
	// OnSelect - 
	OnSelect func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */, value any /* string | any // SelectOptionObject  */, isPlaceholder bool)
	// OnClose - 
	OnClose func()
	// OnFavorite - 
	OnFavorite func(itemId string, isFavorite bool)
	// Variant - 
	Variant string
	// InputIdPrefix - 
	InputIdPrefix string
	// ShouldResetOnSelect - 
	ShouldResetOnSelect bool
}

// contents of react-core/src/components/Select/selectConstants.tsx
// import * as React from 'react';
// import { SelectOptionObject } from './SelectOption';
// 
// export interface SelectContextInterface {
//   onSelect: (
//     event: React.MouseEvent<any, MouseEvent> | React.ChangeEvent<HTMLInputElement>,
//     value: string | SelectOptionObject,
//     isPlaceholder?: boolean
//   ) => void;
//   onClose: () => void;
//   onFavorite: (itemId: string, isFavorite: boolean) => void;
//   variant: string;
//   inputIdPrefix: string;
//   shouldResetOnSelect: boolean;
// }
// 
// export const SelectContext = React.createContext<SelectContextInterface | null>(null);
// 
// export const SelectProvider = SelectContext.Provider;
// export const SelectConsumer = SelectContext.Consumer;
// 
// export enum SelectVariant {
//   single = 'single',
//   checkbox = 'checkbox',
//   typeahead = 'typeahead',
//   typeaheadMulti = 'typeaheadmulti'
// }
// 
// export enum SelectPosition {
//   right = 'right',
//   left = 'left'
// }
// 
// export enum SelectDirection {
//   up = 'up',
//   down = 'down'
// }
// 
// export const SelectFooterTabbableItems = 'input, button, select, textarea, a[href]';
