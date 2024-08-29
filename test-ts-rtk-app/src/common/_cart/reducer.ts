import { cartAction, cartActionTypes } from "./actions";


const CartReducer = (state: number[] = [], action: cartAction) => {    
    if (action.type === cartActionTypes.add) {
        return [
            ...state,
            action.payload?.id
        ]
    }
    if (action.type === cartActionTypes.remove) {
        const deleteIndex = state.findIndex(el => el === action.payload?.id)
        return [
            ...state.slice(0, deleteIndex),
            ...state.slice(deleteIndex + 1)
        ]
    }
    if (action.type === cartActionTypes.wipe) {
        return []
    }

    return state
}

export default CartReducer