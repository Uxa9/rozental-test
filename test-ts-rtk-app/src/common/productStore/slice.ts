import { createSlice } from "@reduxjs/toolkit";
import { loadState } from "../localstorage";

type Product = {
    id: number,
    price: number,
    name: string,
}

type InitialState = Product[]

type ProductStoreAction = {
    type: string,
    payload: Product
}

type ProductStoreActionDelete = {
    type: string,
    payload: number
}

const getInitialState = (): InitialState => {
    const localStorage: Product[] = loadState("products")    
    return localStorage.length === 0 ? [] : localStorage
}

const productSlice = createSlice({
    name: "product",
    initialState: getInitialState(),
    reducers: {
        addProduct(state: InitialState, action: ProductStoreAction) {
            state.push(action.payload)
        },
        editProduct(state: InitialState, action: ProductStoreAction) {
            const editIndex = state.findIndex(item => item.id === action.payload.id)
            state[editIndex] = action.payload
        },
        deleteProduct(state: InitialState, action: ProductStoreActionDelete) {
            state.filter(item => item.id !== action.payload)
        }
    }
})

export const {addProduct, editProduct, deleteProduct} = productSlice.actions
export const ProductReducer = productSlice.reducer