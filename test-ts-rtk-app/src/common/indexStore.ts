import CartDispatcher from "./cart/dispatcher";
import CartReducer from "./cart/reducer";
import productStoreDispatcher, { DispatcherProps as ProductDispatcherProps } from "./productStore/dispatcher";
import { DispatcherProps as CartDispatcherProps } from "./cart/dispatcher";
import { Product } from "./productStore/product";
import productStoreReducer from "./productStore/reducer";
import { combineReducers } from "redux";

const rootReducer = combineReducers({
    cart: CartReducer,
    products: productStoreReducer
})

export default rootReducer

export type CommonStore = StoreProps & ProductDispatcherProps & CartDispatcherProps

export interface StoreProps {
    cart: number[],
    elements: Product[]
}

export const StoreState = (state: any) => ({
    cart: state.cart,
    elements: state.products
})

export type StoreDispatcherProps = ProductDispatcherProps & CartDispatcherProps

export const StoreDispathcer = {
    ...CartDispatcher,
    ...productStoreDispatcher
}