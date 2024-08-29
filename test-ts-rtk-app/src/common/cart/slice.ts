import { createSlice } from "@reduxjs/toolkit";
import { loadState } from "../localstorage";

type CartItem = number

type InitialState = CartItem[]

type CartStoreAction = {
    type: string,
    payload: CartItem
}

type CartStoreActionRemove = {
    type: string,
    payload: number
}

const getInitialState = (): InitialState => {
    const localStorage: CartItem[] = loadState("cart")
    return localStorage.length === 0 ? [] : localStorage
}

const cartSlice = createSlice({
    name: "cart",
    initialState: getInitialState(),
    reducers: {
        addToCart(state: InitialState, action: CartStoreAction) {
            state.push(action.payload)
        },
        removeFromCart(state: InitialState, action: CartStoreAction) {
            const removeCandidate = state.findIndex(item => item === action.payload)                      
            state.splice(removeCandidate, 1)
        },
        wipeCart(state: InitialState) { 
            state = []
        } 
    }
})

export const {addToCart, removeFromCart, wipeCart} = cartSlice.actions
export const CartReducer = cartSlice.reducer 