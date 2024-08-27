import { Product } from "../productStore/product"
import { DispatcherProps } from "./dispatcher"


export type CartStoreProps = StateProps & DispatcherProps

export type CartStoreAndOwnProps = StatePropsAndProduct & DispatcherProps

export type CartViewProps = StatePropsForCartView & DispatcherProps

export interface StateProps {
    cart: number[]
}

export interface StatePropsAndProduct {
    cart: number[],
    own: Product
}

export interface StatePropsForCartView {
    cart: number[],
    own: {
        product: Product,
        amount: number
    }
}

export const cartStoreState = (state: any) => ({
    cart: state.cart
})

export const cartStoreStateAndProps = (state: any) => ({
    cart: state.cart
})