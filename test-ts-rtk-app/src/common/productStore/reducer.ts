import {productAction, productActionTypes} from "./actions";
import { Product } from "./product";


const productStoreReducer = (state: Product[] = [], action: productAction) => {
    if (action.type === productActionTypes.add) {
        return [
            ...state,
            action.payload
        ]
    }
    if (action.type === productActionTypes.edit) {
        const editIndex = state.findIndex(el => el.id === action.payload.id)
        return [
            ...state.slice(0, editIndex),
            action.payload,
            ...state.slice(editIndex + 1)
        ]
    }
    if (action.type === productActionTypes.delete) {
        const deleteIndex = state.findIndex(el => el.id === action.payload.id)
        return [
            ...state.slice(0, deleteIndex),
            ...state.slice(deleteIndex + 1)
        ]
    }
    return state
}

export default productStoreReducer