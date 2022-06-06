package dragdrop

type DraggableItemPosition struct {
	// DroppableId - Parent droppableId.
	DroppableId string
	// Index - Index of item in parent Droppable.
	Index int
}

type Props struct {
	// Children - Potentially Droppable and Draggable children. Optional.
	Children any //  // React.ReactNode 
	// OnDrag - Callback for drag event. Return true to allow drag, false to disallow. Optional.
	OnDrag func(source any // DraggableItemPosition ) bool
	// OnDragMove - Callback on mouse move while dragging. Optional.
	OnDragMove func(source any // DraggableItemPosition , dest any // DraggableItemPosition )
	// OnDrop - Callback for drop event. Return true to allow drop, false to disallow. Optional.
	OnDrop func(source any // DraggableItemPosition , dest any // DraggableItemPosition ) bool
}

// contents of react-core/src/components/DragDrop/DragDrop.tsx
// import * as React from 'react';
// 
// interface DraggableItemPosition {
//   /** Parent droppableId */
//   droppableId: string;
//   /** Index of item in parent Droppable */
//   index: number;
// }
// 
// export const DragDropContext = React.createContext({
//   onDrag: (_source: DraggableItemPosition) => true as boolean,
//   onDragMove: (_source: DraggableItemPosition, _dest?: DraggableItemPosition) => {},
//   onDrop: (_source: DraggableItemPosition, _dest?: DraggableItemPosition) => false as boolean
// });
// 
// interface DragDropProps {
//   /** Potentially Droppable and Draggable children */
//   children?: React.ReactNode;
//   /** Callback for drag event. Return true to allow drag, false to disallow. */
//   onDrag?: (source: DraggableItemPosition) => boolean;
//   /** Callback on mouse move while dragging. */
//   onDragMove?: (source: DraggableItemPosition, dest?: DraggableItemPosition) => void;
//   /** Callback for drop event. Return true to allow drop, false to disallow. */
//   onDrop?: (source: DraggableItemPosition, dest?: DraggableItemPosition) => boolean;
// }
// 
// export const DragDrop: React.FunctionComponent<DragDropProps> = ({
//   children,
//   onDrag = () => true,
//   onDragMove = () => {},
//   onDrop = () => false
// }: DragDropProps) => (
//   <DragDropContext.Provider value={{ onDrag, onDragMove, onDrop }}>{children}</DragDropContext.Provider>
// );
// DragDrop.displayName = 'DragDrop';
