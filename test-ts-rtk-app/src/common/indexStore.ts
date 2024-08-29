import CartDispatcher from "./_cart/dispatcher";
import { CartReducer } from "./cart/slice";
import productStoreDispatcher, { DispatcherProps as ProductDispatcherProps } from "./_productStore/dispatcher";
import { DispatcherProps as CartDispatcherProps } from "./_cart/dispatcher";
import { Product } from "./_productStore/product";
import { combineReducers } from "redux";
import { ProductReducer } from "./productStore/slice";
import { PostReducer } from "./post/slice";

const rootReducer = combineReducers({
    cart: CartReducer,
    products: ProductReducer,
    post: PostReducer
})

export type RootState = ReturnType<typeof rootReducer>

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